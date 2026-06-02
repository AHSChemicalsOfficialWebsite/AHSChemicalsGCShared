package mocks

import (
	"time"

	"github.com/AHSChemicalsOfficialWebsite/AHSChemicalsGCShared/models"
)

func CreateMockProduct() *models.Product {
	return &models.Product{
		ID:            "prod-001",
		IsActive:      true,
		Brand:         "Acme",
		Name:          "Industrial Cleaner",
		SKU:           "ACM-CLN-500ML",
		Size:          500.0,
		SizeUnit:      "ml",
		PackOf:        12,
		Category:      "Chemicals",
		Price:         129.99,
		PurchasePrice: 50.00,
		Desc:          "Highly effective industrial cleaning solution.",
		Slug:          "industrial-cleaner-500ml",
		NameKey:       "industrialcleaner",
		CreatedAt:     time.Now().UTC(),
		UpdatedAt:     time.Now().UTC(),
	}
}

func CreateMockOrderItems(numberOfItems int) []*models.CartItem {
	products := make([]*models.CartItem, numberOfItems)
	for i := range numberOfItems {
		products[i].Product = CreateMockProduct()
		products[i].Quantity = 5
	}
	return products
}
