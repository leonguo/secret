package libs

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"echoplus/config"
)

func ConnectPG() (db *gorm.DB) {
	CONNECT := "host=" + config.AppConfig.GetString("db.postgre_host") + " port=" + config.AppConfig.GetString("db.postgre_port") + " user= " + " dbname=" + config.AppConfig.GetString("db.postgre_dbname") + " password=" + config.AppConfig.GetString("db.postgre_password")
	db, err := gorm.Open("postgres", CONNECT)
	if err != nil {
		panic(err)
	}
	defer db.Close()
	return
}
