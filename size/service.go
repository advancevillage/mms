//author: richard
package size

import (
	"fmt"
	"github.com/advancevillage/3rd/logs"
	"github.com/advancevillage/3rd/storages"
	"github.com/advancevillage/3rd/times"
	"github.com/advancevillage/3rd/utils"
	"mms/api"
)

func NewService(storage storages.Storage, logger logs.Logs) *Service {
	return &Service{repo:NewRepoMongo(storage), logger:logger}
}

func (s *Service) QuerySizeById(sizeId string) (*api.Size, error) {
	size, err := s.repo.QuerySize(sizeId)
	if err != nil {
		s.logger.Error(err.Error())
		return nil, err
	}
	return size, nil
}

func (s *Service) QuerySizes(page int, perPage int, group string, lang string) ([]api.Size, int64, error) {
	where := make(map[string]interface{})
	sort := make(map[string]interface{})
	where[fmt.Sprintf("%s.%s", "group", lang)]  = group
	sort["value"] = s.asc()
	sizes, total, err := s.repo.QuerySizes(where, page, perPage, sort)
	if err != nil {
		s.logger.Error(err.Error())
		return nil, 0, err
	}
	return sizes, total, nil
}

func (s *Service) CreateSize(value *api.Size) error {
	value.Id = utils.SnowFlakeIdString()
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

func (s *Service) UpdateSize(size *api.Size) error {
	value, err := s.QuerySizeById(size.Id)
	if err != nil {
		s.logger.Error(err.Error())
		return err
	}
	value.Value = size.Value
	value.UpdateTime   = times.Timestamp()
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
	value.UpdateTime  = times.Timestamp()
	value.DeleteTime  = times.Timestamp()
	err = s.repo.UpdateSize(value)
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
