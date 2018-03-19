package main

import (
	"./app"
	"./config"
)

func main() {
	// 初始化app配置文件
	config.Init()
	app.Init()
	app.Server.Logger.Fatal(app.Server.Start(config.AppConfig.GetString("system.port")))
}
