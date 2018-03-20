package models

import (
	"github.com/jinzhu/gorm"
)

type User struct {
	Id        int64  `gorm:"primary_key" json:"id"`
	UserName  string `gorm:"column:username" json:"username"`
	Age       int    `json:"age"`
	IsDeleted int    `json:"is_deleted"`
}

func (User) TableName() string {
	return "public.user"
}

// 根据ID获取用户信息
func GetUserById(db *gorm.DB, userId int64) User {
	var user User
	db.First(&user, userId)
	return user
}

func CreateUser(db *gorm.DB, user User) User {
	db.Create(&user)
	return user
}

func DeleteUser(db *gorm.DB, user User) User {
	db.Model(&user).Update("is_deleted", 1)
	return user
}
