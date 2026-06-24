package repositories

import (
	"context"
	"fmt"
	"time"

	"cloud.google.com/go/firestore"
	"github.com/AHSChemicalsOfficialWebsite/AHSChemicalsGCShared/firebase"
	"github.com/AHSChemicalsOfficialWebsite/AHSChemicalsGCShared/models"
	gonanoid "github.com/matoous/go-nanoid/v2"
)

// Only for the auth role "user"
func CanPlaceOrder(ctx context.Context, orderUID string) error {
	docs, err := firebase.FirestoreClient.Collection(firebase.OrdersCollection).Where("status", "==", models.OrderStatusPending).Where("uid", "==", orderUID).Documents(ctx).GetAll()
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
	err = SetProductsInOrderItem(ctx, order.Items)
	if err != nil {
		return nil, err
	}

	return &order, nil
}

func BuildOrderFromRequest(ctx context.Context, req *models.OrderRequest, prices map[string]float64) (*models.Order, error) {
	customer, err := FetchCustomerFromFirestore(ctx, req.CustomerID)
	if err != nil {
		return nil, err
	}
	order := &models.Order{
		Status:              models.OrderStatusPending,
		ID:                  gonanoid.Must(8),
		Customer:            customer,
		SpecialInstructions: req.SpecialInstructions,
		CreatedAt:           time.Now().UTC(),
	}

	//Get the products from firestore for each item
	docRefs := make([]*firestore.DocumentRef, len(req.Items))
	for i, item := range req.Items {
		docRefs[i] = firebase.FirestoreClient.Collection(firebase.ProductsCollection).Doc(item.ProductID)
	}

	docSnapshots, err := firebase.FirestoreClient.GetAll(ctx, docRefs)
	if err != nil{
		return nil, err
	}

	for i, docSnapshot := range docSnapshots {
		var product models.Product
		if err := docSnapshot.DataTo(&product); err != nil {
			return nil, err
		}
		order.Items = append(order.Items, &models.OrderItem{
			Product:    &product,
			Quantity:   req.Items[i].Quantity,
			Price:      prices[product.ID],
			ProductID:  product.ID,
		})
	}
	return order, nil
}
