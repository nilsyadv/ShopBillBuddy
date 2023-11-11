package config

import (
	"github.com/spf13/viper"
)

type InterfaceConfig interface {
	GetString(string) string
	GetInt(string) int
}

type conf struct {
	cnf *viper.Viper
}

func (cf conf) GetString(key string) string {
	return cf.cnf.GetString(key)
}

func (cf conf) GetInt(key string) int {
	return cf.cnf.GetInt(key)
}
