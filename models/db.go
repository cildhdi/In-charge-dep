package models

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"

	"github.com/cildhdi/In-charge/config"
)

var db *gorm.DB

func init() {
	var err error
	var dbCfg = config.IcCfg().DatabaseCfg
	db, err = gorm.Open(dbCfg.Name,
		fmt.Sprintf("host=%s port=%d user=%s dbname=%s password=%s sslmode=%s",
			dbCfg.Host, dbCfg.Port, dbCfg.User, dbCfg.DbName, dbCfg.Password, dbCfg.SSLMode))
	if err != nil {
		panic(err)
	}
	//migrates

	db.AutoMigrate(&IcUser{})
	db.AutoMigrate(&VerificationCode{})

}

func IcDb() *gorm.DB {
	return db
}
