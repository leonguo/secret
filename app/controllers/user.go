package controllers

import (
	"github.com/labstack/echo"
	"net/http"
	"../models"
	"../db"
	"strconv"
	"../lib"
)

func GetUser(c echo.Context) error {
	// get params
	userId, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	conn := db.ConnectPG()
	defer db.ClosePG(conn)
	user := models.GetUserById(conn, userId)
	if user.Id == 0 {
		c.Logger().Debug("----------- %v ", user)
		return lib.Resp(c, http.StatusNotFound, "not found", "")
	}
	return c.JSON(http.StatusOK, user)
}

func AddUser(c echo.Context) error {
	// get params
	body := c.Request().Body
	conn := db.ConnectPG()
	defer db.ClosePG(conn)
	c.Logger().Printf("body : >>>> %v", body)
	var user models.User
	user.UserName = "guo2"
	user.Age = 23
	user = models.CreateUser(conn, user)
	// return user
	return c.JSON(http.StatusOK, user)
}
