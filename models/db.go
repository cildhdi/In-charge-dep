package models

import (
	"fmt"
	"github.com/cildhdi/In-charge/config"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var db *gorm.DB

func init() {
	var err error
	var dbCfg = config.IcCfg().DatabaseCfg
	db, err = gorm.Open(dbCfg.Name,
		fmt.Sprintf("host=%s port=%d user=%s dbname=%s password=%s",
			dbCfg.Host, dbCfg.Port, dbCfg.User, dbCfg.DbName, dbCfg.Password))
	if err != nil {
		panic(err)
	}
	//migrates
}

func IcDb() *gorm.DB {
	return db
}
