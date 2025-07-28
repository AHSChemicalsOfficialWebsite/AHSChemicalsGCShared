package models

import "time"

type CustomerPrices struct {
	ProductName string    `json:"product_name" firsetore:"product_name"`
	Brand       string    `json:"brand" firsetore:"brand"`
	ProductID   string    `json:"product_id" firsetore:"product_id"`
	CustomerID  string    `json:"customer_id" firsetore:"customer_id"`
	Price       float64   `json:"price" firsetore:"price"`
	CreatedAt   time.Time `json:"createdAt" firsetore:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt" firsetore:"updatedAt"`
} 