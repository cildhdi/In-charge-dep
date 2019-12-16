package models

import (
	"github.com/jinzhu/gorm"
)

const (
	CustomerUser = int(iota)
	MerChantUser
	AdminUser
	SuperUser
)

var UserType map[int]string = map[int]string{
	SuperUser:    "超级用户",
	AdminUser:    "管理员",
	MerChantUser: "商家",
	CustomerUser: "顾客",
}

func CreateRole(userType int) *int {
	role := userType
	return &role
}

type IcUser struct {
	gorm.Model
	Phone     string `gorm:"not null;unique"`
	Name      string
	Role      *int `gorm:"default:0;not null"`
	Banned    *int `gorm:"default:0;not null"`
	Points    *int `gorm:"default:0"`
	AvatarUrl string
}
