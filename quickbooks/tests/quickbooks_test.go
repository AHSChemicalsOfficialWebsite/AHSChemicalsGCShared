package tests

import (
	"context"
	"testing"

	"github.com/AHSChemicalsOfficialWebsite/AHSChemicalsGCShared/mocks"
	"github.com/AHSChemicalsOfficialWebsite/AHSChemicalsGCShared/quickbooks/qbmodels"
	"github.com/AHSChemicalsOfficialWebsite/AHSChemicalsGCShared/quickbooks/qbservices"
)

func TestGetQBEntity(t *testing.T) {
	context := context.Background()
	tokenResp, err := qbservices.GetValidToken(context)
	if err != nil {
		t.Error(err)
	}
	query := "SELECT * FROM Customer WHERE Id = '106'"
	customer, err := qbservices.GetQBEntity[qbmodels.QBCustomersResponse](context, query, tokenResp.RealmId, tokenResp.AccessToken)
	if err != nil {
		t.Error(err)
	}

	t.Logf("Customer fetched from quickbooks: %v", customer)
}

func TestQuickbooksEstimate(t *testing.T) {
	context := context.Background()
	order := mocks.CreateMockOrder(2)
	tokenResp, err := qbservices.GetValidToken(context)
	if err != nil {
		t.Error(err)
	}
	estimate, err := qbservices.CreateOrderQBEstimate(context, order, tokenResp)
	if err != nil {
		t.Error(err)
	}
	t.Logf("Estimate fetched tax rate from quickbooks for the order is: %v", estimate.GetTotalTaxRate())

	err = qbservices.DeleteOrderQBEstimate(context, estimate, tokenResp)
	if err != nil {
		t.Error(err)
	}
	t.Logf("Estimate deleted from quickbooks")
}
