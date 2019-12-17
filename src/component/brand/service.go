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
	return &Service{repo:NewBrandRepoEs7(storage), logger:logger}
}

func (s *Service) QueryBrand(BrandId int64) (*Brand, error) {
	brd, err := s.repo.QueryBrand(BrandId)
	if err != nil {
		s.logger.Error(err.Error())
		return nil, err
	}
	return brd, nil
}

func (s *Service) QueryBrands(BrandId ...int64) ([]*Brand, error) {
	var length = len(BrandId)
	var brands = make([]*Brand, 0, length)
	for i := 0; i < length; i++ {
		brd , err := s.repo.QueryBrand(BrandId[i])
		if err != nil {
			s.logger.Info(err.Error())
		} else {
			brands = append(brands, brd)
		}
	}
	return brands, nil
}

func (s *Service) CreateBrand(brandName string, BrandStatus int) error {
	brd := &Brand{}
	brd.BrandId = utils.SnowFlakeId()
	brd.BrandName = brandName
	brd.BrandStatus = BrandStatus
	brd.BrandCreateTime = times.Timestamp()
	brd.BrandUpdateTime = times.Timestamp()
	brd.BrandDeleteTime = 0
	err := s.repo.CreateBrand(brd)
	if err != nil {
		s.logger.Error(err.Error())
		return err
	}
	return nil
}

