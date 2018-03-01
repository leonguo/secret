package main

import (
	"net/http"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"echoplus/controllers"
	"echoplus/util"
)

func main() {
	e := echo.New()
	//e.Use(middleware.Secure())
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	// api接口组
	api := e.Group("/api")
	api.Use(middleware.KeyAuthWithConfig(middleware.KeyAuthConfig{
		AuthScheme: "api_v1",
		Validator: func(key string, c echo.Context) (bool, error) {
			util.SignAuth(c)
			return key == "valid-key", nil
		},
	}))
	api.GET("/users/:id", controllers.GetUser)

	e.Logger.Fatal(e.Start(":1323"))
}
