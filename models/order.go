package models

import (
	"fmt"
	"strings"
	"time"
)

type OrderDocument string

const (
	OrderDocumentPurchaseOrder    OrderDocument = "purchase_order"
	OrderDocumentShippingManifest OrderDocument = "shipping_manifest"
	OrderDocumentInvoice          OrderDocument = "invoice"
	OrderDocumentRevenueReport    OrderDocument = "revenue_report"
)

type OrderIDPaylod struct {
	OrderID string `json:"orderId"`
}

type OrderStatus string

const (
	OrderStatusPending   OrderStatus = "pending"
	OrderStatusApproved  OrderStatus = "approved"
	OrderStatusDelivered OrderStatus = "delivered"
	OrderStatusCancelled OrderStatus = "cancelled"
)

type OrderRequest struct {
	Items               []*OrderRequestItem `json:"items"`
	CustomerID          string              `json:"customerId"`
	SpecialInstructions string              `json:"specialInstructions"`
}

type OrderRequestItem struct {
	ProductID string `json:"productId"`
	Quantity  int    `json:"quantity"`
}

func (or *OrderRequest) ToProductIDs() []string{
	productIDs := make([]string, 0)
	for _, item := range or.Items {
		productIDs = append(productIDs, item.ProductID)
	}
	return productIDs
}

type Order struct {
	ID                  string      `json:"id" firestore:"-"`
	InvoiceNumber       string      `json:"invoiceNumber" firestore:"invoiceNumber,omitempty"`
	Customer            *Customer   `json:"customer" firestore:"customer"`
	Uid                 string      `json:"uid" firestore:"uid"`
	SpecialInstructions string      `json:"specialInstructions" firestore:"specialInstructions"`
	Items               []*OrderItem `json:"items" firestore:"items"`
	TaxRate             float64     `json:"taxRate" firestore:"taxRate"`
	TaxAmount           float64     `json:"taxAmount" firestore:"taxAmount"`
	SubTotal            float64     `json:"subTotal" firestore:"subTotal"`
	Total               float64     `json:"total" firestore:"total"`
	Status              OrderStatus `json:"status" firestore:"status"`
	CreatedAt           time.Time   `json:"createdAt" firestore:"createdAt"`
	UpdatedAt           time.Time   `json:"updatedAt" firestore:"updatedAt"`
}

type OrderItem struct {
	Quantity  int      `json:"quantity" firestore:"quantity"`
	ProductID string   `json:"productId" firestore:"productId"`
	Price     float64  `json:"price" firestore:"price"`
	Product   *Product `json:"product" firestore:"-"`
}

func (item *OrderItem) GetFormattedQuantity() string {
	return fmt.Sprintf("%d", item.Quantity)
}
func (item *OrderItem) GetFormattedPrice() string {
	return fmt.Sprintf("$%.2f", item.Price)
}
func (item *OrderItem) GetFormattedItemTotalPrice() string {
	return fmt.Sprintf("$%.2f", float64(item.Quantity)*item.Price)
}

func (o *Order) UpdateBill() {
	o.GetSubTotal()
	o.GetTaxAmount()
	o.GetTotal()
	o.SetUpdatedAt(time.Now().UTC())
}

func (o *Order) SetID(id string) {
	o.ID = id
}
func (o *Order) SetInvoiceNumber(invoiceNumber string) {
	o.InvoiceNumber = invoiceNumber
}
func (o *Order) SetCustomer(customer *Customer) {
	o.Customer = customer
}
func (o *Order) SetUID(uid string) {
	o.Uid = uid
}
func (o *Order) SetStatus(status OrderStatus) {
	o.Status = status
}
func (o *Order) SetTaxRate(taxRate float64) {
	o.TaxRate = taxRate
}
func (o *Order) SetItemPrices(correctPrices map[string]float64) {
	for i, item := range o.Items {
		o.Items[i].Price = correctPrices[item.Product.ID]
	}
}
func (o *Order) SetCreatedAt(createdAt time.Time) {
	o.CreatedAt = createdAt
}
func (o *Order) SetUpdatedAt(updatedAt time.Time) {
	o.UpdatedAt = updatedAt
}

func (o *Order) GetSubTotal() {
	subTotal := 0.0
	for _, item := range o.Items {
		subTotal += float64(item.Quantity) * item.Price
	}
	o.SubTotal = subTotal
}
func (o *Order) GetTaxAmount() {
	o.TaxAmount = o.TaxRate * o.SubTotal
}
func (o *Order) GetTotal() {
	o.Total = o.SubTotal + o.TaxAmount
}

