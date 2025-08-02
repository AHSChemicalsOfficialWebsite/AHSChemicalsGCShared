package services

import (
	"time"

	"github.com/HarshMohanSason/AHSChemicalsGCShared/shared/models"
)

func GetUpdatedCustomerDetails(updated, oldCustomer *models.Customer) map[string]any {
	if updated == nil || oldCustomer == nil {
		return nil
	}

	changedValues := make(map[string]any)

	if updated.IsActive != oldCustomer.IsActive {
		changedValues["isActive"] = updated.IsActive
	}
	if updated.Name != oldCustomer.Name {
		changedValues["name"] = updated.Name
	}
	if updated.Email != oldCustomer.Email {
		changedValues["email"] = updated.Email
	}
	if updated.Phone != oldCustomer.Phone {
		changedValues["phone"] = updated.Phone
	}
	if updated.Address1 != oldCustomer.Address1 {
		changedValues["address1"] = updated.Address1
	}
	if updated.City != oldCustomer.City {
		changedValues["city"] = updated.City
	}
	if updated.State != oldCustomer.State {
		changedValues["state"] = updated.State
	}
	if updated.Zip != oldCustomer.Zip {
		changedValues["zip"] = updated.Zip
	}
	if updated.Country != oldCustomer.Country {
		changedValues["country"] = updated.Country
	}

	if len(changedValues) == 0 {
		return nil
	}
	changedValues["updatedAt"] = time.Now()

	return changedValues
}
