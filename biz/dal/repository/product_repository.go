package repository

import "github.com/U1traVeno/tiktok-shop/biz/model"

type ProductRepository interface {
	Create(product *model.Product) (*model.Product, error)
	Update(product *model.Product) (*model.Product, error)
	Delete(id uint32) error
	GetById(id uint32) (*model.Product, error)
	List() ([]*model.Product, error)
}
