package model

import (
	"github.com/lib/pq"
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Name        string         `json:"name"`
	Amount      int            `json:"amount"`
	Description string         `json:"description"`
	Picture     string         `json:"picture"`
	Price       float32        `json:"price"`
	Categories  pq.StringArray `gorm:"type:text[]" json:"categories"`
}
