package app

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/spf13/viper"
)

var Server *echo.Echo

func Init() {
	Server = echo.New()
	// load default config
	Server.Use(
		func(next echo.HandlerFunc) echo.HandlerFunc {
			var config *viper.Viper
			LoadDefaultConfig(config)
			return func(c echo.Context) error {
			c.Set("Config", config)
			return next(c)
		}
	})
	Server.Use(middleware.Logger())
	Server.Use(middleware.Recover())
	Server.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET, echo.HEAD, echo.PUT, echo.PATCH, echo.POST, echo.DELETE},
	}))
}

func LoadDefaultConfig(config *viper.Viper) {
	config.Set("debug",true)
}