package models

import (
	pgorm "../../db/gorm"
)

type User struct {
	Id        int64  `gorm:"primary_key" json:"id"`
	UserName  string `gorm:"column:username" json:"username"`
	Age       int    `gorm:"column:age" json:"age"`
	IsDeleted int    `gorm:"column:is_deleted" json:"is_deleted"`
}

func (User) TableName() string {
	return "public.users"
}

// 根据ID获取用户信息
func (u *User) GetUserById(userId int64) (err error) {
	err = pgorm.DBManager().First(&u, userId).Error
	return
}

func (u *User) CreateUser() {
	pgorm.DBManager().Create(&u)
	return
}

func (u *User) DeleteUser(userId int64) {
	u.Id = userId
	pgorm.DBManager().Model(u).Update("is_deleted", 1)
	return
}
