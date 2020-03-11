//author: richard
package order

import (
	"github.com/advancevillage/3rd/logs"
	"mms/api"
)

const (
	EventPay   = 0x01  // 支付事件
	EventShip  = 0x02  // 发货事件
)

const (
	Schema      = "orders"
	StockSchema = "stocks"
	GoodsSchema = "goods"
	CreditSchema = "credit"
	CartSchema   = "carts"
)
const (
	Project   = 0  //show'u
	TypeOrder  = 0 //订单
	TypeReturn = 1 //退货单
	TypeRefund = 2 //退款单

)

const (
	StateOrdered     = "ordered"
	StatePendingPay  = "pending_pay"
	StatePaying      = "paying"
	StatePayed       = "payed"
	StatePendingShip = "pending_ship"
	StateShipping    = "shipping"
	StateShipped     = "shipped"
)

type IOrder interface {
	CreateOrder(o *api.Order) error
	QueryOrders(user *api.User, where map[string]interface{}, page int, perPage int, sort map[string]interface{}) ([]api.Order, int64, error)
	IStock
	IPay
	ICart
}

type IStock interface {
	//乐观锁 多版本实现 库存添加version字段属性 自增 核心CAS(compare and set) CAS 原子性由mongodb保证
	//https://docs.mongodb.com/manual/core/write-operations-atomicity/#update-if-current
	//https://stackoverflow.com/questions/16523621/atomicity-and-cas-operations-in-mongodb
	UpdateStock(stock *api.Stock) error
	QueryStock(stock *api.Stock) (*api.Stock, error)
}

type IPay interface {
	QueryPay(user *api.User, card *api.CreditCard) (*api.CreditCard, error)
}

type ICart interface {
	ClearCart(user *api.User, cartId ...string) error
}

type Service struct {
	repo IOrder
	logger logs.Logs
	sm   *fsm
}

