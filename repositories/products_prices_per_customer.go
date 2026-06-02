package repositories

import (
	"context"
	"fmt"

	firebase "github.com/AHSChemicalsOfficialWebsite/AHSChemicalsGCShared/firebase"
	"github.com/AHSChemicalsOfficialWebsite/AHSChemicalsGCShared/models"
)

func SaveProductsPricesPerCustomerToFirestore(ctx context.Context) error {
	customers, err := FetchAllCustomersFromFirestore(ctx)
	if err != nil {
		return err
	}
	products, err := FetchAllProductsFromFirestore(ctx)
	if err != nil {
		return err
	}

	bulkWriter := firebase.FirestoreClient.BulkWriter(ctx)

	//Just blindly updating all prices per customers for each product since we first need to make
	//a read request to firestore to get all customers and products then check if prices have changed
	//or not which is way more reads than just directly bulk updating. In addition to this, the amount
	//of customers and products are not that much so it really does not matter.
	for _, product := range products {
		for _, customer := range customers {
			key := fmt.Sprintf("%s|%s", product.ID, customer.ID)
			productToAdd := models.CreateProductPricePerCustomer(product, customer.ID).ToMap()
			_, err := bulkWriter.Set(firebase.FirestoreClient.Collection(firebase.ProductsPricesPerCustCollection).Doc(key), productToAdd)
			if err != nil {
				return err
			}
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
func GetProductPricesFromCustomerID(customerID string, ctx context.Context) (map[string]float64, error) {

	docs, err := firebase.FirestoreClient.Collection(firebase.ProductsPricesPerCustCollection).Where("customerId", "==", customerID).Documents(ctx).GetAll()
	if err != nil {
		return nil, err
	}
	pricesMap := make(map[string]float64) //Map of product id to price
	for _, doc := range docs {
		var pppc models.ProductPricePerCustomer
		err := doc.DataTo(&pppc)
		if err != nil {
			return nil, err
		}
		pricesMap[pppc.ProductID] = pppc.Price
	}
	return pricesMap, nil
}
