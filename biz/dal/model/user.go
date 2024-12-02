package model

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string `gorm:"unique"`
	Password string `gorm:"not null"`
	Role     string `gorm:"default:'user'"`
	Email    string `gorm:"unique;default:null"` // 如果不设置default:null，会默认为"", 与unique冲突
}
