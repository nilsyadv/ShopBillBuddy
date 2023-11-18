package main

import (
	"log"

	"github.com/gorilla/mux"
	"github.com/nilsyadv/ShopBillBuddy/common/pkg/logger"
	"github.com/nilsyadv/ShopBillBuddy/customer/config"
	"github.com/nilsyadv/ShopBillBuddy/customer/server"
)

func main() {
	// Initialize application configuration from a JSON file named "config.json"
	config, configErr := config.InitConf("customer", "json")
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

	// Create and initialize the HTTP server for the Customer service
	customerServer := server.NewCustomerServer(router, config, logger)
	customerServer.InitServer()
}
