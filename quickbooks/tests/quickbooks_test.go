package tests

import (
	"context"
	"testing"

	"github.com/AHSChemicalsOfficialWebsite/AHSChemicalsGCShared/quickbooks/qbmodels"
	"github.com/AHSChemicalsOfficialWebsite/AHSChemicalsGCShared/quickbooks/qbservices"
	"github.com/AHSChemicalsOfficialWebsite/AHSChemicalsGCShared/repositories"
)

func TestGetTokenFromFirestore(t *testing.T) {

	tokenResp, err := qbservices.GetTokenUIDFromFirestore(context.Background())
	if err != nil{
		t.Error(err)
	}
	t.Log(tokenResp)
}

//query := fmt.Sprintf("SELECT * FROM Customer WHERE Id = '%s' AND Active IN (true, false)", customerID)
func TestCreateQBCustomerFromEntityID(t *testing.T){
	context := context.Background()
	customer, err := qbservices.GetQBEntity[qbmodels.QBCustomersResponse](context, "SELECT * FROM Customer WHERE Id = '68' in ACTIVE (true, false)", "realmID", "")
	if err != nil{
		t.Error(err)
	}
	t.Logf("Customer fetched from quickbooks: %v",customer)
}

//SELECT * FROM Item WHERE Id = '%s' AND Active IN (true, false)"
func TestQuickbooksEstimate(t *testing.T){
	context := context.Background()
	order, err := repositories.FetchOrderFromFirestore("CTZAWb", context)
	if err != nil{
		t.Error(err)
	}
	t.Logf("Order fetched from firestore: %v", order)
	
	estimate, err := qbservices.CreateOrderQBEstimate(context,order)
	if err != nil{
		t.Error(err)
	}
	t.Logf("Estimate fetched tax rate from quickbooks for the order is: %v", estimate.GetTotalTaxRate())

	err = qbservices.DeleteOrderQBEstimate(context, estimate)
	if err != nil{
		t.Error(err)
	}
	t.Logf("Estimate deleted from quickbooks")
}