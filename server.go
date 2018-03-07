package main

import (
	"echoplus/app"
	//"echoplus/config"
	"fmt"
	//"github.com/labstack/echo"
)

func main() {

	app.Init()
	fmt.Printf("%v","DD")
	app.Server.Logger.Fatal(app.Server.Start(":1323"))
}
