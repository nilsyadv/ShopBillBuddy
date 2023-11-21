package service

import (
	wraperror "github.com/nilsyadv/ShopBillBuddy/common/pkg/error"
	"github.com/nilsyadv/ShopBillBuddy/common/pkg/logger"
	"github.com/nilsyadv/ShopBillBuddy/common/pkg/repository"
	"github.com/nilsyadv/ShopBillBuddy/customer/internal/model"
	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

type customerService struct {
	logger *logger.InterfaceLogger
	repo   repository.Repository
	db     *gorm.DB
}

func NewCustomerService(logger *logger.InterfaceLogger, db *gorm.DB, repo repository.Repository) *customerService {
	return &customerService{
		logger: logger,
		db:     db,
		repo:   repo,
	}
}

func (custser *customerService) InsertNewCustomer(customer *model.Customer) *wraperror.WrappedError {
	uow := repository.NewUnitOfWork(custser.db, false, *custser.logger)
	return custser.repo.Add(uow, customer)
}

func (custser *customerService) UpdateCustomer(customer *model.Customer) *wraperror.WrappedError {
	uow := repository.NewUnitOfWork(custser.db, false, *custser.logger)
	return custser.repo.Update(uow, customer)
}

func (custser *customerService) DeleteCustomer(customer *model.Customer) *wraperror.WrappedError {
	uow := repository.NewUnitOfWork(custser.db, false, *custser.logger)
	return custser.repo.Delete(uow, customer, repository.Filter("id = ?", customer.ID))
}

func (custser *customerService) GetCustomer(customerID string, customer *model.Customer) *wraperror.WrappedError {
	uow := repository.NewUnitOfWork(custser.db, false, *custser.logger)
	return custser.repo.Get(uow, customer, uuid.UUID{})
}

func (custser *customerService) GetsCustomer() *wraperror.WrappedError {
	return nil
}
