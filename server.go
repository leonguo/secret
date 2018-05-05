package main

import (
	"./app"
	"./config"
	"./db/gorm"
	"./db/redis"
)

func main() {
	// 初始化app配置文件
	config.Init()
	app.Init()
	redis.Init()
	gorm.PostgresConn()
	app.Server.Logger.Fatal(app.Server.Start(config.AppConfig.GetString("system.port")))
}
