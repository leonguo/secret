package controllers

import (
	"github.com/labstack/echo"
	"net/http"
	"echoplus/app/models"
	"echoplus/config"
)

func GetUser(c echo.Context) error {
	// get params
	userId := c.Param("id")
	c.Logger().Printf(userId)
	user := models.GetUserById()
	// return user
	c.Logger().Printf(" cfg >>>> %s",config.AppConfig.GetString("title"))
	return c.JSON(http.StatusOK, user)
}


func AddUser(c echo.Context) error {
	// get params
	body := c.Request().Body
	c.Logger().Printf("body : >>>> %v",body)
	user := models.GetUserById()
	// return user
	return c.JSON(http.StatusOK, user)
}
