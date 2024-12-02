package model

import "gorm.io/gorm"

type Cart struct {
	// Your code here
	gorm.Model
	UserId    uint64 `json:"user_id"`
	ProductID uint64 `json:"product_id"`
	Quantity  uint64 `json:"quantity"`
}
