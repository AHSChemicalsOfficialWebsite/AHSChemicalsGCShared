package repositories

import (
	"context"
	"fmt"

	"cloud.google.com/go/firestore"
	firebase_shared "github.com/HarshMohanSason/AHSChemicalsGCShared/shared/firebase"
	"github.com/HarshMohanSason/AHSChemicalsGCShared/shared/models"
	"github.com/HarshMohanSason/AHSChemicalsGCShared/shared/quickbooks/qbmodels"
)

//FetchAllProductsByIDs fetches all products from firestore from collection ('products')
//Returns map of product id and the product object in order to search in O(1)
//
//Params:
//  - ctx: context
//  - productIDs: []string, product ids
//
//Returns:
//  - map[string]models.Product. Key is product id
//  - error
func FetchAllProductsByIDs(ctx context.Context, productIDs []string) (map[string]models.Product, error) {
	docRefs := make([]*firestore.DocumentRef, len(productIDs))
	for i, productID := range productIDs {
		docRefs[i] = firebase_shared.FirestoreClient.Collection("products").Doc(productID)
	}
	docSnapshots, err := firebase_shared.FirestoreClient.GetAll(ctx, docRefs)
	if err != nil{
		return nil, err
	}

	productMap := make(map[string]models.Product)
	for i, docSnapshot := range docSnapshots {
		var product models.Product
		if err := docSnapshot.DataTo(&product); err != nil {
			return nil, fmt.Errorf("error decoding product %s: %v", productIDs[i], err)
		}
		productMap[product.ID] = product
	}
	return productMap, nil
}

// FetchAllProductsFromFirestore fetches all products from firestore from collection ('products')
func FetchAllProductsFromFirestore(ctx context.Context) ([]models.Product, error) {

	products, err := firebase_shared.FirestoreClient.Collection("products").Where("isActive", "==", true).Documents(ctx).GetAll()
	if err != nil {
		return nil, err
	}

	productList := make([]models.Product, len(products))
	for i, product := range products {
		var item models.Product
		if err := product.DataTo(&item); err != nil {
			return nil, fmt.Errorf("error decoding product %s: %v", product.Ref.ID, err)
		}
		productList[i] = item
	}
	return productList, nil
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

	bulkWriter := firebase_shared.FirestoreClient.BulkWriter(ctx)

	for _, item := range qbItemsResponse.QueryResponse.Item {
		docRef := firebase_shared.FirestoreClient.Collection("products").Doc(item.ID)
		_, err := bulkWriter.Set(docRef, item.MapToProduct().ToMap(), firestore.MergeAll)
		if err != nil {
			return err
		}
	}
	bulkWriter.Flush()
	return nil
}