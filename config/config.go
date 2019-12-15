package config

import (
	"fmt"
	"github.com/olebedev/config"
	"io/ioutil"
)

type Cfg struct {
	DatabaseCfg struct {
		Name     string
		Host     string
		Port     int
		User     string
		DbName   string
		Password string
		SSLMode  string
	}

	AdminCfg struct {
		Phone string
	}

	Auth struct {
		SecretKey string
	}
}

var icCfg Cfg

func init() {
	file, err := ioutil.ReadFile("config.yaml")
	if err != nil {
		fmt.Println("read config from config.yaml failed, using default config")
	}
	cfg, err := config.ParseYaml(string(file))

	//database
	icCfg.DatabaseCfg.Name = cfg.UString("db.name", "postgres")
	icCfg.DatabaseCfg.Host = cfg.UString("db.host", "127.0.0.1")
	icCfg.DatabaseCfg.Port = cfg.UInt("db.port", 5432)
	icCfg.DatabaseCfg.User = cfg.UString("db.user", "ic")
	icCfg.DatabaseCfg.DbName = cfg.UString("db.db_name", "ic")
	icCfg.DatabaseCfg.Password = cfg.UString("db.password", "ic_password")
	icCfg.DatabaseCfg.SSLMode = cfg.UString("db.sslmode", "disable")

	//admin
	icCfg.AdminCfg.Phone = cfg.UString("admin.phone", "12345678900")

	//auth
	icCfg.Auth.SecretKey = cfg.UString("auth.secret", "ic_secret_key")
}

func IcCfg() *Cfg {
	return &icCfg
}
