package service

import (
	wraperror "github.com/nilsyadv/ShopBillBuddy/common/pkg/error"
	"github.com/nilsyadv/ShopBillBuddy/common/pkg/logger"
	"github.com/nilsyadv/ShopBillBuddy/common/pkg/repository"
	"github.com/nilsyadv/ShopBillBuddy/customer/internal/model"
	"gorm.io/gorm"
)

// CustomerService handles business logic related to customers
type CustomerService struct {
	logger *logger.InterfaceLogger
	repo   repository.Repository
	db     *gorm.DB
}

// NewCustomerService creates a new instance of CustomerService
func NewCustomerService(logger *logger.InterfaceLogger, db *gorm.DB, repo repository.Repository) *CustomerService {
	return &CustomerService{
		logger: logger,
		db:     db,
		repo:   repo,
	}
}

// InsertNewCustomer inserts a new customer into the database
func (custSer *CustomerService) InsertNewCustomer(customer *model.Customer) *wraperror.WrappedError {
	uow := repository.NewUnitOfWork(custSer.db, false, *custSer.logger)
	return custSer.repo.Add(uow, customer)
}

// UpdateCustomer updates an existing customer in the database
func (custSer *CustomerService) UpdateCustomer(customer *model.Customer) *wraperror.WrappedError {
	uow := repository.NewUnitOfWork(custSer.db, false, *custSer.logger)
	return custSer.repo.Update(uow, customer)
}

// DeleteCustomer deletes an existing customer from the database
func (custSer *CustomerService) DeleteCustomer(customer *model.Customer) *wraperror.WrappedError {
	uow := repository.NewUnitOfWork(custSer.db, false, *custSer.logger)
	return custSer.repo.Delete(uow, customer, repository.Filter("id = ?", customer.ID))
}

// GetCustomer retrieves a customer by ID from the database
func (custSer *CustomerService) GetCustomer(customerID string, customer *model.Customer) *wraperror.WrappedError {
	uow := repository.NewUnitOfWork(custSer.db, false, *custSer.logger)
	return custSer.repo.Get(uow, customer, customerID, []string{})
}

// GetsCustomer retrieves all customers from the database
func (custSer *CustomerService) GetCustomers(customers []*model.Customer) *wraperror.WrappedError {
	uow := repository.NewUnitOfWork(custSer.db, false, *custSer.logger)
	return custSer.repo.GetAll(uow, customers, []repository.QueryProcessor{})
}
