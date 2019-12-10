package models

import (
	"github.com/jinzhu/gorm"
)

const (
	MerChantUser uint = 1
	CustomerUser uint = 2
)

var UserType map[uint]string = map[uint]string{
	MerChantUser: "商家",
	CustomerUser: "顾客",
}

type IcUser struct {
	gorm.Model
	Phone  string `gorm:"not null;unique"`
	Name   string
	Points *int `gorm:"default:0"`
}
