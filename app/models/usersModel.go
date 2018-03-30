package models

import (
	"github.com/jinzhu/gorm"
)

type Users struct {
	Id        int64  `gorm:"primary_key" json:"id"`
	UserName  string `gorm:"column:username" json:"username"`
	Age       int    `gorm:"column:age" json:"age"`
	IsDeleted int    `gorm:"column:is_deleted" json:"is_deleted"`
}

func (Users) TableName() string {
	return "public.users"
}

// 根据ID获取用户信息
func (u *Users) GetUserById(db *gorm.DB, userId int64) {
	db.First(u, userId)
	return
}

func (u *Users) CreateUser(db *gorm.DB) {
	db.Create(u)
	return
}

func (u *Users) DeleteUser(db *gorm.DB, userId int64) {
	u.Id = userId
	db.Model(u).Update("is_deleted", 1)
	return
}
