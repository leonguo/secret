package app

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

var Server *echo.Echo

func Init() {
	Server = echo.New()
	Server.Debug = false
	// load default config
	//Server.Use(middleware.Logger())
	//Server.Use(middleware.Recover())
	Server.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET, echo.HEAD, echo.PUT, echo.PATCH, echo.POST, echo.DELETE},
	}))

	// route
	InitRoute()
}