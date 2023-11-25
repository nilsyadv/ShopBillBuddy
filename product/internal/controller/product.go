package controller

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/nilsyadv/ShopBillBuddy/common/pkg/https"
	"github.com/nilsyadv/ShopBillBuddy/common/pkg/logger"
	"github.com/nilsyadv/ShopBillBuddy/product/internal/model"
	"github.com/nilsyadv/ShopBillBuddy/product/internal/service"
)

// productController handles HTTP requests related to products
type productController struct {
	custSer *service.ProductService
	logger  logger.InterfaceLogger
}

// NewProductController creates a new instance of productController
func NewProductController(custSer *service.ProductService, logger logger.InterfaceLogger) *productController {
	return &productController{
		custSer: custSer,
		logger:  logger,
	}
}

// CreateProduct handles HTTP POST request to create a new product
func (custCtrl *productController) CreateProduct(w http.ResponseWriter, r *http.Request) {
	var product model.Product
	werr := https.RequestParse(r, &product)
	if werr != nil {
		https.RespondErrorMessage(w, werr.ErrCode, werr.Context)
		return
	}

	err := custCtrl.custSer.InsertNewProduct(&product)
	if err != nil {
		https.RespondErrorMessage(w, werr.ErrCode, werr.Context)
		return
	}

	https.RespondJSON(w, http.StatusCreated, &product)
}

// UpdateProduct handles HTTP PUT request to update an existing product
func (custCtrl *productController) UpdateProduct(w http.ResponseWriter, r *http.Request) {
	urlquery := mux.Vars(r)
	productID := urlquery["productid"]
	if productID == "" {
		https.RespondErrorMessage(w, http.StatusBadRequest, "Product ID cannot be empty")
		return
	}

	var product model.Product
	werr := https.RequestParse(r, &product)
	if werr != nil {
		https.RespondErrorMessage(w, werr.ErrCode, werr.Context)
		return
	}

	product.ID = productID

	werr = custCtrl.custSer.UpdateProduct(&product)
	if werr != nil {
		https.RespondErrorMessage(w, werr.ErrCode, werr.Context)
		custCtrl.logger.Error("Failed to update product:", werr)
		return
	}

	https.RespondJSON(w, http.StatusOK, &product)
}

// DeleteProduct handles HTTP DELETE request to delete an existing product
func (custCtrl *productController) DeleteProduct(w http.ResponseWriter, r *http.Request) {
	urlquery := mux.Vars(r)
	product := model.Product{
		ID: urlquery["productid"],
	}
	werr := custCtrl.custSer.DeleteProduct(&product)
	if werr != nil {
		https.RespondErrorMessage(w, werr.ErrCode, werr.Context)
		return
	}

	https.RespondJSON(w, http.StatusOK, map[string]string{"msg": "product deleted successfully"})
}

// GetProduct handles HTTP GET request to retrieve a product by ID
func (custCtrl *productController) GetProduct(w http.ResponseWriter, r *http.Request) {
	urlquery := mux.Vars(r)
	if urlquery["productid"] == "" {
		https.RespondErrorMessage(w, http.StatusBadRequest, "Product ID can not be empty")
		return
	}

	var product model.Product
	werr := custCtrl.custSer.GetProduct(urlquery["productid"], &product)
	if werr != nil {
		https.RespondErrorMessage(w, werr.ErrCode, werr.Context)
		return
	}

	https.RespondJSON(w, http.StatusOK, &product)
}

// GetProducts handles HTTP GET request to retrieve a list of products
func (custCtrl *productController) GetProducts(w http.ResponseWriter, r *http.Request) {
	var products []*model.Product
	werr := custCtrl.custSer.GetProducts(products)
	if werr != nil {
		https.RespondErrorMessage(w, werr.ErrCode, werr.Context)
		return
	}

	https.RespondJSON(w, http.StatusOK, &products)
}
