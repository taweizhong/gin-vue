package model

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name     string `gorm:"varchar(20);not null"`
	Iphone   string `gorm:"varchar(20);not null;unique"`
	Password string `gorm:"varchar(20);not null"`
}
