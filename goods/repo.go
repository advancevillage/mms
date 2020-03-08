//author: richard
package goods

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

func (s *Mongo) CreateGoods(g *api.Goods) error {
	body, err := json.Marshal(g)
	if err != nil {
		return err
	}
	return s.storage.CreateStorageV2(Schema, g.Id, body)
}

func (s *Mongo) UpdateGoods(g *api.Goods) error {
	body, err := json.Marshal(g)
	if err != nil {
		return err
	}
	return s.storage.UpdateStorageV2(Schema, g.Id, body)
}

func (s *Mongo) QueryOneGoods(id string) (*api.Goods, error) {
	buf, err := s.storage.QueryStorageV2(Schema, id)
	if err != nil {
		return nil, err
	}
	value := api.Goods{}
	err = json.Unmarshal(buf, &value)
	if err != nil {
		return nil, err
	}
	return &value, nil
}

func (s *Mongo) QueryGoods(where map[string]interface{}, page int, perPage int, sort map[string]interface{}) ([]api.Goods, int64, error) {
	items, total, err := s.storage.QueryStorageV3(Schema, where, perPage, page * perPage, sort)
	if err != nil {
		return nil, 0, err
	}
	values := make([]api.Goods, 0, len(items))
	for i := range items {
		buf := items[i]
		value := api.Goods{}
		err = json.Unmarshal(buf, &value)
		if err != nil {
			return nil, 0, err
		} else {
			values = append(values, value)
		}
	}
	return values, total, nil
}

func (s *Mongo) IncreaseStock(stock *api.Stocks) error {
	if stock == nil {
		return errors.New("stock is nil")
	}
	buf, err := json.Marshal(stock)
	if err != nil {
		return err
	}
	where := make(map[string]interface{})
	where["stock.goodsId"] = stock.GoodsId
	where["stock.colorId"] = stock.ColorId
	where["stock.sizeId"]  = stock.SizeId

	err = s.storage.UpdateStorageV2Exd(Schema, Schema, where,)
	if err != nil {
		return err
	}
	return nil
}

func (s *Mongo) DecreaseStock(stock *api.Stocks) error {
	return nil
}
