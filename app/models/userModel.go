package models

import (
	"echoplus/app/libs"
)

type User struct {
	Id       int64  `gorm:"primary_key"`
	UserName string `gorm:"column:username"`
	Age      int
}

func (User) TableName() string {
	return "public.user"
}

// 根据ID获取用户信息
func GetUserById() User {
	db := libs.ConnectPG()
	var user User
	db.First(&user)
	return user
}
