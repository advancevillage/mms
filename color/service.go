//author: richard
package color

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

//业务校验
func (s *Service) CreateColor(value *api.Color) error {
	if value == nil {
		return errors.New("color is nil")
	}

	if len(value.RGB) == 0 {
		return errors.New("rgb is empty")
	}

	value.Id = utils.SnowFlakeIdString()
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

func (s *Service) QueryColorById(colorId string) (*api.Color, error) {
	color, err := s.repo.QueryColor(colorId)
	if err != nil {
		s.logger.Error(err.Error())
		return nil, err
	}
	return color, nil
}

func (s *Service) QueryColors(page int, perPage int) ([]api.Color, int64, error) {
	where := make(map[string]interface{})
	sort := make(map[string]interface{})
	sort["createTime"] = s.desc()
	colors, total, err := s.repo.QueryColors(where, page, perPage, sort)
	if err != nil {
		s.logger.Error(err.Error())
		return nil, 0, err
	}
	return colors, total, nil
}

func (s *Service) UpdateColor(color *api.Color) error {
	value, err := s.QueryColorById(color.Id)
	if err != nil {
		s.logger.Error(err.Error())
		return err
	}
	value.RGB  = color.RGB
	value.Name = color.Name
	value.UpdateTime   = times.Timestamp()
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
	value.UpdateTime  = times.Timestamp()
	value.DeleteTime  = times.Timestamp()
	err = s.repo.UpdateColor(value)
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
