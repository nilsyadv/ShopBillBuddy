package v1

import (
	"net/http"

	"github.com/gorilla/mux"
	"gorm.io/gorm"

	"github.com/nilsyadv/ShopBillBuddy/common/pkg/config"
	"github.com/nilsyadv/ShopBillBuddy/common/pkg/logger"
	"github.com/nilsyadv/ShopBillBuddy/common/pkg/repository"
	"github.com/nilsyadv/ShopBillBuddy/product/internal/controller"
	"github.com/nilsyadv/ShopBillBuddy/product/internal/service"
)

// ProductRouter handles routing for the Core Utility service.
type productRouter struct {
	config config.InterfaceConfig
	logger logger.InterfaceLogger
	db     *gorm.DB
	repo   repository.Repository
}

// NewRouter creates a new ProductRouter with the provided dependencies.
func NewProductRouter(conf config.InterfaceConfig, lg logger.InterfaceLogger) *productRouter {
	return &productRouter{
		// rds:    rds,
		config: conf,
		logger: lg,
		// db
	}
}

// InitRouter initializes the ProductRouter by setting up its routes.
func (custroutes *productRouter) InitRouter(route *mux.Router) {
	// Product router initiation
	custroutes.initProductRouter(route)
}

// InitProductRouter initializes the router for the Product service.
func (custroutes *productRouter) initProductRouter(routr *mux.Router) *mux.Router {
	// Create instances of the service and controller with the provided dependencies
	custService := service.NewProductService(&custroutes.logger, custroutes.db, custroutes.repo)
	custController := controller.NewProductController(custService, custroutes.logger)

	// Define routes for the product service
	routr.HandleFunc("/products", custController.GetProducts).Methods(http.MethodGet)
	routr.HandleFunc("/product", custController.CreateProduct).Methods(http.MethodPost)
	routr.HandleFunc("/product/{productid}", custController.GetProduct).Methods(http.MethodGet)
	routr.HandleFunc("/product/{productid}", custController.UpdateProduct).Methods(http.MethodPut)
	routr.HandleFunc("/product/{productid}", custController.DeleteProduct).Methods(http.MethodDelete)

	return routr
}
