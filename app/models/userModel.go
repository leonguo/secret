package models

import (
	"github.com/jinzhu/gorm"
)

type User struct {
	Id       int64  `gorm:"primary_key" json:"id"`
	UserName string `gorm:"column:username" json:"username"`
	Age      int    `json:"age"`
}

func (User) TableName() string {
	return "public.user"
}

// 根据ID获取用户信息
func GetUserById(db *gorm.DB) User {
	var user User
	db.First(&user)
	return user
}

func CreateUser() {
	//db := libs.ConnectPG()
	//user := &User{
	//	UserName:"ddd",
	//	Age:22,
	//}
	//db.Create(user)
}
