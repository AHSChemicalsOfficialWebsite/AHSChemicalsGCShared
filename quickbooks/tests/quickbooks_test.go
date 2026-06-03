package tests

import (
	"context"
	"testing"

	"github.com/AHSChemicalsOfficialWebsite/AHSChemicalsGCShared/mocks"
	"github.com/AHSChemicalsOfficialWebsite/AHSChemicalsGCShared/quickbooks/qbmodels"
	"github.com/AHSChemicalsOfficialWebsite/AHSChemicalsGCShared/quickbooks/qbservices"
)

func TestGetTokenFromFirestore(t *testing.T) {

	tokenResp, err := qbservices.GetTokenUIDFromFirestore(context.Background())
	if err != nil{
		t.Error(err)
	}
	t.Log(tokenResp)
}

func TestCreateQBCustomerFromEntityID(t *testing.T){
	context := context.Background()
	customer, err := qbservices.GetQBEntity[qbmodels.QBCustomersResponse](context, "SELECT * FROM Customer WHERE Id = '68' in ACTIVE (true, false)", "realmID", "")
	if err != nil{
		t.Error(err)
	}
	t.Logf("Customer fetched from quickbooks: %v",customer)
}

func TestQuickbooksEstimate(t *testing.T){
	context := context.Background()
	order := mocks.CreateMockOrder(2)
	
	estimate, err := qbservices.CreateOrderQBEstimate(context, order, nil)
	if err != nil{
		t.Error(err)
	}
	t.Logf("Estimate fetched tax rate from quickbooks for the order is: %v", estimate.GetTotalTaxRate())

	err = qbservices.DeleteOrderQBEstimate(context, estimate, nil)
	if err != nil{
		t.Error(err)
	}
	t.Logf("Estimate deleted from quickbooks")
}
