package controller

import (
	"net/http"

	"github.com/nilsyadv/ShopBillBuddy/common/pkg/https"
	"github.com/nilsyadv/ShopBillBuddy/common/pkg/logger"
	"github.com/nilsyadv/ShopBillBuddy/shop/internal/model"
	"github.com/nilsyadv/ShopBillBuddy/shop/internal/service"
)

// orderController handles HTTP requests related to orders
type orderController struct {
	custSer *service.OrderService
	logger  logger.InterfaceLogger
}

// NewOrderController creates a new instance of orderController
func NewOrderController(custSer *service.OrderService, logger logger.InterfaceLogger) *orderController {
	return &orderController{
		custSer: custSer,
		logger:  logger,
	}
}

// CreateOrder handles HTTP POST request to create a new order
func (custCtrl *orderController) CreateOrder(w http.ResponseWriter, r *http.Request) {
	var order model.Order
	werr := https.RequestParse(r, &order)
	if werr != nil {
		https.RespondErrorMessage(w, werr.ErrCode, werr.Context)
		return
	}

	err := custCtrl.custSer.InsertNewOrder(&order)
	if err != nil {
		https.RespondErrorMessage(w, werr.ErrCode, werr.Context)
		return
	}

	https.RespondJSON(w, http.StatusCreated, &order)
}
