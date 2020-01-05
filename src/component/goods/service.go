//author: richard
package goods

import (
	"github.com/advancevillage/3rd/logs"
	"github.com/advancevillage/3rd/storages"
	"github.com/advancevillage/3rd/times"
	"github.com/advancevillage/3rd/utils"
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

func (s *Service) QueryManufacturers(status int, page int, perPage int) ([]Goods, error) {
	where := make(map[string]interface{})
	where["goodsStatus"] = s.Status(status)
	goods, err := s.repo.QueryMerchandises(where, page, perPage)
	if err != nil {
		s.logger.Error(err.Error())
		return nil, err
	}
	return goods, nil
}

func (s *Service) CreateManufacturer(titleEn string, descEn string, costPrice float64) error {
	value := &Goods{}
	value.Id = utils.SnowFlakeIdString()
	value.Title.English = titleEn
	value.DetailedDescription.English = descEn
	value.Status = StatusActive
	value.CostPrice = costPrice
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


