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
	OrderCode = 3000
	OrderMsg  = "order"
	//结算页令牌 token
	PayTokenCode = 3100
	PayTokenMsg  = "pay page token"
	//Session
	SessionCode = 2200
	SessionMsg  = "session"
	//购物车 cart
	StockCode   = 3300
	StockMsg    = "stock"
	//收货地址 address
	AddressCode = 2400
	AddressMsg  = "address"
	//图片 image
	CreditCode  = 2500
	CreditMsg   = "credit"

	SnowFlakeIdLength = 18


	PendingHandle = "pending_handle"
	Handling      = "handling"
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
	StockNotEnough = "stock not enough"
	InvalidCreditCard = "invalid credit card"
	InvalidAddress    = "invalid address"
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
