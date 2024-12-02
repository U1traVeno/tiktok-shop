// Code generated by hertz generator.

package cart

import (
	"context"
	"github.com/U1traVeno/tiktok-shop/biz/service"

	cart "github.com/U1traVeno/tiktok-shop/biz/model/cart"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
)

// AddItem .
// @router /cart/add [POST]
func AddItem(ctx context.Context, c *app.RequestContext) {
	var err error
	var req cart.AddItemReq
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	//add item
	_, err = service.NewCartService(ctx, c).AddItem(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	//response
	resp := new(cart.AddItemResp)
	resp.StatusCode = consts.StatusOK
	resp.StatusMsg = "成功加入购物车"

	c.JSON(consts.StatusOK, resp)
}

// GetCart .
// @router /cart [GET]
func GetCart(ctx context.Context, c *app.RequestContext) {
	var err error
	var req cart.GetCartReq
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	resp := new(cart.GetCartResp)

	c.JSON(consts.StatusOK, resp)
}

// EmptyCart .
// @router /cart/empty [DELETE]
func EmptyCart(ctx context.Context, c *app.RequestContext) {
	var err error
	var req cart.EmptyCartReq
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	//response
	resp := new(cart.EmptyCartResp)
	resp.StatusCode = consts.StatusOK
	resp.StatusMsg = "成功清空购物车"

	c.JSON(consts.StatusOK, resp)
}