// Total cost of goods (purchase price * quantity)
func (o *Order) GetTotalCOG() float64 {
	totalCOG := 0.0
	for _, item := range o.Items {
		totalCOG += item.Product.GetTotalPurchasePrice(item.Quantity)
	}
	return totalCOG
}
func (o *Order) GetFormattedTotal() string {
	return fmt.Sprintf("$%.2f", o.Total)
}
func (o *Order) GetFormattedSubTotal() string {
	return fmt.Sprintf("$%.2f", o.SubTotal)
}
func (o *Order) GetFormattedTaxAmount() string {
	return fmt.Sprintf("$%.2f", o.TaxAmount)
}
func (o *Order) GetFormattedTaxRate() string {
	return fmt.Sprintf("%.2f%%", o.TaxRate*100)
}
func (o *Order) GetFormattedTotalItems() string {
	var totalUnits int = 0
	for _, item := range o.Items {
		totalUnits += item.Quantity
	}
	return fmt.Sprintf("%d", totalUnits)
}
func (o *Order) GetFormattedNetWeight() string {
	weight := 0.0
	for _, item := range o.Items {
		weight += item.Product.GetCorrectWeightInGallons(item.Quantity)
	}
	return fmt.Sprintf("%.2f gal", weight)
}

func (o *Order) GetFormattedCOG() string {
	return fmt.Sprintf("$%.2f", o.GetTotalCOG())
}

// Subtotal - COG. Does not include tax as per the client requirement
func (o *Order) GetFormattedTotalRevenue() string {
	return fmt.Sprintf("$%.2f", o.SubTotal-o.GetTotalCOG())
}

func (o *Order) ToMap() map[string]any {
	return map[string]any{
		"customerId":          o.Customer.ID,
		"customerName":        strings.ToLower(o.Customer.Name),
		"uid":                 o.Uid,
		"specialInstructions": o.SpecialInstructions,
		"items":               o.Items,
		"taxRate":             o.TaxRate,
		"taxAmount":           o.TaxAmount,
		"subTotal":            o.SubTotal,
		"total":               o.Total,
		"status":              o.Status,
		"createdAt":           o.CreatedAt,
		"updatedAt":           o.UpdatedAt,
	}
}

// Returns an array of product IDs. Comes in handy for bulk firestore operations
func (o *Order) ToProductIDs() []string {
	productIDs := make([]string, 0)
	for _, item := range o.Items {
		productIDs = append(productIDs, item.Product.ID)
	}
	return productIDs
}

// Converts the order items array to a map of ID:Product
func (o *Order) ToItemMap() map[string]*Product {
	idMap := make(map[string]*Product)
	for _, item := range o.Items {
		idMap[item.Product.ID] = item.Product
	}
	return idMap
}

func (order *Order) CreateItemsDataForAdminEmail() []map[string]any {
	orderItems := make([]map[string]any, 0)
	for _, item := range order.Items {
		mappedItem := make(map[string]any)
		mappedItem["sku"] = item.Product.SKU
		mappedItem["description"] = item.Product.GetFormattedDescription()
		mappedItem["quantity"] = item.GetFormattedQuantity()
		mappedItem["price"] = item.GetFormattedPrice()

		orderItems = append(orderItems, mappedItem)
	}
	return orderItems
}

func (order *Order) CreateItemsDetailedItemsDataForAdminEmail() []map[string]any {
	orderItems := make([]map[string]any, 0)
	for _, item := range order.Items {
		mappedItem := make(map[string]any)
		mappedItem["sku"] = item.Product.SKU
		mappedItem["description"] = item.Product.GetFormattedDescription()
		mappedItem["quantity"] = item.GetFormattedQuantity()
		mappedItem["unit_price"] = item.GetFormattedPrice()
		mappedItem["total_price"] = item.GetFormattedItemTotalPrice()

		orderItems = append(orderItems, mappedItem)
	}
	return orderItems
}

func (order *Order) CreateItemsDataForUserEmail() []map[string]any {
	orderItems := make([]map[string]any, 0)
	for _, item := range order.Items {
		mappedItem := make(map[string]any)
		mappedItem["sku"] = item.Product.SKU
		mappedItem["description"] = item.Product.GetFormattedDescription()
		mappedItem["quantity"] = item.GetFormattedQuantity()

		orderItems = append(orderItems, mappedItem)
	}
	return orderItems
}
