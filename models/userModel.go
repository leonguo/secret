package models

type User struct {
	Id   int64  `json:"id"`
	Name string `json:"user_name"`
	Age  int    `json:"age"`
}

type Users struct {
	Users []User `json:"users"`
}

// 根据ID获取用户信息
func GetUserById() User {
	user := User{1, "name", 2}
	return user
}
