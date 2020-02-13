//author: richard
package route

import (
	"github.com/advancevillage/3rd/https"
	"mms/config"
	"mms/language"
	"mms/session"
	"mms/user"
)

const (
	//用户 user
	UserCode = 2000
	UserMsg  = "user"
	//令牌 token
	TokenCode = 2100
	TokenMsg  = "token"
	//Session
	SessionCode = 2200
	SessionMsg  = "session"
	//购物车 cart
	CartCode    = 2300
	CartMsg     = "cart"
	//图片 image
	ImageCode = 1700
	ImageMsg  = "category"

	SnowFlakeIdLength = 18
)

const (
	BodyErrorCode = 11
	BodyErrorMsg  = "request query body error"

	ContextErrorCode = 12
	ContextErrorMsg  = "request body struct format error"

	CreateErrorCode = 13
	CreateErrorMsg  = "create error"
	QueryErrorCode  = 14
	UpdateErrorCode = 15
	DeleteErrorCode = 16
	IDErrorCode     = 20
	IDErrorMsg      = "id error"

	OperateSuccess = "operate success"
)

var router = func (api API) []https.Router {
	return []https.Router {
		//service
		{"GET", "/v1/service/ping", api.ping},
		{"GET", "/v1/service/version", api.version},
		//user
		{"GET", "/v1/users", api.QueryUser},
		{"POST", "/v1/users", api.CreateUser},
		{"OPTIONS", "/v1/users", api.ping},
		//token
		{"POST", "/v1/tokens", api.CreateToken},
		{"OPTIONS", "/v1/tokens", api.ping},
		//carts
		{"POST", "/v1/carts", api.CreateCart},
		{"GET", "/v1/carts", api.QueryCart},
		{"PUT", "/v1/carts/:pathId", api.UpdateCart},
		{"DELETE", "/v1/carts/:pathId", api.DeleteCart},
		{"OPTIONS", "/v1/carts", api.ping},
		{"OPTIONS", "/v1/carts/:pathId", api.ping},
	}
}

type API interface {
	ping(ctx *https.Context)
	version(ctx *https.Context)
	//user
	QueryUser (ctx *https.Context)
	CreateUser(ctx *https.Context)
	//token
	CreateToken(ctx *https.Context)
	//carts
	CreateCart(ctx *https.Context)
	QueryCart(ctx *https.Context)
	DeleteCart(ctx *https.Context)
	UpdateCart(ctx *https.Context)
}

type httpError struct {
	Code    string `json:"code,omitempty"`
	Message string `json:"message,omitempty"`
}

type httpOk struct {
	Code    int `json:"code,omitempty"`
	Message string `json:"message,omitempty"`
}

type Service struct {
	configService  *config.Service
	langService    *language.Service
	userService    *user.Service
	sessionService *session.Service
}
