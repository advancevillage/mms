//author: richard
package goods

import (
	"github.com/advancevillage/3rd/logs"
	"github.com/advancevillage/3rd/storages"
	"github.com/advancevillage/3rd/times"
	"github.com/advancevillage/3rd/utils"
)

type Service struct {
	repo   IGoods
	logger logs.Logs
}

func NewGoodsService(storage storages.Storage, logger logs.Logs) *Service {
	return &Service{repo:NewGoodsRepoEs7(storage), logger:logger}
}

func (s *Service) QueryGood(gId int64) (*Goods, error) {
	g, err := s.repo.QueryGoods(gId)
	if err != nil {
		s.logger.Error(err.Error())
		return nil, err
	}
	return g, nil
}

func (s *Service) QueryGoods(gId ...int64) ([]*Goods, error) {
	var length = len(gId)
	var g = make([]*Goods, 0, length)
	for i := 0; i < length; i++ {
		brd , err := s.repo.QueryGoods(gId[i])
		if err != nil {
			s.logger.Info(err.Error())
		} else {
			g = append(g, brd)
		}
	}
	return g, nil
}

func (s *Service) CreateGoods(g *Goods) error {
	g.GoodsId = utils.SnowFlakeId()
	g.GoodsCreateTime = times.Timestamp()
	g.GoodsUpdateTime = times.Timestamp()
	g.GoodsDeleteTime = 0
	err := s.repo.CreateGoods(g)
	if err != nil {
		s.logger.Error(err.Error())
		return err
	}
	return nil
}


