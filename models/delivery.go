package models

import (
	"time"
)

type Delivery struct {
	Order          *Order
	ReceivedBy     string
	DeliveredBy    string
	Signature      []byte
	DeliveryImages [][]byte
	DeliveredAt    time.Time //In UTC
	Timezone       string
}