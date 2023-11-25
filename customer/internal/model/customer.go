package model

import "encoding/json"

type Customer struct {
	ID      string
	Name    string
	Remark  string
	Contact json.RawMessage
	Address json.RawMessage
}

type Contact struct {
	MobileNo    string
	Email       string
	AlternateNo string
}

type Address struct {
	ID         string
	Address    string
	Zipcode    string
	Landmark   string
	CustomerID string
}
