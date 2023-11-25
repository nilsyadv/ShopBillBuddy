package v1

import (
	"net/http"

	"github.com/gorilla/mux"
	"gorm.io/gorm"

	"github.com/nilsyadv/ShopBillBuddy/common/pkg/config"
	"github.com/nilsyadv/ShopBillBuddy/common/pkg/logger"
	"github.com/nilsyadv/ShopBillBuddy/common/pkg/repository"
	"github.com/nilsyadv/ShopBillBuddy/customer/internal/controller"
	"github.com/nilsyadv/ShopBillBuddy/customer/internal/service"
)

// CustomerRouter handles routing for the Core Utility service.
type customerRouter struct {
	config config.InterfaceConfig
	logger logger.InterfaceLogger
	db     *gorm.DB
	repo   repository.Repository
}

// NewRouter creates a new CustomerRouter with the provided dependencies.
func NewCustomerRouter(conf config.InterfaceConfig, lg logger.InterfaceLogger) *customerRouter {
	return &customerRouter{
		// rds:    rds,
		config: conf,
		logger: lg,
		// db
	}
}

// InitRouter initializes the CustomerRouter by setting up its routes.
func (custroutes *customerRouter) InitRouter(route *mux.Router) {
	// Customer router initiation
	custroutes.initCustomerRouter(route)
}

// InitCustomerRouter initializes the router for the Customer service.
func (custroutes *customerRouter) initCustomerRouter(routr *mux.Router) *mux.Router {
	// Create instances of the service and controller with the provided dependencies
	custService := service.NewCustomerService(&custroutes.logger, custroutes.db, custroutes.repo)
	custController := controller.NewCustomerController(custService, custroutes.logger)

	// Define routes for the customer service
	routr.HandleFunc("/customers", custController.GetCustomers).Methods(http.MethodGet)
	routr.HandleFunc("/customer", custController.CreateCustomer).Methods(http.MethodPost)
	routr.HandleFunc("/customer/{customerid}", custController.GetCustomer).Methods(http.MethodGet)
	routr.HandleFunc("/customer/{customerid}", custController.UpdateCustomer).Methods(http.MethodPut)
	routr.HandleFunc("/customer/{customerid}", custController.DeleteCustomer).Methods(http.MethodDelete)

	return routr
}
