package v1

import (
	"github.com/gorilla/mux"

	"github.com/nilsyadv/ShopBillBuddy/common/pkg/config"
	"github.com/nilsyadv/ShopBillBuddy/common/pkg/logger"
)

// CustomerRouter handles routing for the Core Utility service.
type customerRouter struct {
	// rds    *db.RedisDB
	config config.InterfaceConfig
	logger logger.InterfaceLogger
}

// NewRouter creates a new CustomerRouter with the provided dependencies.
func NewCustomerRouter(conf config.InterfaceConfig, lg logger.InterfaceLogger) *customerRouter {
	return &customerRouter{
		// rds:    rds,
		config: conf,
		logger: lg,
	}
}

// InitRouter initializes the CustomerRouter by setting up its routes.
func (corutlyroutes *customerRouter) InitRouter(route *mux.Router) {
	// Number generator router initiation
	// corutlyroutes.initNumGeneratorRouter(route)
}

// initNumGeneratorRouter initializes the router for the Number Generator service.
// func (corutlyroutes *customerRouter) InitNumGeneratorRouter(routr *mux.Router) *mux.Router {
// Create instances of the service and controller with the provided dependencies
// numservice := service.NewNumGenService(corutlyroutes.rds, corutlyroutes.logger)
// numcontroller := controller.NewNumGen(numservice, corutlyroutes.logger)

// // Define routes for the Number Generator service
// routr.HandleFunc("/Customer/nxtkey/{prefix}", numcontroller.NextKey).Methods(http.MethodGet)
// routr.HandleFunc("/Customer/setkey/{prefix}/{length}/{intialvalue}", numcontroller.SetKey).Methods(http.MethodGet)

// 	return routr
// }
