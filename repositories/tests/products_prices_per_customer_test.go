package tests

import (
	"context"
	"testing"

	"github.com/AHSChemicalsOfficialWebsite/AHSChemicalsGCShared/repositories"
)

func TestProductsPricesPerCustomer(t *testing.T) {
	
	err := repositories.SaveProductsPricesPerCustomerToFirestore(context.Background())
	if err != nil {
		t.Error(err)
	}
}

func TestGetProductPricesFromCustomerID(t *testing.T) {
	
	pricesMap, err := repositories.GetProductPricesFromCustomerID(context.Background(),"1")
	if err != nil {
		t.Error(err)
	}
	t.Log("Fetched Map is", pricesMap)
}