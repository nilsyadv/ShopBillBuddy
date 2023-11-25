package model

import "time"

// Bill represents a bill entity
type Bill struct {
	ID          string    `json:"id"`
	CustomerID  string    `json:"customer_id"`
	OrderID     string    `json:"order_id"`
	Amount      float64   `json:"amount"`
	PaymentDate time.Time `json:"payment_date"`
}

// NewBill creates a new instance of Bill
func NewBill(id, customerID, orderID string, amount float64, paymentDate time.Time) *Bill {
	return &Bill{
		ID:          id,
		CustomerID:  customerID,
		OrderID:     orderID,
		Amount:      amount,
		PaymentDate: paymentDate,
	}
}
