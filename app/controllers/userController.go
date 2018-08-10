package controllers

import (
	"github.com/labstack/echo"
	"net/http"
	"../models"
	"github.com/davecgh/go-spew/spew"
	pgorm "../../db/gorm"
	"../lib"
	"secret/util"
)

func GetUser(c echo.Context) error {
	users := new(models.User)
	users.GetUserById(1)
	spew.Dump(users)
	test := make([]map[string]interface{}, 0)
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
	return c.JSON(http.StatusOK, m)
}

func GetUsers(c echo.Context) error {
	page := util.StringToInt(c.QueryParam("page"))
	limit := util.StringToInt(c.QueryParam("limit"))
	var users []models.User
	db := pgorm.DBManager()
	if page < 1 {
		page = 1
	}
	if limit < 1 {
		limit = 1
	}
	p := lib.Paginator{DB: db, Limit: limit, Page: page, OrderBy: []string{"id desc"}}
	data := p.Paginate(&users)
	return c.JSON(http.StatusOK, data)
}
