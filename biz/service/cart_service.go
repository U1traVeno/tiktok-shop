package service

import (
	"context"
	"fmt"
	"github.com/U1traVeno/tiktok-shop/biz/dal/model"
	query "github.com/U1traVeno/tiktok-shop/biz/dal/query/cart"
	"github.com/U1traVeno/tiktok-shop/biz/model/cart"
	"github.com/cloudwego/hertz/pkg/app"
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

func (s *CartService) AddItem(req *cart.AddItemReq) (*cart.AddItemResp, error) {
	//判断商品数量
	if req.Item.Quantity <= 0 {
		return nil, fmt.Errorf("invalid quantity")
	}

	//提供商品信息
	item := &model.Cart{
		UserId:    req.UserId,
		ProductID: req.Item.ProductId,
		Quantity:  req.Item.Quantity,
	}

	//向数据库添加购物车商品信息
	cartTable := query.Cart
	_ = cartTable.Create(item)

	return nil, nil
}

func (s *CartService) GetCart(req *cart.GetCartReq) (*cart.GetCartResp, error) {
	return nil, nil
}

func (s *CartService) Empty(req *cart.EmptyCartReq) (*cart.EmptyCartResp, error) {
	return nil, nil
}
