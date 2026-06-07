package qbservices

import (
	"bytes"
	"context"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"time"

	"cloud.google.com/go/firestore"
	firebase "github.com/AHSChemicalsOfficialWebsite/AHSChemicalsGCShared/firebase"
	"github.com/AHSChemicalsOfficialWebsite/AHSChemicalsGCShared/quickbooks"
	"github.com/AHSChemicalsOfficialWebsite/AHSChemicalsGCShared/quickbooks/qbmodels"
)

func ExchangeTokenForAuthCode(ctx context.Context, authCode string) (*qbmodels.QBReponseToken, error) {
	tokenURL := "https://oauth.platform.intuit.com/oauth2/v1/tokens/bearer"

	data := url.Values{}
	data.Set("grant_type", "authorization_code")
	data.Set("code", authCode)
	data.Set("redirect_uri", quickbooks.AUTH_CALLBACK_URL)

	authStr := fmt.Sprintf("%s:%s", quickbooks.CLIENT_ID, quickbooks.CLIENT_SECRET)
	baseAuth := base64.StdEncoding.EncodeToString([]byte(authStr))

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, tokenURL, bytes.NewBufferString(data.Encode()))
	if err != nil {
		return nil, err
	}

	authValue := fmt.Sprintf("Basic %s", baseAuth)
	req.Header.Set("Authorization", authValue)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	client := &http.Client{Timeout: 10 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, ReturnErrorFromQBResp(body, "ExchangeTokenForAuthCode")
	}

	var tokenResp qbmodels.QBReponseToken
	if err := json.Unmarshal(body, &tokenResp); err != nil {
		return nil, err
	}

	return &tokenResp, nil
}

func refreshToken(ctx context.Context, refreshToken string) (*qbmodels.QBReponseToken, error) {
	tokenURL := "https://oauth.platform.intuit.com/oauth2/v1/tokens/bearer"

	data := url.Values{}
	data.Set("grant_type", "refresh_token")
	data.Set("refresh_token", refreshToken)

	authStr := fmt.Sprintf("%s:%s", quickbooks.CLIENT_ID, quickbooks.CLIENT_SECRET)
	baseAuth := base64.StdEncoding.EncodeToString([]byte(authStr))

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, tokenURL, bytes.NewBufferString(data.Encode()))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", fmt.Sprintf("Basic %s", baseAuth))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	client := &http.Client{Timeout: 10 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, ReturnErrorFromQBResp(body, "RefreshToken")
	}

	var tokenResp qbmodels.QBReponseToken
	if err := json.Unmarshal(body, &tokenResp); err != nil {
		return nil, err
	}
	return &tokenResp, nil
}

func GetValidToken(ctx context.Context) (*qbmodels.QBReponseToken, error) {
    docRefs, err := firebase.FirestoreClient.Collection(firebase.QuickBooksTokenCollection).Documents(ctx).GetAll()
    if err != nil {
        return nil, err
    }
    if len(docRefs) == 0 {
        return nil, errors.New("No QuickBooks authentication, please re-authenticate it")
    }
    for _, docRef := range docRefs {
        var token qbmodels.QBReponseToken
        if err := docRef.DataTo(&token); err != nil {
            continue
        }
        if token.IsRefreshTokenExpired() {
            continue
        }
        if token.IsExpired() {
            // refresh it and return
            newToken, err := refreshToken(ctx, token.RefreshToken)
            if err != nil {
                continue // try next token
            }
            newToken.SetRealmID(token.RealmId)
            newToken.SetState(token.State)
            SaveTokenToFirestore(ctx, newToken, docRef.Ref.ID)
            return newToken, nil
        }
        return &token, nil
    }
    return nil, ErrNoQuickBooksAuth
}

func SaveTokenToFirestore(ctx context.Context, t *qbmodels.QBReponseToken, uid string) error {
	time := time.Now().UTC()
	t.SetObtainedAt(time)
	t.SetExpiresAt(time)

	_, err := firebase.FirestoreClient.Collection(firebase.QuickBooksTokenCollection).Doc(uid).Set(ctx, t.ToMap(), firestore.MergeAll)
	return err
}