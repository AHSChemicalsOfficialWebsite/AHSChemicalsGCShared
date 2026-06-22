package mocks

import (
	"time"

	"github.com/AHSChemicalsOfficialWebsite/AHSChemicalsGCShared/models"
)

func CreateMockOrder(itemLength int) *models.Order{
	return &models.Order{
		ID:"1",
		InvoiceNumber: "1234",
		Customer: CreateMockCustomer(),
		Items: CreateMockOrderItems(itemLength),
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