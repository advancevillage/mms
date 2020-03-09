//author: richard
package goods

import (
	"errors"
	"github.com/advancevillage/3rd/logs"
	"github.com/advancevillage/3rd/storages"
	"github.com/advancevillage/3rd/times"
	"github.com/advancevillage/3rd/utils"
	"mms/api"
)

type Service struct {
	repo   IGoods
	logger logs.Logs
}

func NewService(storage storages.StorageExd, logger logs.Logs) *Service {
	return &Service{repo:NewRepoMongo(storage), logger:logger}
}

func (s *Service) CreateGoods(g *api.Goods) error {
	//校验逻辑
	if g == nil {
		return errors.New("goods is nil")
	}

	g.Id = utils.SnowFlakeIdString()
	g.CreateTime = times.Timestamp()
	g.UpdateTime = times.Timestamp()
	g.DeleteTime = 0

	stocks := g.Stocks  //库存信息和商品信息分开存储
	g.Stocks = nil

	for i := range stocks {
		stocks[i].GoodsId = g.Id
		stocks[i].CreateTime = g.CreateTime
		stocks[i].UpdateTime = g.UpdateTime
		stocks[i].DeleteTime = g.DeleteTime
		stocks[i].Version    = 0
		stocks[i].Id         = utils.RandsNumberString(6)
	}
	//上架商品信息
	err := s.repo.CreateGoods(g)
	if err != nil {
		s.logger.Error(err.Error())
		return err
	}
	//设置库存信息
	for i := range stocks {
		err = s.repo.CreateStock(&stocks[i])
		if err != nil {
			s.logger.Error(err.Error())
		} else {
			continue
		}
	}

	return nil
}

func (s *Service) QueryOneGoods(id string) (*api.Goods,error) {
	//校验逻辑
	return s.repo.QueryOneGoods(id)
}

func (s *Service) QueryGoods(page int, perPage int) ([]api.Goods, int64, error) {
	where := make(map[string]interface{})
	sort := make(map[string]interface{})
	sort["createTime"] = s.desc()
	goods, total, err := s.repo.QueryGoods(where, page, perPage, sort)
	if err != nil {
		s.logger.Error(err.Error())
		return nil, 0, err
	}

	for i := range goods {
		goods[i].Stocks, _, _ = s.repo.QueryStocks(goods[i].Id)
	}

	return goods, total, nil
}

func (s *Service) asc() int {
	return 1
}

func (s *Service) desc() int {
	return -1
}