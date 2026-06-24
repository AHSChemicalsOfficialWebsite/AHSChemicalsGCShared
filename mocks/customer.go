package mocks

import (

	"github.com/AHSChemicalsOfficialWebsite/AHSChemicalsGCShared/models"
)

func CreateMockCustomer() *models.CustomerRequest {	
	return &models.CustomerRequest{
		ID: "106",
		Name: "Alexis Party Rental",
		Email: "test_email@gmail.com",
		Address1: "2040 N Preisker lane",
		City: "Las Vegas",
		State: "NV",
		Zip: "89101",
		Phone: "702-555-5555",
		IsActive: true,
		Country: "US",
	}
}

func CreateMockCustomers(val int) []*models.CustomerRequest{
	customers := make([]*models.CustomerRequest, val)
	for range val {
		customers = append(customers, CreateMockCustomer())
	}
	return customers
}