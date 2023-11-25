package service

import (
	wraperror "github.com/nilsyadv/ShopBillBuddy/common/pkg/error"
	"github.com/nilsyadv/ShopBillBuddy/common/pkg/logger"
	"github.com/nilsyadv/ShopBillBuddy/common/pkg/repository"
	"github.com/nilsyadv/ShopBillBuddy/product/internal/model"
	"gorm.io/gorm"
)

// ProductService handles business logic related to products
type ProductService struct {
	logger *logger.InterfaceLogger
	repo   repository.Repository
	db     *gorm.DB
}

// NewProductService creates a new instance of ProductService
func NewProductService(logger *logger.InterfaceLogger, db *gorm.DB, repo repository.Repository) *ProductService {
	return &ProductService{
		logger: logger,
		db:     db,
		repo:   repo,
	}
}

// InsertNewProduct inserts a new product into the database
func (custSer *ProductService) InsertNewProduct(product *model.Product) *wraperror.WrappedError {
	uow := repository.NewUnitOfWork(custSer.db, false, *custSer.logger)
	return custSer.repo.Add(uow, product)
}

// UpdateProduct updates an existing product in the database
func (custSer *ProductService) UpdateProduct(product *model.Product) *wraperror.WrappedError {
	uow := repository.NewUnitOfWork(custSer.db, false, *custSer.logger)
	return custSer.repo.Update(uow, product)
}

// DeleteProduct deletes an existing product from the database
func (custSer *ProductService) DeleteProduct(product *model.Product) *wraperror.WrappedError {
	uow := repository.NewUnitOfWork(custSer.db, false, *custSer.logger)
	return custSer.repo.Delete(uow, product, repository.Filter("id = ?", product.ID))
}

// GetProduct retrieves a product by ID from the database
func (custSer *ProductService) GetProduct(productID string, product *model.Product) *wraperror.WrappedError {
	uow := repository.NewUnitOfWork(custSer.db, false, *custSer.logger)
	return custSer.repo.Get(uow, product, productID, []string{})
}

// GetsProduct retrieves all products from the database
func (custSer *ProductService) GetProducts(products []*model.Product) *wraperror.WrappedError {
	uow := repository.NewUnitOfWork(custSer.db, false, *custSer.logger)
	return custSer.repo.GetAll(uow, products, []repository.QueryProcessor{})
}
