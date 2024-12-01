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

// User handles getting user information
// @Summary Get user information
// @Description Get user information by user ID and token
// @Tags User
// @Produce json
// @Param id query int true "Request body with user ID and token"
// @Success 200 {object} user.UserResp
// @Router /user/ [get]
func User(ctx context.Context, c *app.RequestContext) {
	var req user.UserReq
	if err := c.BindAndValidate(&req); err != nil {
		handleError(c, consts.StatusBadRequest, err)
		return
	}

	// validate token
	userID, err := validateToken(req.Token)
	if err != nil {
		handleError(c, consts.StatusUnauthorized, err)
		return
	}

	// get user
	userQuery := query.User
	u, err := userQuery.WithContext(ctx).Where(userQuery.ID.Eq(uint(userID))).First()
	if err != nil {
		handleError(c, consts.StatusNotFound, fmt.Errorf("user not found"))
		return
	}

	// response
	resp := new(user.UserResp)
	resp.StatusCode = consts.StatusOK
	resp.StatusMsg = consts.StatusMessage(consts.StatusOK)
	resp.Message = "Hello, " + u.Username
	c.JSON(consts.StatusOK, resp)
}

// UserRegister handles user registration
// @Summary Register a new user
// @Description Register a new user with username and password
// @Tags User
// @Accept json
// @Produce json
// @Param data body user.UserRegisterReq true "Request body with username and password"
// @Success 200 {object} user.UserRegisterResp
// @Router /user/register [post]
func UserRegister(ctx context.Context, c *app.RequestContext) {
	var req user.UserRegisterReq
	if err := c.BindAndValidate(&req); err != nil {
		handleError(c, consts.StatusBadRequest, err)
		return
	}

	// check if username and password are empty
	if req.Username == "" || req.Password == "" {
		handleError(c, consts.StatusBadRequest, fmt.Errorf("username and password are required"))
		return
	}

	// check if username exists
	userQuery := query.User
	u, err := userQuery.WithContext(ctx).Where(userQuery.Username.Eq(req.Username)).Take()
	if u != nil {
		handleError(c, consts.StatusConflict, fmt.Errorf("username exists"))
		return
	}

	// hash password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		handleError(c, consts.StatusInternalServerError, err)
		return
	}

	// create user
	newUser := model.User{
		Username: req.Username,
		Password: string(hashedPassword),
	}
	if err := query.User.Create(&newUser); err != nil {
		handleError(c, consts.StatusInternalServerError, err)
		return
	}

	// response
	resp := new(user.UserRegisterResp)
	resp.StatusCode = consts.StatusOK
	resp.StatusMsg = consts.StatusMessage(consts.StatusOK)
	resp.UserId = int64(newUser.ID)
	resp.Token = generateToken(int64(newUser.ID))

	c.JSON(consts.StatusOK, resp)
}

// UserLogin handles user login
// @Summary User Login
// @Description User login with username and password
// @Tags User
// @Accept json
// @Produce json
// @Param data body user.UserLoginReq true "Request body with username and password"
// @Success 200 {object} user.UserLoginResp
// @Router /user/login [post]
func UserLogin(ctx context.Context, c *app.RequestContext) {
	var req user.UserLoginReq
	if err := c.BindAndValidate(&req); err != nil {
		handleError(c, consts.StatusBadRequest, err)
		return
	}

	// check if username and password are empty
	userQuery := query.User
	u, err := userQuery.WithContext(ctx).Where(userQuery.Username.Eq(req.Username)).First()
	if err != nil {
		handleError(c, consts.StatusNotFound, fmt.Errorf("user not found"))
		return
	}

	// compare password
	if err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(req.Password)); err != nil {
		handleError(c, consts.StatusUnauthorized, fmt.Errorf("password error"))
		return
	}

	// response
	resp := new(user.UserLoginResp)
	resp.StatusCode = consts.StatusOK
	resp.StatusMsg = consts.StatusMessage(consts.StatusOK)
	resp.UserId = int64(u.ID)
	resp.Token = generateToken(int64(u.ID))

	// update token
	if err := updateToken(ctx, int64(u.ID), resp.Token); err != nil {
		handleError(c, consts.StatusInternalServerError, err)
		return
	}

	c.JSON(consts.StatusOK, resp)
}

// AddEmail TODO handles adding an email to a user account
// @Summary Add email
// @Description Add an email to a user account by user ID and token
// @Tags User
// @Accept json
// @Produce json
// @Param data body user.AddEmailReq true "Request body with user ID, token, and email"
// @Success 200 {object} user.AddEmailResp
// @Router /user/add_email [post]
func AddEmail(ctx context.Context, c *app.RequestContext) {
	var err error
	var req user.AddEmailReq
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	resp := new(user.AddEmailResp)

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
