package service

import (
	"context"
	"fmt"
	"github.com/U1traVeno/tiktok-shop/biz/dal/model"
	query "github.com/U1traVeno/tiktok-shop/biz/dal/query/user"
	"github.com/U1traVeno/tiktok-shop/biz/model/user"
	"github.com/cloudwego/hertz/pkg/app"
	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	ctx context.Context
	c   *app.RequestContext
}

func NewUserService(ctx context.Context, c *app.RequestContext) *UserService {
	return &UserService{
		ctx: ctx,
		c:   c,
	}
}

func (s *UserService) GetUserInfo() {

}

func (s *UserService) UserRegister(req *user.UserRegisterReq) (userId int64, err error) {
	// check if username and password are empty
	if req.Username == "" || req.Password == "" {
		return 0, fmt.Errorf("username or password is empty")
	}

	// check if username exists
	userQuery := query.User
	// TODO 这一行可以用上WithContext, 需要优化
	u, err := userQuery.Where(userQuery.Username.Eq(req.Username)).Take()
	if u != nil {
		return 0, fmt.Errorf("username already exists")
	}

	// hash password
	// TODO 将这里的加密逻辑封装到utils里
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return 0, err
	}

	// create user
	newUser := model.User{
		Username: req.Username,
		Password: string(hashedPassword),
	}
	if err := query.User.Create(&newUser); err != nil {
		return 0, err
	}
	return userId, nil
}
