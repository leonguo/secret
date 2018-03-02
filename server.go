package main

import (
	"net/http"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"echoplus/controllers"
	//"echoplus/util"
)

func main() {
	e := echo.New()
	//e.Use(middleware.Secure())
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.BodyDump(func(c echo.Context, reqBody, resBody []byte) {
	}))

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	// api接口组
	api := e.Group("/api")
	e.Use(middleware.RequestID())
	// 对API接口签名
	//api.Use(middleware.KeyAuthWithConfig(middleware.KeyAuthConfig{
	//	AuthScheme: "api_v1",
	//	Validator: func(key string, c echo.Context) (bool, error) {
	//		checkOk, err := util.SignAuth(c)
	//		return checkOk, err
	//	},
	//}))
	api.GET("/users/:id", controllers.GetUser)
	api.POST("/users", controllers.AddUser)

	e.Logger.Fatal(e.Start(":1323"))
}
