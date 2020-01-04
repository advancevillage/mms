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
	return &Service{repo:NewGoodsRepoMgo(storage), logger:logger}
}

func (s *Service) QueryGood(goodsId string) (*Goods, error) {
	g, err := s.repo.QueryGoods(goodsId)
	if err != nil {
		s.logger.Error(err.Error())
		return nil, err
	}
	return g, nil
}

func (s *Service) QueryGoods(goodsIds ...string) ([]*Goods, error) {
	var length = len(goodsIds)
	var goods = make([]*Goods, 0, length)
	for i := 0; i < length; i++ {
		good , err := s.repo.QueryGoods(goodsIds[i])
		if err != nil {
			s.logger.Info(err.Error())
		} else {
			goods = append(goods, good)
		}
	}
	return goods, nil
}

func (s *Service) CreateGoods(g *Goods) error {
	g.GoodsId = utils.SnowFlakeIdString()
	g.CreateTime = times.Timestamp()
	g.UpdateTime = times.Timestamp()
	g.DeleteTime = 0
	err := s.repo.CreateGoods(g)
	if err != nil {
		s.logger.Error(err.Error())
		return err
	}
	return nil
}


