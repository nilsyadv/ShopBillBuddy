package config

import (
	"log"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

// InitConfig initializes a configuration using the Viper library in Go.
// It sets up a Viper instance, specifies the configuration file's path,
// name, and type based on input parameters, and attempts to read the
// configuration file. Error handling is implemented for various scenarios,
// including file not found, unsupported file type, and parsing/marshaling issues.
// Additionally, a callback is registered to log a message when the configuration
// file changes. The function returns an instance of the conf struct, which
// encapsulates the Viper configuration for further use.
func InitConfig(envname, envtype string) (InterfaceConfig, error) {
	vp := viper.New()
	vp.AddConfigPath(".")
	vp.AddConfigPath("/config/")
	vp.AddConfigPath("./config/")
	vp.SetConfigName(envname)
	vp.SetConfigType(envtype)
	if err := vp.ReadInConfig(); err != nil {
		switch err.(type) {
		case viper.ConfigFileNotFoundError:
			log.Println("configuration file not found on mentioned path '/config'")
			return nil, err
		case viper.UnsupportedConfigError:
			log.Println("unknown defined configuration file type")
			return nil, err
		case viper.ConfigMarshalError, viper.ConfigParseError:
			log.Println("something wrong in the configuration file: ", err.Error())
		default:
			log.Println("unknown error: ", err.Error())
			return nil, err
		}
	}

	vp.OnConfigChange(func(in fsnotify.Event) {
		log.Println("Configuration file changed")
	})

	return conf{
		cnf: vp,
	}, nil
}
