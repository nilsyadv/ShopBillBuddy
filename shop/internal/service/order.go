package service

import (
	wraperror "github.com/nilsyadv/ShopBillBuddy/common/pkg/error"
	"github.com/nilsyadv/ShopBillBuddy/common/pkg/logger"
	"github.com/nilsyadv/ShopBillBuddy/common/pkg/repository"
	"github.com/nilsyadv/ShopBillBuddy/shop/internal/model"
	"gorm.io/gorm"
)

// OrderService handles business logic related to Orders
type OrderService struct {
	logger *logger.InterfaceLogger
	repo   repository.Repository
	db     *gorm.DB
}

// NewOrderService creates a new instance of OrderService
func NewOrderService(logger *logger.InterfaceLogger, db *gorm.DB, repo repository.Repository) *OrderService {
	return &OrderService{
		logger: logger,
		db:     db,
		repo:   repo,
	}
}

// InsertNewOrder inserts a new Order into the database
func (custSer *OrderService) InsertNewOrder(Order *model.Order) *wraperror.WrappedError {
	uow := repository.NewUnitOfWork(custSer.db, false, *custSer.logger)
	return custSer.repo.Add(uow, Order)
}
