package libs

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"echoplus/config"
	"fmt"
)

func ConnectPG() (db *gorm.DB) {
	connectString := "host=" + config.AppConfig.GetString("db.pg_host") + " port=" + config.AppConfig.GetString("db.pg_port") + " user=" + config.AppConfig.GetString("db.pg_user") + " dbname=" + config.AppConfig.GetString("db.pg_dbname") + " password=" + config.AppConfig.GetString("db.pg_password") + " sslmode=disable"
	db, err := gorm.Open("postgres", connectString)
	if err != nil {
		panic(fmt.Errorf("Fatal err when db connect: %s \n", err))
	}
	db.LogMode(true)
	return
}
