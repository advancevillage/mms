//author: richard
package brand

import (
	"errors"
	"github.com/advancevillage/3rd/logs"
	"github.com/advancevillage/3rd/storages"
	"github.com/advancevillage/3rd/times"
	"github.com/advancevillage/3rd/utils"
	"mms/api"
)

func NewService(storage storages.Storage, logger logs.Logs) *Service {
	return &Service{repo:NewRepoMongo(storage), logger:logger}
}

func (s *Service) QueryBrandById(brandId string) (*api.Brand, error) {
	brd, err := s.repo.QueryBrand(brandId)
	if err != nil {
		s.logger.Error(err.Error())
		return nil, err
	}
	return brd, nil
}

func (s *Service) QueryBrands(page int, perPage int) ([]api.Brand, int64, error) {
	where := make(map[string]interface{})
	sort := make(map[string]interface{})
	sort["createTime"] = s.desc()
	brands, total, err := s.repo.QueryBrands(where, page, perPage, sort)
	if err != nil {
		s.logger.Error(err.Error())
		return nil, 0, err
	}
	return brands, total, nil
}

func (s *Service) CreateBrand(value *api.Brand) error {
	if value == nil {
		return  errors.New("brand is nil")
	}

	value.Id = utils.SnowFlakeIdString()
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

func (s *Service) UpdateBrand(brand *api.Brand ) error {
	value, err := s.QueryBrandById(brand.Id)
	if err != nil {
		s.logger.Error(err.Error())
		return err
	}
	value.Name = brand.Name
	value.UpdateTime = times.Timestamp()
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
	value.UpdateTime  = times.Timestamp()
	value.DeleteTime  = times.Timestamp()
	err = s.repo.UpdateBrand(value)
	if err != nil {
		s.logger.Error(err.Error())
		return err
	}
	return nil
}

func (s *Service) asc() int {
	return 1
}

func (s *Service) desc() int {
	return -1
}
