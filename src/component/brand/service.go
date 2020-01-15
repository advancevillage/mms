//author: richard
package brand

import (
	"github.com/advancevillage/3rd/logs"
	"github.com/advancevillage/3rd/storages"
	"github.com/advancevillage/3rd/times"
	"github.com/advancevillage/3rd/utils"
)

type Service struct {
	repo   IBrand
	logger logs.Logs
}

func NewBrandService(storage storages.Storage, logger logs.Logs) *Service {
	return &Service{repo:NewBrandRepoMgo(storage), logger:logger}
}

func (s *Service) QueryBrandById(brandId string) (*Brand, error) {
	brd, err := s.repo.QueryBrand(brandId)
	if err != nil {
		s.logger.Error(err.Error())
		return nil, err
	}
	return brd, nil
}

func (s *Service) QueryBrands(status int, page int, perPage int) ([]Brand, int64, error) {
	where := make(map[string]interface{})
	where["brandStatus"] = s.Status(status)
	brands, total, err := s.repo.QueryBrands(where, page, perPage)
	if err != nil {
		s.logger.Error(err.Error())
		return nil, 0, err
	}
	return brands, total, nil
}

func (s *Service) CreateBrand(english string, chinese string) error {
	value := &Brand{}
	value.Id = utils.SnowFlakeIdString()
	value.Name.English = english
	value.Name.Chinese = chinese
	value.Status = StatusActive
	value.CreateTime = times.Timestamp()
	value.UpdateTime = times.Timestamp()
	value.DeleteTime = 0
	err := s.repo.CreateBrand(value)
	if err != nil {
		s.logger.Error(err.Error())
		return err
	}
	return nil
}

func (s *Service) UpdateBrand(id string, english, chinese string, status int) error {
	value, err := s.QueryBrandById(id)
	if err != nil {
		s.logger.Error(err.Error())
		return err
	}
	status = s.Status(status)
	value.Name.English = english
	value.Name.Chinese = chinese
	value.Status = s.Status(status)
	err = s.repo.UpdateBrand(value)
	if err != nil {
		s.logger.Error(err.Error())
		return err
	}
	return nil
}

func (s *Service) DeleteBrand(brandId string) error {
	value, err := s.QueryBrandById(brandId)
	if err != nil {
		s.logger.Error(err.Error())
		return err
	}
	value.Status = StatusDeleted
	value.UpdateTime  = times.Timestamp()
	value.DeleteTime  = times.Timestamp()
	err = s.repo.UpdateBrand(value)
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

