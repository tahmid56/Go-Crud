package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name string
	Password string
	Phone string
	Email string `gorm:"unique"`
	Post []Post `gorm:"foreignKey:Title;references:Email"`
}