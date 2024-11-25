// Code generated by hertz generator.

package auth

import (
	"context"

	auth "github.com/U1traVeno/tiktok-shop/biz/model/auth"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
)

// DeliverTokenByRPC .
// @router /auth/deliver [GET]
func DeliverTokenByRPC(ctx context.Context, c *app.RequestContext) {
	var err error
	var req auth.DeliverTokenReq
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	resp := new(auth.DeliveryResp)

	c.JSON(consts.StatusOK, resp)
}

// VerifyTokenByRPC .
// @router /auth/verify [POST]
func VerifyTokenByRPC(ctx context.Context, c *app.RequestContext) {
	var err error
	var req auth.VerifyTokenReq
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	resp := new(auth.VerifyResp)

	c.JSON(consts.StatusOK, resp)
}