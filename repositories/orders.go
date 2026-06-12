package repositories

import (
	"context"
	"fmt"

	"cloud.google.com/go/firestore"
	"github.com/AHSChemicalsOfficialWebsite/AHSChemicalsGCShared/firebase"
	"github.com/AHSChemicalsOfficialWebsite/AHSChemicalsGCShared/models"
)

//Only for the auth role "user"
func CanPlaceOrder(ctx context.Context, orderUID string) error {
	docs, err := firebase.FirestoreClient.Collection(firebase.OrdersCollection).Where("status", "==", models.OrderStatusPending).Where("uid", "==", orderUID).Where("role", "==", models.UserRole).Documents(ctx).GetAll()
	if err != nil {
		return err
	}
	if len(docs) > 0 {
		return fmt.Errorf("There is already an order pending for your account. Please contact the admin to change the status of the order to either approved or Cancelled in order to place a new order")
	}
	return nil
}

func UpdateOrderInFirestore(ctx context.Context, orderID string, updates []firestore.Update) error {
	_, err := firebase.FirestoreClient.Collection(firebase.OrdersCollection).Doc(orderID).Update(ctx, updates)
	if err != nil {
		return err
	}
	return nil
}

func FetchOrderFromFirestore(ctx context.Context, orderID string) (*models.Order, error) {
	docSnapshot, err := firebase.FirestoreClient.Collection(firebase.OrdersCollection).Doc(orderID).Get(ctx)
	if err != nil {
		return nil, err
	}
	if !docSnapshot.Exists() {
		return nil, fmt.Errorf("order with ID %s not found", orderID)
	}
	//Assign the order object
	var order models.Order
	if err := docSnapshot.DataTo(&order); err != nil {
		return nil, err
	}
	order.SetID(docSnapshot.Ref.ID)

	//Set customer
	customerID := docSnapshot.Data()["customerId"].(string)
	customer, err := FetchCustomerFromFirestore(ctx, customerID)
	if err != nil {
		return nil, err
	}
	order.Customer = customer

	//Set products
	err = SetProductsInCartItem(ctx, order.Items)
	if err != nil {
		return nil, err
	}

	return &order, nil
}
