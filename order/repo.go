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

func (s *Mongo) UpdateStock(goods *api.Goods) error {
	if goods == nil {
		return errors.New("goods is nil")
	}
	where := make(map[string]interface{})
	where["goodsId"] = goods.GoodsId
	where["colorId"] = goods.ColorId
	where["sizeId"]  = goods.SizeId
	where["version"] = goods.Version
	return nil
}

func (s *Mongo) QueryStock(goods *api.Goods) (*api.Goods, error) {
	if goods == nil {
		return nil, errors.New("goods is nil")
	}
	where := make(map[string]interface{})
	sort  := make(map[string]interface{})
	where["goodsId"] = goods.GoodsId
	where["colorId"] = goods.ColorId
	where["sizeId"]  = goods.SizeId
	sort["createTime"] = -1
	items, total, err := s.storage.SearchStorageV2Exd(StockSchema, goods.GoodsId, where, 1, 0, sort)
	if err != nil {
		return nil, err
	}
	if total <= 0 {
		return nil, errors.New("total <= 0")
	}
	value := api.Goods{}
	err = json.Unmarshal(items[0], &value)
	if err != nil {
		return nil, err
	}
	return &value, nil
}
