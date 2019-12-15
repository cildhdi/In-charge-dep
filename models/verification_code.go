package models

import (
	"github.com/jinzhu/gorm"
)

type VerificationCode struct {
	gorm.Model
	Phone string
	Code  uint
}
