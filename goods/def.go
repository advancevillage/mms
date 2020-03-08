//author: richard
package goods

import (
	"github.com/advancevillage/3rd/logs"
	"mms/api"
)

const (
	Schema = "goods"
)

type IGoods interface {
	CreateGoods(g *api.Goods) error
	UpdateGoods(g *api.Goods) error
	QueryOneGoods (id string) (*api.Goods, error)
	QueryGoods(where map[string]interface{}, page int, perPage int, sort map[string]interface{}) ([]api.Goods, int64, error)
	IStock
}

type IStock interface {
	IncreaseStock(stock *api.Stocks) error
	DecreaseStock(stock *api.Stocks) error
}


type Service struct {
	repo   IGoods
	logger logs.Logs
}
