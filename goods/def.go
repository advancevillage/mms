//author: richard
package goods

import "mms/api"

const (
	Schema = "goods"
	StockSchema = "stocks"
)

type IGoods interface {
	CreateGoods(g *api.Goods) error
	UpdateGoods(g *api.Goods) error
	QueryOneGoods (id string) (*api.Goods, error)
	QueryGoods(where map[string]interface{}, page int, perPage int, sort map[string]interface{}) ([]api.Goods, int64, error)

	IStock
}

type IStock interface {
	CreateStock(stock *api.Stock) error
	QueryStocks(goodsId string) ([]api.Stock, int64, error)
}