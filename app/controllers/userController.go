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
	test := make([]map[string]interface{},0)
	m := map[string]interface{}{"dd": "dd"}
	test = append(test, m)
	spew.Dump(test)
	return c.JSON(http.StatusOK, users)
}
