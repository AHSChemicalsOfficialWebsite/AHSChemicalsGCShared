package qbservices

import (
	"context"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"time"

	"github.com/AHSChemicalsOfficialWebsite/AHSChemicalsGCShared/quickbooks"
	"github.com/go-jose/go-jose/v4/json"
)

func VerifyQuickBooksWebhookSignature(body []byte, signature string) bool {

	mac := hmac.New(sha256.New, []byte(quickbooks.WEBHOOK_VERIFY_TOKEN))
	mac.Write(body)

	expectedMAC := mac.Sum(nil)
	//Quickbooks sends the signature in base64 format
	expectedSignature := base64.StdEncoding.EncodeToString(expectedMAC) 

	return hmac.Equal([]byte(expectedSignature), []byte(signature))
}

func GetQBEntity[T any](ctx context.Context, query, realmID, accessToken string) (*T, error) {
	requestURL := fmt.Sprintf("%s/v3/company/%s/query", quickbooks.API_URL, realmID)
	params := url.Values{}
	params.Set("query", query)
	params.Set("minorversion", "73")
	fullURL := fmt.Sprintf("%s?%s", requestURL, params.Encode())

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, fullURL, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", accessToken))
	req.Header.Set("Accept", "application/json")

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
		return nil, fmt.Errorf("quickbooks request failed: %s", string(body))
	}

	var result T
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, err
	}
	return &result, nil
}