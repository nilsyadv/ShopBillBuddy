package model

import "time"

// Order represents an order entity
type Order struct {
	ID            string         `json:"id"`
	CustomerID    string         `json:"customer_id"`
	ProductID     string         `json:"product_id"`
	Quantity      int            `json:"quantity"`
	TotalAmount   float64        `json:"total_amount"`
	PaidAmount    float64        `json:"paid_amount"`
	PendingAmount float64        `json:"pending_amount"`
	Orders        []OrderProduct `json:"orders"`
	CreatedAt     time.Time      `json:"created_at"`
}

// NewOrder creates a new instance of Order
func NewOrder(id, customerID, productID string, quantity int, totalAmount, paidAmount float64, createdAt time.Time) *Order {
	return &Order{
		ID:            id,
		CustomerID:    customerID,
		ProductID:     productID,
		Quantity:      quantity,
		TotalAmount:   totalAmount,
		PaidAmount:    paidAmount,
		PendingAmount: totalAmount - paidAmount,
		CreatedAt:     createdAt,
	}
}

type OrderProduct struct {
	ProductID        string
	ProductName      string
	Quantity         int
	PricePerQuantity float32
}

func (order *Order) UpdateBill() {
	var totalAmount float64
	for _, product := range order.Orders {
		totalAmount += float64(product.PricePerQuantity)
	}

	order.TotalAmount = totalAmount
	order.PendingAmount = order.TotalAmount - order.PaidAmount
}
