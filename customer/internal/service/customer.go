package service

import (
	wraperror "github.com/nilsyadv/ShopBillBuddy/common/pkg/error"
	"github.com/nilsyadv/ShopBillBuddy/common/pkg/logger"
	"github.com/nilsyadv/ShopBillBuddy/common/pkg/repository"
	"github.com/nilsyadv/ShopBillBuddy/customer/internal/model"
)

type customerService struct {
	logger *logger.InterfaceLogger
	db     repository.Repository
}

func NewCustomerService(logger *logger.InterfaceLogger, db repository.Repository) *customerService {
	return &customerService{
		logger: logger,
		db:     db,
	}
}

func (custser *customerService) InsertNewCustomer(customer *model.Customer) *wraperror.WrappedError {

	return nil
}

func (custser *customerService) UpdateCustomer() *wraperror.WrappedError {
	return nil
}

func (custser *customerService) DeleteCustomer() *wraperror.WrappedError {
	return nil
}

func (custser *customerService) GetCustomer() *wraperror.WrappedError {
	return nil
}

func (custser *customerService) GetsCustomer() *wraperror.WrappedError {
	return nil
}
