package models

import (
	"fmt"
	"strings"
	"time"

	"cloud.google.com/go/firestore"
)

type Product struct {
	ID            string    `json:"id" firestore:"id"`
	IsActive      bool      `json:"is_active" firestore:"isActive"`
	Brand         string    `json:"brand" firestore:"brand"`
	Name          string    `json:"name" firestore:"name"`
	SKU           string    `json:"sku" firestore:"sku"`
	Size          float64   `json:"size" firestore:"size"`
	SizeUnit      string    `json:"sizeUnit" firestore:"sizeUnit"`
	PackOf        int       `json:"packOf" firestore:"packOf"`
	Category      string    `json:"category" firestore:"category"`
	Price         float64   `json:"price" firestore:"price"` //This is only meant for record purposes. Actual prices are set by the super admins for each individual customer. So when an order is placed, the actual price of that product is stored in the cart item struct. This is only read once from `product_prices_per_customers` collection when an order is placed.
	PurchasePrice float64   `json:"purchasePrice" firestore:"purchasePrice"`
	Desc          string    `json:"desc" firestore:"desc"`
	Slug          string    `json:"slug" firestore:"slug"`
	NameKey       string    `json:"nameKey" firestore:"nameKey"`
	Stock         int       `json:"stock" firestore:"stock"`
	CreatedAt     time.Time `json:"created_at" firestore:"createdAt"`
	UpdatedAt     time.Time `json:"updated_at" firestore:"updatedAt"`
}

func (p *Product) ToMap() map[string]any {
	return map[string]any{
		"id":            p.ID,
		"isActive":      p.IsActive,
		"brand":         p.Brand,
		"name":          p.Name,
		"sku":           p.SKU,
		"size":          p.Size,
		"sizeUnit":      p.SizeUnit,
		"packOf":        p.PackOf,
		"category":      p.Category,
		"price":         p.Price,
		"purchasePrice": p.PurchasePrice,
		"desc":          p.Desc,
		"slug":          p.Slug,
		"nameKey":       p.NameKey,
		"createdAt":     firestore.ServerTimestamp,
		"updatedAt":     firestore.ServerTimestamp,
	}
}

/* Setters */

func (p *Product) SetID(id string) {
	p.ID = id
}
func (p *Product) SetIsActive(isActive bool) {
	p.IsActive = isActive
}
func (p *Product) SetBrand(brand string) {
	p.Brand = brand
}
func (p *Product) SetName(name string) {
	p.Name = name
}
func (p *Product) SetSKU(sku string) {
	p.SKU = sku
}
func (p *Product) SetSize(size float64) {
	p.Size = size
}
func (p *Product) SetSizeUnit(sizeUnit string) {
	p.SizeUnit = sizeUnit
}
func (p *Product) SetPackOf(packOf int) {
	p.PackOf = packOf
}
func (p *Product) SetCategory(category string) {
	p.Category = category
}
func (p *Product) SetPrice(price float64) {
	p.Price = price
}
func (p *Product) SetPurchasePrice(purchasePrice float64) {
	p.PurchasePrice = purchasePrice
}
func (p *Product) SetDesc(desc string) {
	p.Desc = desc
}
func (p *Product) SetSlug(slug string) {
	p.Slug = slug
}
func (p *Product) SetNameKey(nameKey string) {
	p.NameKey = nameKey
}
func (p *Product) SetCreatedAt(createdAt time.Time) {
	p.CreatedAt = createdAt
}
func (p *Product) SetUpdatedAt(updatedAt time.Time) {
	p.UpdatedAt = updatedAt
}

/* Getters */

// Purchase Price * Quantity
func (p *Product) GetTotalPurchasePrice(quantity uint16) float64 {
	return p.PurchasePrice * float64(quantity)
}
// (Selling Price - Purchase Price) * Quantity
func (p *Product) GetTotalRevenuePerProduct(sellingPrice float64, quantity uint16) float64 {
	return (sellingPrice - p.PurchasePrice) * float64(quantity)
}

func (p *Product) GetCorrectWeightInGallons(quantity uint16) float64 {
	unit := strings.ToUpper(p.SizeUnit)
	switch unit {
	case "OZ", "OUNCE", "OUNCES":
		return (p.Size / 128) * float64(quantity)
	case "LB", "LBS", "POUND", "POUNDS":
		return (p.Size * 0.125) * float64(quantity)
	case "QT", "QUART", "QUARTS":
		return (p.Size * 0.25) * float64(quantity)
	case "GAL", "GALLON", "GALLONS":
		return p.Size * float64(quantity)
	case "SHEETS": //Fabric softener sheets. Its a rough estimate
		return float64(quantity) * 0.15 / 128
	default:
		return p.Size * float64(quantity)
	}
}
func (p *Product) GetFormattedDescription() string {
	return fmt.Sprintf("%s - %s %.2f %s (Pack of %d)", p.Brand, p.Name, p.Size, strings.ToLower(p.SizeUnit), p.PackOf)
}
func (p *Product) GetShortDescription() string {
	return fmt.Sprintf("%s %s %.2f%s - %s", p.Brand, p.Name, p.Size, strings.ToLower(p.SizeUnit), p.SKU)
}
func (p *Product) GetFormattedPurchasePrice() string {
	return fmt.Sprintf("$%.2f", p.PurchasePrice)
}
func (p *Product) GetFormattedTotalPurchasePrice(quantity uint16) string {
	return fmt.Sprintf("$%.2f", p.GetTotalPurchasePrice(quantity))
}
func (p *Product) GetFormattedTotalRevenuePerProduct(sellingPrice float64, quantity uint16) string {
	return fmt.Sprintf("$%.2f", p.GetTotalRevenuePerProduct(sellingPrice, quantity))
}
func (p *Product) GetFormattedWeight(quantity uint16) string {
	return fmt.Sprintf("%.2f gal", p.GetCorrectWeightInGallons(quantity))
}