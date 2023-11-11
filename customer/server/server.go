package server

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/nilsyadv/ShopBillBuddy/common/pkg/config"
	"github.com/nilsyadv/ShopBillBuddy/common/pkg/logger"
)

type Server struct {
	conf   config.InterfaceConfig
	logger logger.InterfaceLogger
}

func (serv Server) InitServer() {

	routes := mux.NewRouter()

	server := http.Server{
		Addr:    serv.conf.GetString("server.host") + ":" + serv.conf.GetString("server.port"),
		Handler: routes,
	}

	serv.logger.Info("Customer Server started on " + server.Addr)

	if err := server.ListenAndServe(); err != nil {
		serv.logger.Fatal("fatal", err)
	}
}
