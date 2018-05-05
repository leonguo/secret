package app

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"net/http"
	"./controllers"
)

var Server *echo.Echo

func Init() {
	Server = echo.New()
	Server.Debug = false
	Server.DisableHTTP2 = true
	// load default config
	Server.Use(middleware.Logger())
	Server.Use(middleware.Recover())
	Server.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET, echo.HEAD, echo.PUT, echo.PATCH, echo.POST, echo.DELETE},
	}))

	Server.File("/static", "app/public/index.html")

	Server.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	Server.GET("/user",controllers.GetUser)
	Server.POST("/getpng",controllers.GenerateCaptcha)
	Server.POST("/attr/update",controllers.AttrUpdate)
	Server.POST("/keepalive",controllers.KeepAlive)

}