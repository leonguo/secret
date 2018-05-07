package controllers

import (
	"github.com/labstack/echo"
	"net/http"
	"../models"
	"github.com/davecgh/go-spew/spew"
)

func GetUser(c echo.Context) error {
	users := new(models.User)
	users.GetUserById(1)
	spew.Dump(users)
	return c.JSON(http.StatusOK, users)
}
