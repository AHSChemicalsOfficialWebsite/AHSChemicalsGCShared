package validation

import (
	"regexp"
	"strings"
	"github.com/AHSChemicalsOfficialWebsite/AHSChemicalsGCShared/models"
)

var illegalCharsRegex = regexp.MustCompile(`[<>[\]{}^*~|\\]`)

func ValidateOrder(o *models.Order) error {
	if o.Uid == "" {
		return ErrNoUserID
	}
	if len(o.Items) == 0 {
		return ErrNoItems
	}
	for _, item := range o.Items {
		if item.Quantity == 0 {
			return ErrItemZeroQty
		}
		if item.Product == nil {
			return ErrItemNoProduct
		}
	}
	return validateSpecialInstructions(o.SpecialInstructions)
}

func validateSpecialInstructions(specialInstructions string) error {
	if specialInstructions == "" {
		return nil
	}
	instructions := strings.TrimSpace(specialInstructions)
	if illegalCharsRegex.MatchString(instructions) {
		return ErrSpecialInstructionsInvalidChars
	}
	if len(instructions) > 150 {
		return ErrSpecialInstructionsTooLong
	}
	return nil
}