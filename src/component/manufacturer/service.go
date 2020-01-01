//author: richard
package manufacturer

import (
	"github.com/advancevillage/3rd/logs"
	"github.com/advancevillage/3rd/storages"
	"github.com/advancevillage/3rd/times"
	"github.com/advancevillage/3rd/utils"
)

type Service struct {
	repo   IManufacturer
	logger logs.Logs
}

func NewManufacturerService(storage storages.Storage, logger logs.Logs) *Service {
	return &Service{repo:NewManufacturerRepoMgo(storage), logger:logger}
}

func (s *Service) QueryManufacturer(mfId string) (*Manufacturer, error) {
	mf, err := s.repo.QueryManufacturer(mfId)
	if err != nil {
		s.logger.Error(err.Error())
		return nil, err
	}
	return mf, nil
}

func (s *Service) QueryManufacturers(mfId ...string) ([]*Manufacturer, error) {
	var length = len(mfId)
	var mfs = make([]*Manufacturer, 0, length)
	for i := 0; i < length; i++ {
		mf , err := s.repo.QueryManufacturer(mfId[i])
		if err != nil {
			s.logger.Info(err.Error())
		} else {
			mfs = append(mfs, mf)
		}
	}
	return mfs, nil
}

func (s *Service) CreatManufacturer(mf *Manufacturer) error {
	mf.ManufacturerId = utils.SnowFlakeIdString()
	mf.ManufacturerCreateTime = times.Timestamp()
	mf.ManufacturerUpdateTime = times.Timestamp()
	mf.ManufacturerDeleteTime = 0
	err := s.repo.CreateManufacturer(mf)
	if err != nil {
		s.logger.Error(err.Error())
		return err
	}
	return nil
}
