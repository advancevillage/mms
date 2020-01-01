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

func (s *Service) QuerySize(sizeId string) (*Size, error) {
	size, err := s.repo.QuerySize(sizeId)
	if err != nil {
		s.logger.Error(err.Error())
		return nil, err
	}
	return size, nil
}

func (s *Service) QuerySizes(sizeIds ...string) ([]*Size, error) {
	var length = len(sizeIds)
	var sizes = make([]*Size, 0, length)
	for i := 0; i < length; i++ {
		size , err := s.repo.QuerySize(sizeIds[i])
		if err != nil {
			s.logger.Info(err.Error())
		} else {
			sizes = append(sizes, size)
		}
	}
	return sizes, nil
}

func (s *Service) CreateSize(size *Size) error {
	size.SizeId = utils.SnowFlakeIdString()
	size.SizeCreateTime = times.Timestamp()
	size.SizeUpdateTime = times.Timestamp()
	size.SizeDeleteTime = 0
	err := s.repo.CreateSize(size)
	if err != nil {
		s.logger.Error(err.Error())
		return err
	}
	return nil
}