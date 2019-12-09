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

}

func IcCfg() *Cfg {
	return &icCfg
}
