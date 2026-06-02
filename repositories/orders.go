package repositories

import (
	"context"
	"encoding/base64"
	"fmt"

	"cloud.google.com/go/firestore"
	"cloud.google.com/go/storage"
	"github.com/AHSChemicalsOfficialWebsite/AHSChemicalsGCShared/constants"
	"github.com/AHSChemicalsOfficialWebsite/AHSChemicalsGCShared/firebase"
	"github.com/AHSChemicalsOfficialWebsite/AHSChemicalsGCShared/models"
)

func CreateOrderInFirestore(order *models.Order, ctx context.Context) error {

	_, err := firebase.FirestoreClient.Collection(firebase.OrdersCollection).Doc(order.ID).Set(ctx, order.ToMap())
	if err != nil {
		return err
	}
	return nil
}

//Only for the auth role "user"
func CanPlaceOrder(orderUID string, ctx context.Context) error {
	docs, err := firebase.FirestoreClient.Collection(firebase.OrdersCollection).Where("status", "==", constants.OrderStatusPending).Where("uid", "==", orderUID).Documents(ctx).GetAll()
	if err != nil {
		return err
	}
	if len(docs) > 0 {
		return fmt.Errorf("There is already an order pending for your account. Please contact the admin to change the status of the order to either approved or Cancelled in order to place a new order")
	}
	return nil
}

func UploadOrderFileToStorage(orderID string, base64Str string, fileName string, ctx context.Context) error {
	bucket, err := firebase.StorageClient.Bucket(firebase.StorageBucket)
	if err != nil {
		return err
	}
	object := bucket.Object(fmt.Sprintf("%s/%s/%s", firebase.OrdersCollection, orderID, fileName))
	writer := object.NewWriter(ctx)

	data, err := base64.StdEncoding.DecodeString(base64Str)
	if err != nil {
		return fmt.Errorf("failed to decode base64 string: %w", err)
	}

	if _, err := writer.Write(data); err != nil {
		return fmt.Errorf("failed to write to Cloud Storage: %w", err)
	}

	if err := writer.Close(); err != nil {
		return fmt.Errorf("failed to close writer: %w", err)
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

func FetchOrderFromFirestore(orderID string, ctx context.Context) (*models.Order, error) {
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
	customer, err := FetchCustomerFromFirestore(customerID, ctx)
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

func DoesShippingManifestExist(ctx context.Context, orderID string) (bool, error) {
	bucket, err := firebase.StorageClient.Bucket(firebase.StorageBucket)
	if err != nil {
		return false, err
	}
	obj := bucket.Object(fmt.Sprintf("%s/%s", firebase.OrdersCollection, orderID))

	// Try to get object attributes
	_, err = obj.Attrs(ctx)
	if err != nil {
		if err == storage.ErrObjectNotExist {
			return false, nil
		}
		return false, err
	}
	return true, nil
}
