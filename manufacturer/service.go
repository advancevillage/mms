//author: richard
package manufacturer

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

func (s *Service) QueryManufacturerById(id string) (*api.Manufacturer, error) {
	color, err := s.repo.QueryManufacturer(id)
	if err != nil {
		s.logger.Error(err.Error())
		return nil, err
	}
	return color, nil
}

func (s *Service) QueryManufacturers(page int, perPage int) ([]api.Manufacturer, int64, error) {
	where := make(map[string]interface{})
	sort := make(map[string]interface{})
	sort["createTime"] = s.desc()
	manufacturers, total, err := s.repo.QueryManufacturers(where, page, perPage, sort)
	if err != nil {
		s.logger.Error(err.Error())
		return nil, 0, err
	}
	return manufacturers, total, nil
}

func (s *Service) CreateManufacturer(value *api.Manufacturer) error {
	if value == nil {
		return errors.New("manufacturer is nil")
	}
	value.Id = utils.SnowFlakeIdString()
	value.CreateTime = times.Timestamp()
	value.UpdateTime = times.Timestamp()
	value.DeleteTime = 0
	err := s.repo.CreateManufacturer(value)
	if err != nil {
		s.logger.Error(err.Error())
		return err
	}
	return nil
}

func (s *Service) UpdateManufacturer(m *api.Manufacturer) error {
	value, err := s.QueryManufacturerById(m.Id)
	if err != nil {
		s.logger.Error(err.Error())
		return err
	}
	value.Contact = m.Contact
	value.Phone   = m.Phone
	value.Email   = m.Email
	value.UpdateTime   = times.Timestamp()
	err = s.repo.UpdateManufacturer(value)
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
	value.UpdateTime  = times.Timestamp()
	value.DeleteTime  = times.Timestamp()
	err = s.repo.UpdateManufacturer(value)
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
