//author: richard
package order

import (
	"errors"
	"fmt"
	"github.com/advancevillage/3rd/logs"
	"github.com/advancevillage/3rd/storages"
	"github.com/advancevillage/3rd/times"
	"mms/api"
)

func NewService(storage storages.StorageExd, logger logs.Logs) *Service {
	s := &Service{
		repo: NewRepoMongo(storage),
		logger: logger,
		sm: &fsm{},
	}
	//初始化状态自动机
	s.sm.CreateState(StateOrdered, StatePendingPay).CreateEvent(EventPay, s.ActionPayOrder)
	s.sm.CreateState(StatePayed, StatePendingShip).CreateEvent(EventShip, s.ActionShipping)
	return s
}

func (s *Service) ActionPayOrder (o *api.Order) error {

	return nil
}

func (s *Service) ActionShipping(o *api.Order) error {
	return nil
}

func (s *Service) CreateOrder(user *api.User, addr *api.Address, goods []api.Goods, credit *api.Transaction) error {
	//锁定库存
	//生成订单
	//支付订单
	return nil
}

func (s *Service) QueryOrders(user *api.User, page, perPage int) ([]api.Order, int64, error) {
	if user == nil {
		return nil, 0, errors.New("user is nil")
	}
	where := make(map[string]interface{})
	sort  := make(map[string]interface{})
	sort["orderTime"] = -1
	orders, total, err := s.repo.QueryOrders(user, where, page, perPage, sort)
	if err != nil {
		s.logger.Error(err.Error())
		return nil, 0, err
	}
	return orders, total, nil
}

func (s *Service) CreateOrderId(user *api.User) (string, error) {
	if user == nil {
		return "", errors.New("user is nil")
	}
	project  := Project
	types    := TypeOrder
	datetime := times.TimeFormatString("060102150405")
	where := make(map[string]interface{})
	sort  := make(map[string]interface{})
	where["id"] = map[string]interface{} {
		"$regex": fmt.Sprintf("%d%d%s", project, types, datetime),
	}
	_, total, err := s.repo.QueryOrders(user, where, 0, 0, sort)
	if err != nil {
		s.logger.Error(err.Error())
		return "", err
	}
	id := fmt.Sprintf("%d%d%s%04d", project, types, datetime, total)
	return id, nil
}


