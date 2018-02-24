package controllers

import (
	"github.com/labstack/echo"
	"net/http"
	"echoplus/models"
)

func GetUser(c echo.Context) error {
	user := models.GetUserById()
	return c.JSON(http.StatusOK, user)
}
