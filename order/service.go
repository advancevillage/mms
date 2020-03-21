//author: richard
package order

import (
	"errors"
	"fmt"
	"github.com/advancevillage/3rd/logs"
	"github.com/advancevillage/3rd/pay"
	"github.com/advancevillage/3rd/storages"
	"github.com/advancevillage/3rd/times"
	"mms/api"
)

func NewService(storage storages.StorageExd, pay pay.IPay, logger logs.Logs) *Service {
	s := &Service{
		repo: NewRepoMongo(storage),
		pay: pay,
		logger: logger,
		sm: &fsm{},
	}
	//初始化状态自动机
	s.sm.CreateState(StateOrdered, StatePendingPay).CreateEvent(EventPay, s.ActionPayOrder)
	s.sm.CreateState(StatePayed, StatePendingShip).CreateEvent(EventShip, s.ActionShipping)
	return s
}

func (s *Service) ActionPayOrder (o *api.Order) error {
	//移除购物车
	if o == nil {
		return errors.New("order is nil")
	}
	cartId := make([]string, len(o.Stocks))
	for i := range o.Stocks {
		cartId = append(cartId, o.Stocks[i].CartId)
	}
	err := s.repo.ClearCart(o.User, cartId[:]...)
	if err != nil {
		s.logger.Error(err.Error())
	}
	//锁定库存
	for i := range o.Stocks {
		err = s.repo.UpdateStock(&o.Stocks[i])
		if err != nil {
			s.logger.Error(err.Error())
		}
	}
	//支付订单
	callback := make(map[string]string)
	err = s.pay.Transaction(o.Pay.Nonce, o.Total, &callback)
	if err != nil {
		s.logger.Alert(err.Error())
		return err
	}
	o.T = &api.Transaction{}
	o.T.Id = callback["id"]
	o.T.Status = callback["status"]
	o.State = StatePayed
	o.NextState = StatePendingShip
	o.UpdateTime = times.Timestamp()
	o.PayTime    = times.Timestamp()
	o.SnapshotPayed = true
	err = s.UpdateOder(o.User, o)
	if err != nil {
		return err
	}
	return nil
}

func (s *Service) ActionShipping(o *api.Order) error {
	return nil
}

func (s *Service) CreateOrder(user *api.User, o *api.Order) error {
	if user == nil || o == nil {
		return errors.New("user or order is nil")
	}

	//生成订单
	orderId, err := s.CreateOrderId(user)
	if err != nil {
		s.logger.Error(err.Error())
		return err
	}
	o.Id      = orderId
	o.User    = user
	o.State   = StateOrdered
	o.NextState = StatePendingPay
	o.SnapshotOrdered = true
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
	return nil
}

func (s *Service) UpdateOder(user *api.User, o *api.Order) error {
	if user == nil || o == nil {
		return errors.New("user or order is nil")
	}
	o.User = user
	err := s.repo.UpdateOrder(o)
	if err != nil {
		s.logger.Critical("pay order fail %s", err.Error())
		return err
	}
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

func (s *Service) QueryStock(stock *api.Stock) (*api.Stock, error) {
	if stock == nil {
		return nil, errors.New("stock is nil")
	}
	org, err := s.repo.QueryStock(stock)
	if err != nil {
		s.logger.Error(err.Error())
		return nil, err
	}
	return org, nil
}

func (s *Service) CreatePayToken(user *api.User) (string, error) {
	if user == nil {
		return "", errors.New("user is nil")
	}
	var token string
	err := s.pay.ClientToken(&token)
	if err != nil {
		s.logger.Error(err.Error())
		return "", err
	}
	return token, nil
}
