package repositories

import (
	"context"
	"fmt"

	"cloud.google.com/go/firestore"
	"github.com/AHSChemicalsOfficialWebsite/AHSChemicalsGCShared/firebase"
	"github.com/AHSChemicalsOfficialWebsite/AHSChemicalsGCShared/models"
	"github.com/AHSChemicalsOfficialWebsite/AHSChemicalsGCShared/quickbooks/qbmodels"
)

// FetchAllProductsByIDs fetches all products from firestore from collection ('products')
// Returns map of product id and the product object in order to search in O(1)
//
// Params:
//   - ctx: context
//   - productIDs: []string, product ids
//
// Returns:
//   - map[string]models.Product. Key is product id
//   - error
func FetchAllProductsByIDs(ctx context.Context, productIDs []string) (map[string]*models.Product, error) {
	docRefs := make([]*firestore.DocumentRef, len(productIDs))
	for i, productID := range productIDs {
		docRefs[i] = firebase.FirestoreClient.Collection(firebase.ProductsCollection).Doc(productID)
	}
	docSnapshots, err := firebase.FirestoreClient.GetAll(ctx, docRefs)
	if err != nil {
		return nil, err
	}

	productMap := make(map[string]*models.Product)
	for i, docSnapshot := range docSnapshots {
		var product models.Product
		if err := docSnapshot.DataTo(&product); err != nil {
			return nil, fmt.Errorf("error decoding product %s: %v", productIDs[i], err)
		}
		productMap[product.ID] = &product
	}
	return productMap, nil
}

func FetchAllProductsFromFirestore(ctx context.Context) ([]*models.Product, error) {
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
func SyncQuickbookProductRespToFirestore(qbItemsResponse *qbmodels.QBItemsResponse, ctx context.Context) error {
	if qbItemsResponse == nil || qbItemsResponse.QueryResponse.Item == nil {
		return nil
	}

	bulkWriter := firebase.FirestoreClient.BulkWriter(ctx)
	defer bulkWriter.End()

	for _, item := range qbItemsResponse.QueryResponse.Item {
		docRef := firebase.FirestoreClient.Collection(firebase.ProductsCollection).Doc(item.ID)
		bulkWriter.Set(docRef, item.MapToProduct().ToMap(), firestore.MergeAll)
	}

	bulkWriter.Flush()
	return nil
}