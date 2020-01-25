//author: richard
package size


import (
	"fmt"
	"github.com/advancevillage/3rd/logs"
	"github.com/advancevillage/3rd/storages"
	"github.com/advancevillage/3rd/times"
	"github.com/advancevillage/3rd/utils"
	"mms/src/language"
)

type Service struct {
	repo   ISize
	logger logs.Logs
}

func NewSizeService(storage storages.Storage, logger logs.Logs) *Service {
	return &Service{repo:NewSizeRepoMgo(storage), logger:logger}
}

func (s *Service) QuerySizeById(sizeId string) (*Size, error) {
	size, err := s.repo.QuerySize(sizeId)
	if err != nil {
		s.logger.Error(err.Error())
		return nil, err
	}
	return size, nil
}

func (s *Service) QuerySizes(status int, page int, perPage int, group string, lang string) ([]Size, int64, error) {
	where := make(map[string]interface{})
	sort := make(map[string]interface{})
	where["status"] = s.Status(status)
	where[fmt.Sprintf("%s.%s", "group", lang)]  = group
	sort["value"] = s.asc()
	//sort["createTime"] = s.desc()
	sizes, total, err := s.repo.QuerySizes(where, page, perPage, sort)
	if err != nil {
		s.logger.Error(err.Error())
		return nil, 0, err
	}
	return sizes, total, nil
}

func (s *Service) CreateSize(name string, group *language.Languages) error {
	value := &Size{}
	value.Id = utils.SnowFlakeIdString()
	value.Value = name
	value.Group = group
	value.Status = StatusActive
	value.CreateTime = times.Timestamp()
	value.UpdateTime = times.Timestamp()
	value.DeleteTime = 0
	err := s.repo.CreateSize(value)
	if err != nil {
		s.logger.Error(err.Error())
		return err
	}
	return nil
}

func (s *Service) UpdateSize(id string, name string, status int) error {
	value, err := s.QuerySizeById(id)
	if err != nil {
		s.logger.Error(err.Error())
		return err
	}
	value.Value = name
	value.UpdateTime   = times.Timestamp()
	value.Status = s.Status(status)
	err = s.repo.UpdateSize(value)
	if err != nil {
		s.logger.Error(err.Error())
		return err
	}
	return nil
}

func (s *Service) DeleteSize(id string) error {
	value, err := s.QuerySizeById(id)
	if err != nil {
		s.logger.Error(err.Error())
		return err
	}
	value.Status = StatusDeleted
	value.UpdateTime  = times.Timestamp()
	value.DeleteTime  = times.Timestamp()
	err = s.repo.UpdateSize(value)
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
