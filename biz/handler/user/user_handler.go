package user

import (
	"context"
	"encoding/base64"
	"fmt"
	"github.com/U1traVeno/tiktok-shop/biz/dal/model"
	"strconv"
	"strings"
	"time"

	query "github.com/U1traVeno/tiktok-shop/biz/dal/query/user"
	"github.com/U1traVeno/tiktok-shop/biz/model/user"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"golang.org/x/crypto/bcrypt"
)

// User .
// @router /user/ [GET]
func User(ctx context.Context, c *app.RequestContext) {
	var req user.UserReq
	if err := c.BindAndValidate(&req); err != nil {
		handleError(c, consts.StatusBadRequest, err)
		return
	}

	userID, err := validateToken(req.Token)
	if err != nil {
		handleError(c, consts.StatusUnauthorized, err)
		return
	}

	userQuery := query.User
	u, err := userQuery.WithContext(ctx).Where(userQuery.ID.Eq(uint(userID))).First()
	if err != nil {
		handleError(c, consts.StatusNotFound, fmt.Errorf("user not found"))
		return
	}

	resp := new(user.UserResp)
	resp.StatusCode = consts.StatusOK
	resp.StatusMsg = consts.StatusMessage(consts.StatusOK)

	resp.Message = "Hello, " + u.Username

	c.JSON(consts.StatusOK, resp)
}

// UserRegister .
// @router /user/register [POST]
func UserRegister(ctx context.Context, c *app.RequestContext) {
	var req user.UserRegisterReq
	if err := c.BindAndValidate(&req); err != nil {
		handleError(c, consts.StatusBadRequest, err)
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		handleError(c, consts.StatusInternalServerError, err)
		return
	}

	newUser := model.User{
		Username: req.Username,
		Password: string(hashedPassword),
	}

	if err := query.User.Create(&newUser); err != nil {
		handleError(c, consts.StatusInternalServerError, err)
		return
	}

	resp := new(user.UserRegisterResp)
	resp.StatusCode = consts.StatusOK
	resp.StatusMsg = consts.StatusMessage(consts.StatusOK)

	resp.UserId = int64(newUser.ID)
	resp.Token = generateToken(int64(newUser.ID))

	c.JSON(consts.StatusOK, resp)
}

// UserLogin .
// @router /user/login [POST]
func UserLogin(ctx context.Context, c *app.RequestContext) {
	var req user.UserLoginReq
	if err := c.BindAndValidate(&req); err != nil {
		handleError(c, consts.StatusBadRequest, err)
		return
	}

	userQuery := query.User
	u, err := userQuery.WithContext(ctx).Where(userQuery.Username.Eq(req.Username)).First()
	if err != nil {
		handleError(c, consts.StatusNotFound, fmt.Errorf("user not found"))
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(req.Password)); err != nil {
		handleError(c, consts.StatusUnauthorized, fmt.Errorf("password error"))
		return
	}

	resp := new(user.UserLoginResp)
	resp.StatusCode = consts.StatusOK
	resp.StatusMsg = consts.StatusMessage(consts.StatusOK)

	resp.UserId = int64(u.ID)
	resp.Token = generateToken(int64(u.ID))

	if err := updateToken(ctx, int64(u.ID), resp.Token); err != nil {
		handleError(c, consts.StatusInternalServerError, err)
		return
	}

	c.JSON(consts.StatusOK, resp)
}

func handleError(c *app.RequestContext, statusCode int, err error) {
	resp := new(user.UserResp)
	resp.StatusCode = int32(statusCode)
	resp.StatusMsg = consts.StatusMessage(statusCode)
	resp.Message = err.Error()
	c.JSON(statusCode, resp)
}

func generateToken(userID int64) string {
	// TODO: generate token in database
	return generateTempToken(userID)
}

func validateToken(token string) (int64, error) {
	// TODO: validate token in database
	return validateTempToken(token)
}

func updateToken(ctx context.Context, userID int64, token string) error {
	// TODO: update token in database
	return nil
}

func generateTempToken(userID int64) string {
	data := fmt.Sprintf("%d:%d", userID, time.Now().Unix())
	token := base64.StdEncoding.EncodeToString([]byte(data))
	return token
}

func validateTempToken(token string) (int64, error) {
	data, err := base64.StdEncoding.DecodeString(token)
	if err != nil {
		return 0, fmt.Errorf("invalid token format")
	}

	parts := strings.Split(string(data), ":")
	if len(parts) != 2 {
		return 0, fmt.Errorf("invalid token data")
	}

	userID, err := strconv.ParseUint(parts[0], 10, 32)
	if err != nil {
		return 0, fmt.Errorf("invalid userID in token")
	}

	return int64(userID), nil
}
