package main

import (
	"log"

	"github.com/gorilla/mux"
	"github.com/nilsyadv/ShopBillBuddy/common/pkg/logger"
	v1 "github.com/nilsyadv/ShopBillBuddy/product/api/v1"
	"github.com/nilsyadv/ShopBillBuddy/product/config"
	"github.com/nilsyadv/ShopBillBuddy/product/server"
)

func main() {
	// Initialize application configuration from a JSON file named "config.json"
	config, configErr := config.InitConf("config", "json")
	if configErr != nil {
		log.Fatal("Failed to initialize configuration:", configErr.Error())
	}

	// Initialize logger based on configuration
	logger, loggerErr := logger.InitLogger(config.GetString("logger.level"), config.GetString("logger.output"))
	if loggerErr != nil {
		log.Fatal("Failed to initialize logger:", loggerErr.Error())
	}

	// Create a new mux.Router for routing HTTP requests
	router := mux.NewRouter()

	// setup routes
	custroutes := v1.NewProductRouter(config, logger)
	custroutes.InitRouter(router)

	// Create and initialize the HTTP server for the Product service
	productServer := server.NewProductServer(router, config, logger)
	productServer.InitServer()
}
