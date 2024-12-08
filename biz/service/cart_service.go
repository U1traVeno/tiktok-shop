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
	//检查数据是否合法
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

func (s *CartService) GetCart(userid uint64) ([]*cart.CartItem, error) {
	//检查userid是否合法
	if userid <= 0 {
		return nil, fmt.Errorf("invalid user ID")
	}

	cartTable := query.Cart

	//查找用户购物车中的商品
	cartItems, err := cartTable.Where(cartTable.UserId.Eq(userid)).Find()
	if err != nil {
		return nil, fmt.Errorf("failed to find cart items: %w", err)
	}

	//将查找到的商品转换为cartItem
	var items []*cart.CartItem
	for _, item := range cartItems {
		items = append(items, &cart.CartItem{
			ProductId: item.ProductID,
			Quantity:  item.Quantity,
		})
	}

	return items, nil
}

func (s *CartService) Empty(userid uint64) error {
	//检查userid是否合法
	if userid <= 0 {
		return fmt.Errorf("invalid userid")
	}

	cartTable := query.Cart

	//检查是否购物车中是否有商品
	_, err := cartTable.Where(cartTable.UserId.Eq(userid)).First()
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return fmt.Errorf("cart is empty")
		}
		return fmt.Errorf("failed to find cart item: %w", err)
	}

	//删除购物车中商品
	_, err = cartTable.Where(cartTable.UserId.Eq(userid)).Delete()

	if err != nil {
		return fmt.Errorf("failed to delete cart item: %w", err)
	}

	return nil
}
