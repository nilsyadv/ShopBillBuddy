package config

import (
	"log"
	"net/http"

	"github.com/nilsyadv/ShopBillBuddy/common/pkg/config"
	wraperror "github.com/nilsyadv/ShopBillBuddy/common/pkg/error"
)

func InitConf(envname, envtype string) (config.InterfaceConfig, *wraperror.WrappedError) {
	log.Println(envname, envtype)
	conf, err := config.InitConfig(envname, envtype)
	if err != nil {
		werr := wraperror.Wrap(err, "error in new config initiation", "error", http.StatusInternalServerError)
		log.Println(err.Error())
		return nil, werr
	}

	return conf, nil
}
