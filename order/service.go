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

func (s *Service) CreateOrder(user *api.User, o *api.Order) error {
	if user == nil || o == nil {
		return errors.New("user or order is nil")
	}
	//TODO 锁定库存
	//生成订单
	orderId, err := s.CreateOrderId(user)
	if err != nil {
		s.logger.Error(err.Error())
		return err
	}
	o.Id      = orderId
	o.State   = StateOrdered
	o.NextState = StatePendingPay
	o.OrderTime = times.Timestamp()
	err = s.repo.CreateOrder(o)
	if err != nil {
		s.logger.Error(err.Error())
		return err
	}
	//生成订单后交付给自动机处理
	action, err := s.sm.QueryAction(o.State, o.NextState, EventPay)
	if err != nil {
		s.logger.Critical("fsm don't handle %s->%s", o.State, o.NextState)
		return nil
	}
	err = action(o)
	if err != nil {
		s.logger.Critical("pay order fail %s", err.Error())
		return err
	}
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
	id := fmt.Sprintf("%d%d%s%04d", project, types, datetime, total + 1)
	return id, nil
}

func (s *Service) QueryStock(goods *api.Goods) (*api.Goods, error) {
	if goods == nil {
		return nil, errors.New("goods is nil")
	}
	goods, err := s.repo.QueryStock(goods)
	if err != nil {
		s.logger.Error(err.Error())
		return nil, err
	}
	return goods, nil
}


