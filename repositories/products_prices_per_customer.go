package repositories

import (
	"context"
	"fmt"

	firebase "github.com/AHSChemicalsOfficialWebsite/AHSChemicalsGCShared/firebase"
	"github.com/AHSChemicalsOfficialWebsite/AHSChemicalsGCShared/models"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func CreateProductPricesForNewCustomer(ctx context.Context, customer *models.Customer) error {
    products, err := FetchAllActiveProductsFromFirestore(ctx)
    if err != nil {
        return err
    }
    bulkWriter := firebase.FirestoreClient.BulkWriter(ctx)
    for _, product := range products {
        key := fmt.Sprintf("%s|%s", product.ID, customer.ID)
        ref := firebase.FirestoreClient.Collection(firebase.ProductsPricesPerCustCollection).Doc(key)
        _, err := bulkWriter.Create(ref, models.CreateProductPricePerCustomer(product, customer.ID).ToMap())
        if err != nil {
            return err
        }
    }
    bulkWriter.Flush()
    return nil
}

func CreateProductPricesForNewProduct(ctx context.Context, product *models.Product) error {
    customers, err := FetchAllActiveCustomersFromFirestore(ctx)
	if err != nil {
		return err
	}
	bulkWriter := firebase.FirestoreClient.BulkWriter(ctx)
	for _, customer := range customers {
		key := fmt.Sprintf("%s|%s", product.ID, customer.ID)
		ref := firebase.FirestoreClient.Collection(firebase.ProductsPricesPerCustCollection).Doc(key)
		_, err := bulkWriter.Create(ref, models.CreateProductPricePerCustomer(product, customer.ID).ToMap())
		if err != nil {
			return err
		}
	}
	bulkWriter.Flush()
	return nil
}

func EnsureProductPricesForProduct(ctx context.Context, product *models.Product) error {
	customers, err := FetchAllActiveCustomersFromFirestore(ctx)
	if err != nil {
		return err
	}

	bulkWriter := firebase.FirestoreClient.BulkWriter(ctx)
	for _, customer := range customers {
		key := fmt.Sprintf("%s|%s", product.ID, customer.ID)
		ref := firebase.FirestoreClient.Collection(firebase.ProductsPricesPerCustCollection).Doc(key)

		_, err := ref.Get(ctx)
		if err == nil {
			continue // price doc already exists continue to avoid overwriting super admin written price
		}
		if status.Code(err) != codes.NotFound {
			return err 
		}

		_, err = bulkWriter.Create(ref, models.CreateProductPricePerCustomer(product, customer.ID).ToMap())
		if err != nil {
			return err
		}
	}
	bulkWriter.Flush()
	return nil
}

func EnsureProductPricesForCustomer(ctx context.Context, customer *models.Customer) error {
	products, err := FetchAllActiveProductsFromFirestore(ctx)
	if err != nil {
		return err
	}

	bulkWriter := firebase.FirestoreClient.BulkWriter(ctx)
	for _, product := range products {
		key := fmt.Sprintf("%s|%s", product.ID, customer.ID)
		ref := firebase.FirestoreClient.Collection(firebase.ProductsPricesPerCustCollection).Doc(key)

		_, err := ref.Get(ctx)
		if err == nil {
			continue
		}
		if status.Code(err) != codes.NotFound {
			return err
		}

		_, err = bulkWriter.Create(ref, models.CreateProductPricePerCustomer(product, customer.ID).ToMap())
		if err != nil {
			return err
		}
	}
	bulkWriter.Flush()
	return nil
}

// GetProductPricesFromCustomerID returns a map of product id to price.
// Returning a map is done to avoid multiple calls to firestore and to search in O(1)
//
// Params:
//   - customerID: ID of the customer
//   - ctx: context
//
// Returns:
//   - map of product id to price
//   - error
func GetProductPricesFromCustomerID(ctx context.Context, customerID string, productIDs []string) (map[string]float64, error) {
    docs, err := firebase.FirestoreClient.Collection(firebase.ProductsPricesPerCustCollection).
        Where("customerId", "==", customerID).
        Where("productId", "in", productIDs).
        Documents(ctx).GetAll()
    if err != nil {
        return nil, err
    }
    pricesMap := make(map[string]float64)
    for _, doc := range docs {
        var pppc models.ProductPricePerCustomer
        if err := doc.DataTo(&pppc); err != nil {
            return nil, err
        }
        pricesMap[pppc.ProductID] = pppc.Price
    }
    return pricesMap, nil
}