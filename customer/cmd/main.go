package main

import (
	"log"

	"github.com/nilsyadv/ShopBillBuddy/common/pkg/logger"
	"github.com/nilsyadv/ShopBillBuddy/customer/config"
)

func main() {
	///----
	conf, werr := config.InitConf("config", "json")
	if werr != nil {
		log.Fatal(werr.Error())
	}

	custlog, err := logger.InitLogger(conf.GetString("logger.level"), conf.GetString("logger.output"))
	if err != nil {
		log.Fatal(err.Error())
	}

	custlog.Info("Customer server started")
}
