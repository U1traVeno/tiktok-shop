package model

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string `gorm:"unique"`
	Password string `gorm:"not null"`
	Role     string `gorm:"default:'user'"`
	Email    string `gorm:"unique"`
}
