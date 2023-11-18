package v1

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/nilsyadv/ShopBillBuddy/common/pkg/config"
	"github.com/nilsyadv/ShopBillBuddy/common/pkg/logger"
	"github.com/nilsyadv/ShopBillBuddy/coreutility/internal/controller"
	"github.com/nilsyadv/ShopBillBuddy/coreutility/internal/db"
	"github.com/nilsyadv/ShopBillBuddy/coreutility/internal/service"
)

// coreUtilityRouter handles routing for the Core Utility service.
type coreUtilityRouter struct {
	rds    *db.RedisDB
	config config.InterfaceConfig
	logger logger.InterfaceLogger
}

// NewRouter creates a new coreUtilityRouter with the provided dependencies.
func NewCoreUtilityRouter(rds *db.RedisDB, conf config.InterfaceConfig, lg logger.InterfaceLogger) *coreUtilityRouter {
	return &coreUtilityRouter{
		rds:    rds,
		config: conf,
		logger: lg,
	}
}

// InitRouter initializes the coreUtilityRouter by setting up its routes.
func (corutlyroutes *coreUtilityRouter) InitRouter(route *mux.Router) {
	// Number generator router initiation
	corutlyroutes.initNumGeneratorRouter(route)
}

// initNumGeneratorRouter initializes the router for the Number Generator service.
func (corutlyroutes *coreUtilityRouter) initNumGeneratorRouter(routr *mux.Router) *mux.Router {
	// Create instances of the service and controller with the provided dependencies
	numservice := service.NewNumGenService(corutlyroutes.rds, corutlyroutes.logger)
	numcontroller := controller.NewNumGen(numservice, corutlyroutes.logger)

	// Define routes for the Number Generator service
	routr.HandleFunc("/coreutility/nxtkey/{prefix}", numcontroller.NextKey).Methods(http.MethodGet)
	routr.HandleFunc("/coreutility/setkey/{prefix}/{length}/{intialvalue}", numcontroller.SetKey).Methods(http.MethodGet)

	return routr
}
