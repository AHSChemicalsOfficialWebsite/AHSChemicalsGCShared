package models

import "fmt"

type CartItem struct {
	Quantity  int      `json:"quantity" firestore:"quantity"`
	ProductID string   `json:"productId" firestore:"productId"`
	Price     float64  `json:"price" firestore:"price"`
	Product   *Product `json:"product" firestore:"-"`
}

func (item *CartItem) GetFormattedQuantity() string {
	return fmt.Sprintf("%d", item.Quantity)
}
func (item *CartItem) GetFormattedPrice() string {
	return fmt.Sprintf("$%.2f", item.Price)
}
func (item *CartItem) GetFormattedItemTotalPrice() string {
	return fmt.Sprintf("$%.2f", float64(item.Quantity)*item.Price)
}

func AreEqualPrices(a, b []*CartItem) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i].Price != b[i].Price {
			return false
		}
	}
	return true
}

func AreEqualQuantities(a, b []*CartItem) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i].Quantity != b[i].Quantity {
			return false
		}
	}
	return true
}
