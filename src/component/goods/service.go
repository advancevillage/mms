//author: richard
package goods

import (
	"github.com/advancevillage/3rd/logs"
	"github.com/advancevillage/3rd/storages"
	"github.com/advancevillage/3rd/times"
	"github.com/advancevillage/3rd/utils"
	"mms/src/language"
)

type Service struct {
	repo   IMerchandise
	logger logs.Logs
}

func NewGoodsService(storage storages.Storage, logger logs.Logs) *Service {
	return &Service{repo:NewGoodsRepoMgo(storage), logger:logger}
}

func (s *Service) QueryManufacturerById(goodsId string) (*Goods, error) {
	goods, err := s.repo.QueryMerchandise(goodsId)
	if err != nil {
		s.logger.Error(err.Error())
		return nil, err
	}
	return goods, nil
}

func (s *Service) QueryManufacturers(status int, page int, perPage int) ([]Goods, int64, error) {
	where := make(map[string]interface{})
	sort := make(map[string]interface{})
	where["status"] = s.Status(status)
	sort["createTime"] = s.desc()
	goods, total, err := s.repo.QueryMerchandises(where, page, perPage, sort)
	if err != nil {
		s.logger.Error(err.Error())
		return nil, 0, err
	}
	return goods, total, nil
}

func (s *Service) CreateManufacturer(name *language.Languages, title *language.Languages, description *language.Languages, rank int, status int) error {
	value := &Goods{}
	value.Id = utils.SnowFlakeIdString()
	value.Title   = title
	value.Summary = name
	value.Description = description
	value.Status = status
	value.Rank   = rank

	value.CreateTime = times.Timestamp()
	value.UpdateTime = times.Timestamp()
	value.DeleteTime = 0
	err := s.repo.CreateMerchandise(value)
	if err != nil {
		s.logger.Error(err.Error())
		return err
	}
	return nil
}

func (s *Service) DeleteManufacturer(id string) error {
	value, err := s.QueryManufacturerById(id)
	if err != nil {
		s.logger.Error(err.Error())
		return err
	}
	value.Status = StatusDeleted
	value.UpdateTime  = times.Timestamp()
	value.DeleteTime  = times.Timestamp()
	err = s.repo.UpdateMerchandise(value)
	if err != nil {
		s.logger.Error(err.Error())
		return err
	}
	return nil
}

func (s *Service) Status(status int) int {
	switch status {
	case StatusActive:
		status = StatusActive
	case StatusDeleted:
		status = StatusDeleted
	default:
		status = StatusInvalid
	}
	return status
}

func (s *Service) asc() int {
	return 1
}

func (s *Service) desc() int {
	return -1
}


