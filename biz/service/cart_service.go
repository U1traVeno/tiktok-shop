package service

import (
	"context"
	"errors"
	"fmt"
	"github.com/U1traVeno/tiktok-shop/biz/dal/model"
	query "github.com/U1traVeno/tiktok-shop/biz/dal/query/cart"
	"github.com/U1traVeno/tiktok-shop/biz/model/cart"
	"github.com/cloudwego/hertz/pkg/app"
	"gorm.io/gorm"
)

type CartService struct {
	ctx context.Context
	c   *app.RequestContext
}

func NewCartService(ctx context.Context, c *app.RequestContext) *CartService {
	return &CartService{
		ctx: ctx,
		c:   c,
	}
}

func (s *CartService) AddItem(req *cart.AddItemReq) error {
	//判断数据合法性
	if req.Item.Quantity <= 0 || req.Item.ProductId <= 0 || req.UserId <= 0 {
		return fmt.Errorf("invalid data")
	}

	cartTable := query.Cart

	//查找用户购物车是否已存在该商品
	item, err := cartTable.Where(cartTable.UserId.Eq(req.UserId), cartTable.ProductID.Eq(req.Item.ProductId)).First()
	if err != nil {
		//未找到则向数据库添加该商品数据
		if errors.Is(err, gorm.ErrRecordNotFound) {
			item := &model.Cart{
				UserId:    req.UserId,
				ProductID: req.Item.ProductId,
				Quantity:  req.Item.Quantity,
			}
			err = cartTable.Create(item)
			if err != nil {
				return fmt.Errorf("failed to create cart item: %w", err)
			}
			return nil
		}
		//其他报错返回错误
		return fmt.Errorf("failed to find cart item: %w", err)
	}
	//更新商品数量
	_, err = cartTable.Where(cartTable.UserId.Eq(req.UserId), cartTable.ProductID.Eq(req.Item.ProductId)).Update(cartTable.Quantity, req.Item.Quantity+item.Quantity)
	if err != nil {
		return fmt.Errorf("failed to update cart item: %w", err)
	}
	return nil
}

func (s *CartService) GetCart(req *cart.GetCartReq) (*cart.GetCartResp, error) {
	return nil, nil
}

func (s *CartService) Empty(req *cart.EmptyCartReq) (*cart.EmptyCartResp, error) {
	return nil, nil
}
