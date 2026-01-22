package tests

import (
	"context"
	"testing"

	"github.com/AHSChemicalsOfficialWebsite/AHSChemicalsGCShared/shared/quickbooks/qbservices"
	"github.com/AHSChemicalsOfficialWebsite/AHSChemicalsGCShared/shared/repositories"
)

func TestGetTokenFromFirestore(t *testing.T) {

	tokenResp, err := qbservices.GetTokenUIDFromFirestore(context.Background())
	if err != nil{
		t.Error(err)
	}
	t.Log(tokenResp)
}

func TestCreateQBCustomerFromEntityID(t *testing.T){

	customer, err := qbservices.GetQBCustomerFromEntityID("1")
	if err != nil{
		t.Error(err)
	}
	t.Logf("Customer fetched from quickbooks: %v",customer)
}

func TestCreateQBItemFromEntityID(t *testing.T){

	item, err := qbservices.GetQBProductFromEntityID("68")
	if err != nil{
		t.Error(err)
	}
	t.Logf("Item fetched from quickbooks: %v", item)
}

func TestQuickbooksEstimate(t *testing.T){
	order, err := repositories.FetchDetailedOrderFromFirestore("CTZAWb", context.Background())
	if err != nil{
		t.Error(err)
	}
	t.Logf("Order fetched from firestore: %v", order.Items[0].Name)
	
	estimate, err := qbservices.CreateOrderQBEstimate(order)
	if err != nil{
		t.Error(err)
	}
	t.Logf("Estimate fetched tax rate from quickbooks for the order is: %v", estimate.GetTotalTaxRate())

	err = qbservices.DeleteQBEstimate(estimate)
	if err != nil{
		t.Error(err)
	}
	t.Logf("Estimate deleted from quickbooks")
}