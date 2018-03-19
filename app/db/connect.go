package db

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"../../config"
	"fmt"
)

func ConnectPG() (PgDB *gorm.DB){
	connectString := "host=" + config.AppConfig.GetString("db.pg_host") + " port=" + config.AppConfig.GetString("db.pg_port") + " user=" + config.AppConfig.GetString("db.pg_user") + " dbname=" + config.AppConfig.GetString("db.pg_dbname") + " password=" + config.AppConfig.GetString("db.pg_password") + " sslmode=disable"
	PgDB, err := gorm.Open("postgres", connectString)
	if err != nil {
		panic(fmt.Errorf("Fatal err when db connect: %s \n", err))
	}
	PgDB.DB().SetMaxIdleConns(5)
	PgDB.DB().SetMaxOpenConns(100)
	PgDB.LogMode(true)
	return PgDB
}

func ClosePG(PgDB *gorm.DB) {
	err := PgDB.Close()
	if err != nil {
		panic(err)
	}
	return
}