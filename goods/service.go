//author: richard
package goods

import (
	"github.com/advancevillage/3rd/logs"
	"github.com/advancevillage/3rd/storages"
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

	return nil
}