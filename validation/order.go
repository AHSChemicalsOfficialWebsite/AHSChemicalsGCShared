package validation

import (
	"regexp"
	"strings"
	"time"

	"github.com/AHSChemicalsOfficialWebsite/AHSChemicalsGCShared/models"
)

var illegalCharsRegex = regexp.MustCompile(`[<>[\]{}^*~|\\]`)

func ValidateOrderRequest(o *models.OrderRequest) error {
	if o.CustomerID == ""{
		return ErrCustomerIDRequired
	}
	if len(o.Items) == 0 {
		return ErrOrderItemsRequired
	}
	for _, item := range o.Items {
		if item.ProductID == "" {
			return ErrProductIDRequired
		}
		if item.Quantity == 0 {
			return ErrOrderItemQuantityRequired
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

func AreEqualPrices(a, b []*models.OrderItem) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i].Price != b[i].Price {
			return false
		}
	}
	return true
}

func AreEqualQuantities(a, b []*models.OrderItem) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i].Quantity != b[i].Quantity {
			return false
		}
	}
	return true
}