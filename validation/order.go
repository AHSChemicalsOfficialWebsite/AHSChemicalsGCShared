package validation

import (
	"regexp"
	"strings"
	"time"

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
	if o.Customer == nil {
		return ErrNoCustomerFound
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

func CanApproveOrder(originalOrder *models.Order) error {
    switch originalOrder.Status {
    case models.OrderStatusPending:
        return nil
    case models.OrderStatusCancelled:
        if time.Since(originalOrder.UpdatedAt) > 30*24*time.Hour {
            return ErrCannotApproveExpiredCancellation
        }
        return nil
    default:
        return ErrCannotApproveNotPendingOrCancelled
    }
}

func CanCancelOrder(originalOrder *models.Order) error {
    switch originalOrder.Status {
    case models.OrderStatusDelivered:
        return ErrCannotCancelDelivered
    case models.OrderStatusCancelled:
        return ErrCannotCancelAlreadyCancelled
    default:
        return nil
    }
}