package server

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/nilsyadv/ShopBillBuddy/common/pkg/config"
	"github.com/nilsyadv/ShopBillBuddy/common/pkg/logger"
)

// CustomerServer represents the HTTP server for the Core Utility application.
type CustomerServer struct {
	conf   config.InterfaceConfig
	logger logger.InterfaceLogger
	router *mux.Router
}

func NewCustomerServer(router *mux.Router, conf config.InterfaceConfig, log logger.InterfaceLogger) *CustomerServer {
	return &CustomerServer{
		conf:   conf,
		logger: log,
		router: router,
	}
}

// InitServer initializes the HTTP server with the provided router.
// It starts the server on the configured address and port.
func (srv CustomerServer) InitServer() {
	// Create an HTTP server with the specified address, port, and router
	server := http.Server{
		Addr:    srv.conf.GetString("app.addr") + ":" + srv.conf.GetString("app.port"),
		Handler: srv.router,
	}

	// Start the HTTP server and log any errors
	if err := server.ListenAndServe(); err != nil {
		srv.logger.Error("server shutting down", err)
	}
}
