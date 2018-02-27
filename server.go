package main

import (
	"net/http"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"echoplus/controllers"
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
			e.Logger.Printf("url >>>>>>>>>>  method %v",c.Request().Method)
			e.Logger.Printf("url >>>>>>>>>>  URL %v",c.Request().URL)
			e.Logger.Printf("URL >>>>>>>>>>> header  %v",c.Request().Header)
			e.Logger.Printf("URL >>>>>>>>>>> body  %v",c.Request().Body)
			e.Logger.Printf("URL >>>>>>>>>>> query string  %v",c.QueryString())
			return key == "valid-key", nil
		},
	}))
	api.GET("/users/:id", controllers.GetUser)

	e.Logger.Fatal(e.Start(":1323"))
}
