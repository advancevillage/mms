//author: richard
package color

import (
	"github.com/advancevillage/3rd/logs"
	"github.com/advancevillage/3rd/storages"
	"github.com/advancevillage/3rd/times"
	"github.com/advancevillage/3rd/utils"
	"mms/src/language"
)

type Service struct {
	repo   IColor
	logger logs.Logs
}

func NewColorService(storage storages.Storage, logger logs.Logs) *Service {
	return &Service{repo:NewColorRepoMgo(storage), logger:logger}
}

func (s *Service) QueryColorById(colorId string) (*Color, error) {
	color, err := s.repo.QueryColor(colorId)
	if err != nil {
		s.logger.Error(err.Error())
		return nil, err
	}
	return color, nil
}

func (s *Service) QueryColors(status int, page int, perPage int) ([]Color, int64, error) {
	where := make(map[string]interface{})
	sort := make(map[string]interface{})
	where["status"] = s.Status(status)
	sort["createTime"] = s.desc()
	colors, total, err := s.repo.QueryColors(where, page, perPage, sort)
	if err != nil {
		s.logger.Error(err.Error())
		return nil, 0, err
	}
	return colors, total, nil
}

func (s *Service) CreateColor(name *language.Languages, rgba string) error {
	value := &Color{}
	value.Id = utils.SnowFlakeIdString()
	value.Name = name
	value.Status = StatusActive
	value.Value  = rgba
	value.CreateTime = times.Timestamp()
	value.UpdateTime = times.Timestamp()
	value.DeleteTime = 0
	err := s.repo.CreateColor(value)
	if err != nil {
		s.logger.Error(err.Error())
		return err
	}
	return nil
}

func (s *Service) UpdateColor(id string, name *language.Languages, rgba string, status int) error {
	value, err := s.QueryColorById(id)
	if err != nil {
		s.logger.Error(err.Error())
		return err
	}
	status = s.Status(status)
	value.Value = rgba
	value.Name = name
	value.UpdateTime   = times.Timestamp()
	value.Status = s.Status(status)
	err = s.repo.UpdateColor(value)
	if err != nil {
		s.logger.Error(err.Error())
		return err
	}
	return nil
}

func (s *Service) DeleteColor(id string) error {
	value, err := s.QueryColorById(id)
	if err != nil {
		s.logger.Error(err.Error())
		return err
	}
	value.Status = StatusDeleted
	value.UpdateTime  = times.Timestamp()
	value.DeleteTime  = times.Timestamp()
	err = s.repo.UpdateColor(value)
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




