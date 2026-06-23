package qbservices

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"time"

	"github.com/AHSChemicalsOfficialWebsite/AHSChemicalsGCShared/quickbooks/qbmodels"
)

func FetchQuickbooksProducts(ctx context.Context, requestURL, query string, tokenResponse *qbmodels.QBReponseToken) ([]qbmodels.QBItem, error) {
	params := url.Values{}
	params.Set("query", query)
	params.Set("minorversion", "73")

	fullURL := fmt.Sprintf("%s?%s", requestURL, params.Encode())

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, fullURL, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %v", err)
	}
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", tokenResponse.AccessToken))
	req.Header.Set("Accept", "application/json")

	client := &http.Client{Timeout: 15 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("request failed: %v", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response: %v", err)
	}
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("request failed with status code %d: %s", resp.StatusCode, string(body))
	}
	var qbResp qbmodels.QBItemsResponse
	if err := json.Unmarshal(body, &qbResp); err != nil {
		return nil, fmt.Errorf("failed to parse QuickBooks response: %v", err)
	}

	return qbResp.QueryResponse.Item, nil
}
