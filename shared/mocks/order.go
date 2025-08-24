package mocks

import (
	"time"

	"github.com/HarshMohanSason/AHSChemicalsGCShared/shared/models"
)

func CreateMockOrder() *models.Order{
	return &models.Order{
		ID:"1",
		Customer: CreateMockCustomer(),
		Items: CreateMockProducts(5),
		SpecialInstructions: "Test instrucitons",
		Total: 100,
		SubTotal: 200,
		Status: "PENDING",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		TaxRate: 0.0124,
		TaxAmount: 20,
	}
}