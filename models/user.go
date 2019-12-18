package models

import (
	"github.com/jinzhu/gorm"
)

const (
	CustomerUser = int(iota)
	MerchantUser
	AdminUser
	SuperUser
)

var UserType map[int]string = map[int]string{
	SuperUser:    "超级用户",
	AdminUser:    "管理员",
	MerchantUser: "商家",
	CustomerUser: "顾客",
}

type IcUser struct {
	gorm.Model
	Phone     string `gorm:"not null;unique"`
	Name      string
	Role      int `gorm:"default:0"`
	Banned    int `gorm:"default:0"`
	Points    int `gorm:"default:0"`
	AvatarUrl string
}
