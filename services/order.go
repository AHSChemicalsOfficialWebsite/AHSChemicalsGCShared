package services

import (
	"context"
	"errors"
	"time"

	"github.com/AHSChemicalsOfficialWebsite/AHSChemicalsGCShared/models"
	"github.com/AHSChemicalsOfficialWebsite/AHSChemicalsGCShared/repositories"
)

// approveOrder changes the status of an order to APPROVED
// An order can only be approved if
//   - it is in PENDING state
//   - it is in Cancelled state and no more than 30 days have passed since rejection
//
// Parameters:
//   - editedOrder: Pointer to an Order object containing the edited order.
//   - originalOrder: Pointer to an Order object containing the original order.
//
// Returns:
//   - error: An error object if edited status is incorrect .
func approveOrder(editedOrder, originalOrder *models.Order) error {
	switch originalOrder.Status {
	case models.OrderStatusPending:
		editedOrder.Status = models.OrderStatusApproved
		return nil
	case models.OrderStatusCancelled:
		if time.Since(originalOrder.UpdatedAt) > 30*24*time.Hour {
			return errors.New("Cannot approve: more than 30 days have passed since rejection. Place a new order instead")
		}
		editedOrder.Status = models.OrderStatusApproved
		return nil
	default:
		return errors.New("Order can only be approved if it is in PENDING or Cancelled state")
	}
}

// cancelOrder changes the status of an order to CANCELLED.
// An order can only be Cancelled if
//   - it is in PENDING state
//
// Parameters:
//   - editedOrder: Pointer to an Order object containing the edited order.
//   - originalOrder: Pointer to an Order object containing the original order.
//
// Returns:
//   - error: An error object if edited status is incorrect .
func cancelOrder(editedOrder *models.Order, originalOrderStatus models.OrderStatus) error {
	switch originalOrderStatus {
	case models.OrderStatusPending:
		editedOrder.Status = models.OrderStatusCancelled
		return nil
	default:
		return errors.New("Order can only be Cancelled if it is in PENDING state")
	}
}

// deliverOrder changes the status of an order to DELIVERED
// An order can only be delivered if
//   - if the shipping manifest has has been generated. 
//   - it is in APPROVED state.
//
// Once changed, the changes are irreverisble to maintain atomicity and consistency everywhere.
// In Addition to this, A Shipping Manifest is already generated for the order which cannot be redone due
// to the images and signature taken during delivery
//
// Parameters:
//   - editedOrder: Pointer to an Order object containing the edited order.
//
// Returns:
//   - error: An error object if edited status is incorrect .
func deliverOrder(editedOrder, originalOrder *models.Order) error {
	exists, err := repositories.DoesShippingManifestExist(context.Background(), editedOrder.ID)
	if err != nil {
		return errors.New("failed to check for a shipping manifest. Please try again")
	}
	if !exists {
		return errors.New("order has not been delivered yet. Please generate a shipping manifest first")
	}
	switch originalOrder.Status {
	case models.OrderStatusApproved:
		editedOrder.Status = models.OrderStatusDelivered
		return nil
	default:
		return errors.New("Order can only be delivered if it is in APPROVED state")
	}
}

// GetOrderWithUpdatedStatus updates the status of an order based on the edited order
//
// Parameters:
//   - editedOrder: Pointer to an Order object containing the edited order.
//   - originalOrder: Pointer to an Order object containing the original order.
//
// Returns:
//   - error: An error object if edited status is incorrect .
func GetOrderWithUpdatedStatus(editedOrder, originalOrder *models.Order) error {

	switch editedOrder.Status {
	case models.OrderStatusApproved:
		return approveOrder(editedOrder, originalOrder)
	case models.OrderStatusCancelled:
		return cancelOrder(editedOrder, originalOrder.Status)
	case models.OrderStatusDelivered:
		return deliverOrder(editedOrder, originalOrder)
	default:
		return nil
	}
}