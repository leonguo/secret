package app

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/sevenNt/echo-pprof"
)

var Server *echo.Echo

func init() {
	Server = echo.New()
	Server.Debug = true
	// load default config
	Server.Use(middleware.Logger())
	Server.Use(middleware.Recover())
	Server.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET, echo.HEAD, echo.PUT, echo.PATCH, echo.POST, echo.DELETE},
	}))

	// route
	InitRoute()
	echopprof.Wrap(Server)
}