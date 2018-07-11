package config

import (
	"github.com/spf13/viper"
	"fmt"
)

var AppConfig *viper.Viper

func init() {
	AppConfig = viper.New()
	AppConfig.SetConfigName("config")
	AppConfig.SetConfigType("toml")
	AppConfig.AddConfigPath("config")
	err := AppConfig.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
}