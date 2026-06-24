package mocks

import (
	"time"

	"github.com/AHSChemicalsOfficialWebsite/AHSChemicalsGCShared/models"
)

func CreateMockProduct() *models.Product {
	return &models.Product{
		ID:            "245",
		IsActive:      true,
		Brand:         "Acme",
		Name:          "Stain Pro 1 Reclaim",
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

func CreateMockOrderItems(numberOfItems int) []*models.OrderItem {
	items := make([]*models.OrderItem, numberOfItems)
	for i := range numberOfItems {
		items[i] = &models.OrderItem{
			Product:  CreateMockProduct(),
			Quantity: 5,
			Price:    129.99,
		}
	}
	return items
}
