//author: richard
package order

import (
	"encoding/json"
	"errors"
	"fmt"
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

func (s *Mongo) UpdateStock(stock *api.Stock) error {
	if stock == nil {
		return errors.New("stock is nil")
	}
	where := make(map[string]interface{})
	where["goodsId"] = stock.GoodsId
	where["sizeId"]  = stock.SizeId
	where["colorId"] = stock.ColorId
	where["version"] = stock.Version




	return nil
}

func (s *Mongo) QueryStock(stock *api.Stock) (*api.Stock, error) {
	if stock == nil {
		return nil, errors.New("stock is nil")
	}
	where := make(map[string]interface{})
	sort  := make(map[string]interface{})
	where["goodsId"] = stock.GoodsId
	where["colorId"] = stock.ColorId
	where["sizeId"]  = stock.SizeId
	sort["createTime"] = -1
	items, total, err := s.storage.SearchStorageV2Exd(StockSchema, stock.GoodsId, where, 1, 0, sort)
	if err != nil {
		return nil, err
	}
	if total <= 0 {
		return nil, errors.New("total <= 0")
	}
	total = 0
	value := api.Stock{}
	err = json.Unmarshal(items[total], &value)
	if err != nil {
		return nil, err
	}
	return &value, nil
}

func (s *Mongo) QueryPay(user *api.User, card *api.CreditCard) (*api.CreditCard, error) {
	if user == nil || card == nil {
		return nil, errors.New("user or card is nil")
	}
	where := make(map[string]interface{})
	where["id"] = card.Id
	where["number"] = map[string]interface{}{
		"$regex": fmt.Sprintf("^%s[0-9]+%s$", card.Number[:4], card.Number[len(card.Number)-4:]),
	}
	items, total, err := s.storage.SearchStorageV2Exd(CreditSchema, user.Id, where, 1, 0, nil)
	if err != nil {
		return nil, err
	}
	if total > 1 || total < 1 {
		return nil, errors.New("credit card info error")
	}
	total = 0
	credit := api.CreditCard{}
	err = json.Unmarshal(items[total], &credit)
	if err != nil {
		return nil, err
	}
	return &credit, nil
}

func (s *Mongo) ClearCart(user *api.User, cartId ...string) error {
	if user == nil {
		return errors.New("user is nil")
	}
	if len(cartId) <= 0 {
		return nil
	}
	return s.storage.DeleteStorageV2Exd(CartSchema, user.Id, nil, cartId ...)

}
