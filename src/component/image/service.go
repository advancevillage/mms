//author: richard
package image

import (
	"github.com/advancevillage/3rd/logs"
	"github.com/advancevillage/3rd/storages"
	"github.com/advancevillage/3rd/times"
	"github.com/advancevillage/3rd/utils"
	"mms/src/language"
)

type Service struct {
	repo   IImage
	logger logs.Logs
}

func NewImageService(storage storages.Storage, logger logs.Logs) *Service {
	return &Service{repo:NewImageRepoMgo(storage), logger:logger}
}

func (s *Service) QueryImageById(id string) (*Image, error) {
	color, err := s.repo.QueryImage(id)
	if err != nil {
		s.logger.Error(err.Error())
		return nil, err
	}
	return color, nil
}

func (s *Service) QueryImages(status int, page int, perPage int) ([]Image, int64, error) {
	where := make(map[string]interface{})
	sort := make(map[string]interface{})
	where["status"] = s.Status(status)
	sort["createTime"] = s.desc()
	colors, total, err := s.repo.QueryImages(where, page, perPage, sort)
	if err != nil {
		s.logger.Error(err.Error())
		return nil, 0, err
	}
	return colors, total, nil
}

func (s *Service) CreateImage(desc *language.Languages, isDefault bool, url string, customType string, customDirection int) error {
	value := &Image{}
	value.Id = utils.SnowFlakeIdString()
	value.Description = desc
	value.Status = StatusActive
	value.Url = url
	value.IsDefault = isDefault
	value.CustomType = customType
	value.CustomDirection = customDirection
	value.CreateTime = times.Timestamp()
	value.UpdateTime = times.Timestamp()
	value.DeleteTime = 0
	err := s.repo.CreateImage(value)
	if err != nil {
		s.logger.Error(err.Error())
		return err
	}
	return nil
}

func (s *Service) UpdateImage(id string, desc *language.Languages, status int) error {
	value, err := s.QueryImageById(id)
	if err != nil {
		s.logger.Error(err.Error())
		return err
	}
	value.Description = desc
	value.UpdateTime   = times.Timestamp()
	value.Status = s.Status(status)
	err = s.repo.UpdateImage(value)
	if err != nil {
		s.logger.Error(err.Error())
		return err
	}
	return nil
}

func (s *Service) DeleteImage(id string) error {
	value, err := s.QueryImageById(id)
	if err != nil {
		s.logger.Error(err.Error())
		return err
	}
	value.Status = StatusDeleted
	value.UpdateTime  = times.Timestamp()
	value.DeleteTime  = times.Timestamp()
	err = s.repo.UpdateImage(value)
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


