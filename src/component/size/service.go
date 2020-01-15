//author: richard
package size


import (
	"github.com/advancevillage/3rd/logs"
	"github.com/advancevillage/3rd/storages"
	"github.com/advancevillage/3rd/times"
	"github.com/advancevillage/3rd/utils"
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

func (s *Service) QuerySizes(status int, page int, perPage int) ([]Size, int64, error) {
	where := make(map[string]interface{})
	where["sizeStatus"] = s.Status(status)
	sizes, total, err := s.repo.QuerySizes(where, page, perPage)
	if err != nil {
		s.logger.Error(err.Error())
		return nil, 0, err
	}
	return sizes, total, nil
}

func (s *Service) CreateSize(english string) error {
	value := &Size{}
	value.Id = utils.SnowFlakeIdString()
	value.Name.English = english
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

func (s *Service) UpdateSize(id string, english, chinese string, status int) error {
	value, err := s.QuerySizeById(id)
	if err != nil {
		s.logger.Error(err.Error())
		return err
	}
	status = s.Status(status)
	value.Name.English = english
	value.Name.Chinese = chinese
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