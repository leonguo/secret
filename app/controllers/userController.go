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

// 更新用户信息
func PutUser(c echo.Context) error {
	m := echo.Map{}
	if err := c.Bind(&m); err != nil {
		return err
	}
	spew.Dump(m)
	return c.JSON(http.StatusOK, "")
}

func GetTest(c echo.Context) error {
	m := map[string]interface{}{"dd": "dd"}
	return c.JSON(http.StatusOK,m)
}
