package model

type Customer struct {
	Name    string
	Remark  string
	Contact Contact
	Address Address
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
