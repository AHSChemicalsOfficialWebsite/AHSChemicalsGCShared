package tests

import (
	"context"
	"testing"

	"github.com/HarshMohanSason/AHSChemicalsGCShared/shared/repositories"
)

func TestProductsPricesPerCustomer(t *testing.T) {
	
	err := repositories.SaveProductsPricesPerCustomerToFirestore(context.Background())
	if err != nil {
		t.Error(err)
	}
}

func TestGetProductsPricesPerCustomer(t *testing.T) {
	
	pricesMap, err := repositories.GetProductPricesFromCustomerID("1",context.Background())
	if err != nil {
		t.Error(err)
	}
	t.Log("Fetched Map is", pricesMap)
}