package repositories

import (
	"context"
	"fmt"

	"cloud.google.com/go/firestore"
	"github.com/AHSChemicalsOfficialWebsite/AHSChemicalsGCShared/firebase"
	"github.com/AHSChemicalsOfficialWebsite/AHSChemicalsGCShared/models"
	"github.com/AHSChemicalsOfficialWebsite/AHSChemicalsGCShared/quickbooks/qbmodels"
)

func FetchCustomerFromFirestore(ctx context.Context, id string) (*models.Customer, error) {
	docSnapshot, err := firebase.FirestoreClient.Collection(firebase.CustomersCollection).Doc(id).Get(ctx)
	if err != nil {
		return nil, err
	}
	var customer models.Customer
	err = docSnapshot.DataTo(&customer)
	if err != nil {
		return nil, fmt.Errorf("Error getting customer: %v", err)
	}
	return &customer, nil
}

// Only fetches active customers from firestore collection ('customers')
func FetchAllActiveCustomersFromFirestore(ctx context.Context) ([]*models.Customer, error) {
	customers, err := firebase.FirestoreClient.Collection(firebase.CustomersCollection).Where("isActive", "==", true).Documents(ctx).GetAll()
	if err != nil {
		return nil, err
	}
	customersList := make([]*models.Customer, len(customers))
	for i, customer := range customers {
		var customerObj models.Customer
		err = customer.DataTo(&customerObj)
		if err != nil {
			return nil, fmt.Errorf("Error getting customer: %v", err)
		}
		customersList[i] = &customerObj
	}
	return customersList, nil
}

func SyncQuickbookCustomerRespToFirestore(ctx context.Context, qbCustomers []qbmodels.QBCustomer) error {
	bulkWriter := firebase.FirestoreClient.BulkWriter(ctx)
	for _, customer := range qbCustomers {
		docRef := firebase.FirestoreClient.Collection(firebase.CustomersCollection).Doc(customer.ID)
		_, err := bulkWriter.Set(docRef, customer.MapToCustomer().ToMap(), firestore.MergeAll)
		if err != nil {
			return err
		}
	}
	bulkWriter.Flush()
	return nil
}