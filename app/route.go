package app

import (
	"github.com/labstack/echo"
	"net/http"
	"./controllers"
	"../util"
	"github.com/labstack/echo/middleware"
	"github.com/pkg/errors"
)

func InitRoute() {
	Server.File("/static", "app/public/index.html")

	Server.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	// test
	Server.GET("/test", controllers.GetTest)

	Server.GET("/user", controllers.GetUser)
	Server.PUT("/user", controllers.PutUser)

	Server.POST("/getpng", controllers.GenerateCaptcha)
	Server.POST("/attr/update", controllers.AttrUpdate)
	Server.POST("/keepalive", controllers.KeepAlive)

	// api接口组
	accountApi := Server.Group("/v1/accounts")
	accountApi.GET("/:number", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	//验证header的参数
	authApi := Server.Group("/v1")
	authApi.Use(middleware.KeyAuthWithConfig(middleware.KeyAuthConfig{
		AuthScheme: "Basic",
		Validator: func(key string, c echo.Context) (bool, error) {
			checkOk, _ := util.AuthorizationHeader(c)
			return checkOk, errors.New("dd")
		},
	}))

}
