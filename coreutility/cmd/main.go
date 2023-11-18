package main

import (
	"log"

	"github.com/gorilla/mux"
	"github.com/nilsyadv/ShopBillBuddy/common/pkg/logger"
	v1 "github.com/nilsyadv/ShopBillBuddy/coreutility/api/v1"
	"github.com/nilsyadv/ShopBillBuddy/coreutility/config"
	"github.com/nilsyadv/ShopBillBuddy/coreutility/internal/db"
	"github.com/nilsyadv/ShopBillBuddy/coreutility/server"
)

func main() {
	// Initialize application configuration
	config, configErr := config.InitConfig()
	if configErr != nil {
		log.Panic("Failed to initialize configuration:", configErr)
	}

	// Initialize logger based on configuration
	logger, loggerErr := logger.InitLogger(config.GetString("logger.level"), config.GetString("logger.output"))
	if loggerErr != nil {
		log.Panic("Failed to initialize logger:", loggerErr)
	}

	// Create a RedisDB object using the configuration
	redisDB := db.NewRedisObject(config)

	// Create a new mux.Router for routing HTTP requests
	router := mux.NewRouter()

	// Create and initialize the router for the Core Utility service (v1)
	coreUtilityRouter := v1.NewCoreUtilityRouter(redisDB, config, logger)
	coreUtilityRouter.InitRouter(router)

	// Create and initialize the HTTP server for the Core Utility service
	coreUtilityServer := server.NewCoreutilityServer(router, config, logger)
	coreUtilityServer.InitServer()
}
