//author: richard
package order

import (
	"encoding/json"
	"errors"
	"github.com/advancevillage/3rd/storages"
	"mms/api"
)

type Mongo struct {
	storage storages.StorageExd
}

func NewRepoMongo(storage storages.StorageExd) *Mongo {
	return &Mongo{storage:storage}
}

func (s *Mongo) CreateOrder(o *api.Order) error {
	if o == nil {
		return errors.New("order is nil")
	}
	buf, err := json.Marshal(o)
	if err != nil {
		return err
	}
	return s.storage.CreateStorageV2Exd(Schema, o.User.Id, o.Id, buf)
}

func (s *Mongo) QueryOrders(user *api.User, where map[string]interface{}, page int, perPage int, sort map[string]interface{}) ([]api.Order, int64, error) {
	items, total, err := s.storage.SearchStorageV2Exd(Schema, user.Id, where, perPage, page * perPage, sort)
	if err != nil {
		return nil, 0, err
	}
	values := make([]api.Order, 0, len(items))
	for i := range items {
		buf := items[i]
		value := api.Order{}
		err = json.Unmarshal(buf, &value)
		if err != nil {
			return nil, 0, err
		} else {
			values = append(values, value)
		}
	}
	return values, total, nil
}

func (s *Mongo) UpdateStock(goods *api.Stock) error {
	if goods == nil {
		return errors.New("goods is nil")
	}
	return nil
}

func (s *Mongo) QueryStock(stock *api.Stock) (*api.Stock, error) {
	if stock == nil {
		return nil, errors.New("stock is nil")
	}
	buf, err := s.storage.QueryStorageV2(GoodsSchema, stock.GoodsId)
	if err != nil {
		return nil, err
	}
	value := &api.Goods{}
	err = json.Unmarshal(buf, value)
	if err != nil {
		return nil, err
	}
	i := 0
	for i = range value.Stocks {
		if value.Stocks[i].ColorId == stock.ColorId && value.Stocks[i].SizeId == stock.SizeId {
			break
		}
	}
	if i >= len(value.Stocks) {
		return nil, errors.New("stock query error")
	}
	return &value.Stocks[i], nil
}
