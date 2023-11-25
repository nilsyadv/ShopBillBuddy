package config

import (
	"log"
	"net/http"

	"github.com/nilsyadv/ShopBillBuddy/common/pkg/config"
	wraperror "github.com/nilsyadv/ShopBillBuddy/common/pkg/error"
)

func InitConfig() (config.InterfaceConfig, *wraperror.WrappedError) {
	conf, err := config.InitConfig("config", "json")
	if err != nil {
		werr := wraperror.Wrap(err, "encounter error during config setup", "error", http.StatusInternalServerError)
		log.Println(err)
		return nil, werr
	}

	return conf, nil
}
