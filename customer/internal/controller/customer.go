package controller

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/nilsyadv/ShopBillBuddy/common/pkg/https"
	"github.com/nilsyadv/ShopBillBuddy/common/pkg/logger"
	"github.com/nilsyadv/ShopBillBuddy/customer/internal/model"
	"github.com/nilsyadv/ShopBillBuddy/customer/internal/service"
)

// customerController handles HTTP requests related to customers
type customerController struct {
	custSer *service.CustomerService
	logger  logger.InterfaceLogger
}

// NewCustomerController creates a new instance of customerController
func NewCustomerController(custSer *service.CustomerService, logger logger.InterfaceLogger) *customerController {
	return &customerController{
		custSer: custSer,
		logger:  logger,
	}
}

// CreateCustomer handles HTTP POST request to create a new customer
func (custCtrl *customerController) CreateCustomer(w http.ResponseWriter, r *http.Request) {
	var customer model.Customer
	werr := https.RequestParse(r, &customer)
	if werr != nil {
		https.RespondErrorMessage(w, werr.ErrCode, werr.Context)
		return
	}

	err := custCtrl.custSer.InsertNewCustomer(&customer)
	if err != nil {
		https.RespondErrorMessage(w, werr.ErrCode, werr.Context)
		return
	}

	https.RespondJSON(w, http.StatusCreated, &customer)
}

// UpdateCustomer handles HTTP PUT request to update an existing customer
func (custCtrl *customerController) UpdateCustomer(w http.ResponseWriter, r *http.Request) {
	urlquery := mux.Vars(r)
	customerID := urlquery["customerid"]
	if customerID == "" {
		https.RespondErrorMessage(w, http.StatusBadRequest, "Customer ID cannot be empty")
		return
	}

	var customer model.Customer
	werr := https.RequestParse(r, &customer)
	if werr != nil {
		https.RespondErrorMessage(w, werr.ErrCode, werr.Context)
		return
	}

	customer.ID = customerID

	werr = custCtrl.custSer.UpdateCustomer(&customer)
	if werr != nil {
		https.RespondErrorMessage(w, werr.ErrCode, werr.Context)
		custCtrl.logger.Error("Failed to update customer:", werr)
		return
	}

	https.RespondJSON(w, http.StatusOK, &customer)
}

// DeleteCustomer handles HTTP DELETE request to delete an existing customer
func (custCtrl *customerController) DeleteCustomer(w http.ResponseWriter, r *http.Request) {
	urlquery := mux.Vars(r)
	customer := model.Customer{
		ID: urlquery["customerid"],
	}
	werr := custCtrl.custSer.DeleteCustomer(&customer)
	if werr != nil {
		https.RespondErrorMessage(w, werr.ErrCode, werr.Context)
		return
	}

	https.RespondJSON(w, http.StatusOK, map[string]string{"msg": "customer deleted successfully"})
}

// GetCustomer handles HTTP GET request to retrieve a customer by ID
func (custCtrl *customerController) GetCustomer(w http.ResponseWriter, r *http.Request) {
	urlquery := mux.Vars(r)
	if urlquery["customerid"] == "" {
		https.RespondErrorMessage(w, http.StatusBadRequest, "Customer ID can not be empty")
		return
	}

	var customer model.Customer
	werr := custCtrl.custSer.GetCustomer(urlquery["customerid"], &customer)
	if werr != nil {
		https.RespondErrorMessage(w, werr.ErrCode, werr.Context)
		return
	}

	https.RespondJSON(w, http.StatusOK, &customer)
}

// GetCustomers handles HTTP GET request to retrieve a list of customers
func (custCtrl *customerController) GetCustomers(w http.ResponseWriter, r *http.Request) {
	var customers []*model.Customer
	werr := custCtrl.custSer.GetCustomers(customers)
	if werr != nil {
		https.RespondErrorMessage(w, werr.ErrCode, werr.Context)
		return
	}

	https.RespondJSON(w, http.StatusOK, &customers)
}
