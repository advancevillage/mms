//author: richard
package style

import (
	"github.com/advancevillage/3rd/logs"
	"github.com/advancevillage/3rd/storages"
	"github.com/advancevillage/3rd/times"
	"github.com/advancevillage/3rd/utils"
	"mms/src/language"
)

type Service struct {
	repo   IStyle
	logger logs.Logs
}

func NewStyleService(storage storages.Storage, logger logs.Logs) *Service {
	return &Service{repo:NewStyleRepoMgo(storage), logger:logger}
}

func (s *Service) QueryStyleById(styleId string) (*Style, error) {
	style, err := s.repo.QueryStyle(styleId)
	if err != nil {
		s.logger.Error(err.Error())
		return nil, err
	}
	return style, nil
}

func (s *Service) QueryStyles(status int, page int, perPage int) ([]Style, int64, error) {
	where := make(map[string]interface{})
	sort := make(map[string]interface{})
	where["status"] = s.Status(status)
	sort["createTime"] = s.desc()
	styles, total, err := s.repo.QueryStyles(where, page, perPage, sort)
	if err != nil {
		s.logger.Error(err.Error())
		return nil, 0, err
	}
	return styles, total, nil
}

func (s *Service) CreateStyle(name *language.Languages, description *language.Languages) error {
	value := &Style{}
	value.Id = utils.SnowFlakeIdString()
	value.Name = name
	value.Description = description
	value.Status = StatusActive
	value.CreateTime = times.Timestamp()
	value.UpdateTime = times.Timestamp()
	value.DeleteTime = 0
	err := s.repo.CreateStyle(value)
	if err != nil {
		s.logger.Error(err.Error())
		return err
	}
	return nil
}

func (s *Service) UpdateStyle(id string, name *language.Languages, description *language.Languages, status int) error {
	value, err := s.QueryStyleById(id)
	if err != nil {
		s.logger.Error(err.Error())
		return err
	}
	value.Name = name
	value.Description = description
	value.UpdateTime   = times.Timestamp()
	value.Status = s.Status(status)
	err = s.repo.UpdateStyle(value)
	if err != nil {
		s.logger.Error(err.Error())
		return err
	}
	return nil
}

func (s *Service) DeleteStyle(id string) error {
	value, err := s.QueryStyleById(id)
	if err != nil {
		s.logger.Error(err.Error())
		return err
	}
	value.Status = StatusDeleted
	value.UpdateTime  = times.Timestamp()
	value.DeleteTime  = times.Timestamp()
	err = s.repo.UpdateStyle(value)
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