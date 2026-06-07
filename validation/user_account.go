package validation

import (
	"github.com/AHSChemicalsOfficialWebsite/AHSChemicalsGCShared/models"
	"github.com/AHSChemicalsOfficialWebsite/AHSChemicalsGCShared/utils"
)

// Only validating the required fields. Email, name etc can be validated in firebase auth
func ValidateUserAccount(c *models.UserAccountCreate) error {
	if err := validateName(c.Name); err != nil {
		return ErrNameInvalid
	}
	if len(c.Customers) == 0 {
		return ErrNoCustomers
	}
	if utils.HasDuplicateStrings(c.Customers) {
		return ErrDuplicateCustomers
	}
	if len(c.Brands) == 0 {
		return ErrNoBrands
	}
	if utils.HasDuplicateStrings(c.Brands) {
		return ErrDuplicateBrands
	}
	if c.Role == "" {
		return ErrNoRole
	}
	if _, ok := models.ValidRoles[models.Role(c.Role)]; !ok {
		return ErrInvalidRole
	}
	return nil
}
