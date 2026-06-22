package qbservices

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/AHSChemicalsOfficialWebsite/AHSChemicalsGCShared/quickbooks/qbmodels"
	"github.com/AHSChemicalsOfficialWebsite/AHSChemicalsGCShared/models"
	"github.com/AHSChemicalsOfficialWebsite/AHSChemicalsGCShared/quickbooks"
)

func CreateOrderQBEstimate(ctx context.Context, order *models.Order, tokenResp *qbmodels.QBReponseToken) (*qbmodels.QBEstimate, error) {

	newEstimate := qbmodels.NewQBEstimate(order)
	estimateJSON, err := json.Marshal(newEstimate)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal estimate: %w", err)
	}

	// Hit QuickBooks API
	apiURL := fmt.Sprintf("%s/v3/company/%s/estimate", quickbooks.API_URL, tokenResp.RealmId)
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, apiURL, bytes.NewBuffer(estimateJSON))
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}
	req.Header.Set("Authorization", "Bearer "+tokenResp.AccessToken)
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{Timeout: 15 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("request to QuickBooks failed: %w", err)
	}
	defer resp.Body.Close()

	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read QuickBooks response: %w", err)
	}
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("Failed to create estimate in quickbooks: %s", string(bodyBytes))
	}

	var quickbooksItem qbmodels.QBEstimateResponse
	if err := json.Unmarshal(bodyBytes, &quickbooksItem); err != nil {
		return nil, fmt.Errorf("failed to unmarshal QuickBooks response: %w", err)
	}

	return quickbooksItem.Estimate, nil
}

func DeleteOrderQBEstimate(ctx context.Context, estimate *qbmodels.QBEstimate, tokenResp *qbmodels.QBReponseToken) error {
	if estimate.ID == "" || estimate.SyncToken == "" {
		return fmt.Errorf("estimate ID and SyncToken are required to delete")
	}

	deletePayload := map[string]any{
		"Id":        estimate.ID,
		"SyncToken": estimate.SyncToken,
		"sparse":    true,
		"status":    "Deleted",
	}
	payloadBytes, err := json.Marshal(deletePayload)
	if err != nil {
		return fmt.Errorf("failed to marshal delete payload: %w", err)
	}

	// Hit QuickBooks API
	apiURL := fmt.Sprintf("%s/v3/company/%s/estimate?operation=delete", quickbooks.API_URL, tokenResp.RealmId)
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, apiURL, bytes.NewBuffer(payloadBytes))
	if err != nil {
		return fmt.Errorf("failed to create request: %w", err)
	}
	req.Header.Set("Authorization", "Bearer "+tokenResp.AccessToken)
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{Timeout: 15 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("request to QuickBooks failed: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("Failed to delete estimate in quickbooks: %s", string(body))
	}
	return nil
}
