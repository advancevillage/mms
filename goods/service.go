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

func NewService(storage storages.Storage, logger logs.Logs) *Service {
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

	err := s.repo.CreateGoods(g)
	if err != nil {
		s.logger.Error(err.Error())
		return err
	}

	return nil
}