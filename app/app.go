package app

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"net/http"
	"echoplus/app/controllers"
)

var Server *echo.Echo

func Init() {
	Server = echo.New()
	Server.Debug = true
	// load default config
	Server.Use(middleware.Logger())
	Server.Use(middleware.Recover())
	Server.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET, echo.HEAD, echo.PUT, echo.PATCH, echo.POST, echo.DELETE},
	}))

	Server.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	// api接口组
	api := Server.Group("/api")
	Server.Use(middleware.RequestID())
	// 对API接口签名
	//api.Use(middleware.KeyAuthWithConfig(middleware.KeyAuthConfig{
	//	AuthScheme: "api_v1",
	//	Validator: func(key string, c echo.Context) (bool, error) {
	//		checkOk, err := util.SignAuth(c)
	//		return checkOk, err
	//	},
	//}))
	api.GET("/users/:id", controllers.GetUser)
	//api.POST("/users", controllers.AddUser)
}