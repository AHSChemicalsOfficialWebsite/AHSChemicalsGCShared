package tests

import (
	"context"
	"testing"

	"github.com/HarshMohanSason/AHSChemicalsGCShared/shared/quickbooks/qbservices"
)

func TestGetTokenFromFirestore(t *testing.T) {

	tokenResp, err := qbservices.GetTokenUIDFromFirestore(context.Background())
	if err != nil{
		t.Error(err)
	}
	t.Log(tokenResp)
}

func TestCreateQBCustomerFromEntityID(t *testing.T){

	customer, err := qbservices.CreateQBCustomerFromEntityID("1")
	if err != nil{
		t.Error(err)
	}
	t.Logf("Customer fetched from quickbooks: %v",customer)
}

func TestCreateQBItemFromEntityID(t *testing.T){

	item, err := qbservices.CreateQBProductFromEntityID("68")
	if err != nil{
		t.Error(err)
	}
	t.Logf("Item fetched from quickbooks: %v", item)
}
