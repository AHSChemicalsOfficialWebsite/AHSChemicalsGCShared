package validation

import "github.com/AHSChemicalsOfficialWebsite/AHSChemicalsGCShared/models"

func ValidateProductInventoryUpdate(pip *models.ProductInventoryUpdate) error {
	if (len(pip.ProductID) == 0){
		return ErrProductIDRequired
	}
	if (pip.Quantity == 0){
		return ErrQuantityRequired
	}
	if (len(pip.Brand) == 0){
		return ErrBrandIDRequired
	}
	if (len(pip.Name) == 0){
		return ErrLocationIDRequired
	}
	return nil
}