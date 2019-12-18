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

	adminCfg := config.IcCfg().AdminCfg
	count := 0
	if err := db.Table("ic_users").Count(&count).Error; err != nil {
		panic(err)
	}
	if count == 0 {
		su := IcUser{
			Phone: adminCfg.Phone,
			Role:  SuperUser,
		}
		if err := db.Create(&su).Error; err != nil {
			panic(err)
		}
	}
}

func IcDb() *gorm.DB {
	return db
}
