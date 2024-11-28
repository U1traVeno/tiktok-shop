package service

import (
	"github.com/U1traVeno/tiktok-shop/biz/dal/model"
)

type ProductService interface {
	CreateProduct(product *model.Product) (*model.Product, error)
	UpdateProduct(product *model.Product) (*model.Product, error)
	DeleteProduct(id uint32) error
	GetProduct(id uint32) (*model.Product, error)
	ListProducts() ([]*model.Product, error)
}
