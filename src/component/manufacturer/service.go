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

func (s *Service) QueryManufacturerById(id string) (*Manufacturer, error) {
	color, err := s.repo.QueryManufacturer(id)
	if err != nil {
		s.logger.Error(err.Error())
		return nil, err
	}
	return color, nil
}

func (s *Service) QueryManufacturers(status int, page int, perPage int) ([]Manufacturer, int64, error) {
	where := make(map[string]interface{})
	where["manufacturerStatus"] = s.Status(status)
	manufacturers, total, err := s.repo.QueryManufacturers(where, page, perPage)
	if err != nil {
		s.logger.Error(err.Error())
		return nil, 0, err
	}
	return manufacturers, total, nil
}

func (s *Service) CreateManufacturer(contact string, phone, email string, nameEn, addressEn string) error {
	value := &Manufacturer{}
	value.Id = utils.SnowFlakeIdString()
	value.Name.English    = nameEn
	value.Address.English = addressEn
	value.Status = StatusActive
	value.Contact = contact
	value.ContactEmail = email
	value.ContactPhone = phone
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

func (s *Service) UpdateManufacturer(id string, phone, email, contact string, status int) error {
	value, err := s.QueryManufacturerById(id)
	if err != nil {
		s.logger.Error(err.Error())
		return err
	}
	value.Contact = contact
	value.ContactPhone = phone
	value.ContactEmail = email
	value.UpdateTime   = times.Timestamp()
	value.Status = s.Status(status)
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
	value.Status = StatusDeleted
	value.UpdateTime  = times.Timestamp()
	value.DeleteTime  = times.Timestamp()
	err = s.repo.UpdateManufacturer(value)
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
