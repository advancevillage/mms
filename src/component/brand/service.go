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

func (s *Service) QueryBrands(status int, page int, perPage int) ([]Brand, error) {
	where := make(map[string]interface{})
	where["brandStatus"] = s.brandStatus(status)
	brands, err := s.repo.QueryBrands(where, page, perPage)
	if err != nil {
		s.logger.Error(err.Error())
		return nil, err
	}
	return brands, nil
}

func (s *Service) CreateBrand(brandNameEn string) error {
	brd := &Brand{}
	brd.BrandId = utils.SnowFlakeIdString()
	brd.BrandName.English = brandNameEn
	brd.BrandStatus = StatusActive
	brd.CreateTime = times.Timestamp()
	brd.UpdateTime = times.Timestamp()
	brd.DeleteTime = 0
	err := s.repo.CreateBrand(brd)
	if err != nil {
		s.logger.Error(err.Error())
		return err
	}
	return nil
}

func (s *Service) UpdateBrand(brandId string, brandNameEn, brandNameCn string, status int) error {
	brand, err := s.QueryBrandById(brandId)
	if err != nil {
		s.logger.Error(err.Error())
		return err
	}
	status = s.brandStatus(status)
	brand.BrandName.English = brandNameEn
	brand.BrandName.Chinese = brandNameCn
	if status != StatusInvalid {
		brand.BrandStatus = status
	}
	err = s.repo.UpdateBrand(brand)
	if err != nil {
		s.logger.Error(err.Error())
		return err
	}
	return nil
}

func (s *Service) DeleteBrand(brandId string) error {
	brand, err := s.QueryBrandById(brandId)
	if err != nil {
		s.logger.Error(err.Error())
		return err
	}
	brand.BrandStatus = StatusDeleted
	brand.UpdateTime  = times.Timestamp()
	brand.DeleteTime  = times.Timestamp()
	err = s.repo.UpdateBrand(brand)
	if err != nil {
		s.logger.Error(err.Error())
		return err
	}
	return nil
}

func (s *Service) brandStatus(status int) int {
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

