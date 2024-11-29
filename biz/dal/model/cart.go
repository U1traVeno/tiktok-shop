package model

import "gorm.io/gorm"

type CartItems struct {
	// Your code here
	gorm.Model
	CartId    uint32
	ProductId uint32
	Quantity  uint32
}

type Cart struct {
	gorm.Model
	UserId uint32
	CartId uint32
}
