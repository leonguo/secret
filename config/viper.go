package config

import (
	"github.com/spf13/viper"
	"fmt"
)

func InitConf() (Config *viper.Viper) {
	Config = viper.New()
	Config.SetConfigName("config")
	Config.SetConfigType("toml")
	Config.AddConfigPath("config")
	err := Config.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
	return
}
