package repositories

import (
	"context"
	"fmt"

	"cloud.google.com/go/firestore"
	"github.com/AHSChemicalsOfficialWebsite/AHSChemicalsGCShared/firebase"
	"github.com/AHSChemicalsOfficialWebsite/AHSChemicalsGCShared/models"
	"github.com/AHSChemicalsOfficialWebsite/AHSChemicalsGCShared/quickbooks/qbmodels"
)

func FetchAllActiveProductsFromFirestore(ctx context.Context) ([]*models.Product, error) {
	docSnapshots, err := firebase.FirestoreClient.Collection(firebase.ProductsCollection).Documents(ctx).GetAll()
	if err != nil {
		return nil, err
	}
	products := make([]*models.Product, len(docSnapshots))
	for i, docSnapshot := range docSnapshots {
		var product models.Product
		if err := docSnapshot.DataTo(&product); err != nil {
			return nil, fmt.Errorf("error decoding product %s: %v", docSnapshot.Ref.ID, err)
		}
		products[i] = &product
	}
	return products, nil
}

// SetProductsInCartItem sets products in cart item by fetching actual product from firestore
func SetProductsInCartItem(ctx context.Context, cartItems []*models.CartItem) error {
	docRefs := make([]*firestore.DocumentRef, len(cartItems))
	for i, cartItem := range cartItems {
		docRefs[i] = firebase.FirestoreClient.Collection(firebase.ProductsCollection).Doc(cartItem.ProductID)
	}
	docSnapshots, err := firebase.FirestoreClient.GetAll(ctx, docRefs)
	if err != nil {
		return err
	}
	for i, docSnapshot := range docSnapshots {
		var product models.Product
		if err := docSnapshot.DataTo(&product); err != nil {
			return fmt.Errorf("error decoding product %s: %v", cartItems[i].ProductID, err)
		}
		cartItems[i].Product = &product
	}
	return nil
}

// SyncQuickbookProductRespToFirestore syncs quickbook product response to firestore
// collection ('products')
//
// Params:
//   - qbItemsResponse: *qbmodels.QBItemsResponse, a mapped response from quickbooks
//   - ctx: context
//
// Returns:
//   - error: error
func SyncQuickbookProductRespToFirestore(ctx context.Context, qbItems []qbmodels.QBItem) error {
	bulkWriter := firebase.FirestoreClient.BulkWriter(ctx)
	for _, item := range qbItems {
		docRef := firebase.FirestoreClient.Collection(firebase.ProductsCollection).Doc(item.ID)
		_, err := bulkWriter.Set(docRef, item.MapToProduct().ToMap(), firestore.MergeAll)
		if err != nil {
			return err
		}
	}
	bulkWriter.Flush()
	return nil
}
