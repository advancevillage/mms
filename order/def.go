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
	Schema   = "orders"
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
	QueryOrders(user *api.User, where map[string]interface{}, page int, perPage int, sort map[string]interface{}) ([]api.Order, int64, error)
}

type IStock interface {
	//加减库存
}

type Service struct {
	repo IOrder
	logger logs.Logs
	sm   *fsm
}

