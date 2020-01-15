//author: richard
package style

import (
	"github.com/advancevillage/3rd/logs"
	"github.com/advancevillage/3rd/storages"
	"github.com/advancevillage/3rd/times"
	"github.com/advancevillage/3rd/utils"
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
	where["styleStatus"] = s.Status(status)
	styles, total, err := s.repo.QueryStyles(where, page, perPage)
	if err != nil {
		s.logger.Error(err.Error())
		return nil, 0, err
	}
	return styles, total, nil
}

func (s *Service) CreateStyle(name string, description string) error {
	value := &Style{}
	value.Id = utils.SnowFlakeIdString()
	value.Name.English = name
	value.Description.English = description
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

func (s *Service) UpdateStyle(id string, nameEn, nameCn string, descEn, descCn string, status int) error {
	value, err := s.QueryStyleById(id)
	if err != nil {
		s.logger.Error(err.Error())
		return err
	}
	status = s.Status(status)
	value.Name.English = nameEn
	value.Name.Chinese = nameCn
	value.Description.English = descEn
	value.Description.Chinese = descCn
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
