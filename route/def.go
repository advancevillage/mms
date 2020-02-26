//author: richard
package route

import (
	"github.com/advancevillage/3rd/https"
	"mms/config"
	"mms/language"
	"mms/order"
	"mms/session"
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
	//收货地址 address
	AddressCode = 2400
	AddressMsg  = "address"
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
		//order
		{"POST", "/v1/order", api.CreateOrder},
		{"OPTIONS", "/v1/order", api.ping},
		{"OPTIONS", "/v1/order/:pathId", api.ping},
	}
}

type API interface {
	ping(ctx *https.Context)
	version(ctx *https.Context)
	//order
	CreateOrder(ctx *https.Context)
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
	orderService   *order.Service
	sessionService *session.Service
}
