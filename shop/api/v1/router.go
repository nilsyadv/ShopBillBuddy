package v1

import (
	"net/http"

	"github.com/gorilla/mux"
	"gorm.io/gorm"

	"github.com/nilsyadv/ShopBillBuddy/common/pkg/config"
	"github.com/nilsyadv/ShopBillBuddy/common/pkg/logger"
	"github.com/nilsyadv/ShopBillBuddy/common/pkg/repository"
	"github.com/nilsyadv/ShopBillBuddy/shop/internal/controller"
	"github.com/nilsyadv/ShopBillBuddy/shop/internal/service"
)

// OrderRouter handles routing for the Core Utility service.
type orderRouter struct {
	config config.InterfaceConfig
	logger logger.InterfaceLogger
	db     *gorm.DB
	repo   repository.Repository
}

// NewRouter creates a new OrderRouter with the provided dependencies.
func NewOrderRouter(conf config.InterfaceConfig, lg logger.InterfaceLogger) *orderRouter {
	return &orderRouter{
		// rds:    rds,
		config: conf,
		logger: lg,
		// db
	}
}

// InitRouter initializes the OrderRouter by setting up its routes.
func (custroutes *orderRouter) InitRouter(route *mux.Router) {
	// Order router initiation
	custroutes.initOrderRouter(route)
}

// InitOrderRouter initializes the router for the Order service.
func (custroutes *orderRouter) initOrderRouter(routr *mux.Router) *mux.Router {
	// Create instances of the service and controller with the provided dependencies
	custService := service.NewOrderService(&custroutes.logger, custroutes.db, custroutes.repo)
	custController := controller.NewOrderController(custService, custroutes.logger)

	// Define routes for the order service
	routr.HandleFunc("/shop", custController.CreateOrder).Methods(http.MethodPost)

	return routr
}
