package gorm

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"../../config"
	"fmt"
)

var db *gorm.DB

func PostgresConn(){
	connectString := "host=" + config.AppConfig.GetString("db.pg_host") + " port=" + config.AppConfig.GetString("db.pg_port") + " user=" + config.AppConfig.GetString("db.pg_user") + " dbname=" + config.AppConfig.GetString("db.pg_dbname") + " password=" + config.AppConfig.GetString("db.pg_password") + " sslmode=disable"
	db, err := gorm.Open("postgres", connectString)
	if err != nil {
		panic(fmt.Errorf("Fatal err when postgre db connect: %s \n", err))
	}
	db.DB().SetMaxIdleConns(5)
	db.DB().SetMaxOpenConns(100)
	db.LogMode(true)
	return
}

func ClosePg() {
	err := db.Close()
	if err != nil {
		panic(err)
	}
	return
}

func DBManager() *gorm.DB {
	return db
}