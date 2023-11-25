package server

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/nilsyadv/ShopBillBuddy/common/pkg/config"
	"github.com/nilsyadv/ShopBillBuddy/common/pkg/logger"
)

// ShopServer represents the HTTP server for the Core Utility application.
type shopServer struct {
	conf   config.InterfaceConfig
	logger logger.InterfaceLogger
	router *mux.Router
}

func NewShopServer(router *mux.Router, conf config.InterfaceConfig, log logger.InterfaceLogger) *shopServer {
	return &shopServer{
		conf:   conf,
		logger: log,
		router: router,
	}
}

// InitServer initializes the HTTP server with the provided router.
// It starts the server on the configured address and port.
func (srv shopServer) InitServer() {
	// Create an HTTP server with the specified address, port, and router
	server := http.Server{
		Addr:    srv.conf.GetString("app.addr") + ":" + srv.conf.GetString("app.port"),
		Handler: srv.router,
	}

	srv.logger.Infof("Shop running on %s", srv.conf.GetString("app.addr")+":"+srv.conf.GetString("app.port"))

	// Start the HTTP server and log any errors
	if err := server.ListenAndServe(); err != nil {
		srv.logger.Error("server shutting down", err)
	}
}
