package mocks

import (
	"time"

	"github.com/HarshMohanSason/AHSChemicalsGCShared/shared/models"
)

func CreateMockCustomer() models.Customer {	
	return models.Customer{
		ID: "1",
		Name: "Harsh",
		Email: "pFJqA@example.com",
		Address1: "2040 N Preisker lane",
		City: "Las Vegas",
		State: "NV",
		Zip: "89101",
		Phone: "702-555-5555",
		IsActive: true,
		Country: "US",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}

func CreateMockCustomers(val int) []models.Customer{
	customers := make([]models.Customer, val)
	for i := 0; i < val; i++ {
		customers = append(customers, CreateMockCustomer())
	}
	return customers
}