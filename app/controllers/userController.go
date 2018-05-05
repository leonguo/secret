package controllers

import (
	"github.com/labstack/echo"
	"net/http"

)

func GetUser(c echo.Context) error {
	// get params
	return c.JSON(http.StatusOK, "")
}