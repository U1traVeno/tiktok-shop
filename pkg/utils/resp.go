package utils

import (
	"github.com/henrylee2cn/goutil/status"
)

type BaseResp struct {
	StatusCode int32
	StatusMsg  string
}

// BuildBaseResp convert error and build BaseResp
func BuildBaseResp(statu int32, err error, meg string) *BaseResp {
	if err == nil {
		return baseResp(status.OK, "success")
	}

	return baseResp(statu, err.Error()+meg)
}

// baseResp build BaseResp from error
func baseResp(status int32, meg string) *BaseResp {
	return &BaseResp{
		StatusCode: status,
		StatusMsg:  meg,
	}
}
