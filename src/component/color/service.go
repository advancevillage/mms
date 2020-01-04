//author: richard
package color

import (
	"github.com/advancevillage/3rd/logs"
	"github.com/advancevillage/3rd/storages"
	"github.com/advancevillage/3rd/times"
	"github.com/advancevillage/3rd/utils"
)

type Service struct {
	repo   IColor
	logger logs.Logs
}

func NewColorService(storage storages.Storage, logger logs.Logs) *Service {
	return &Service{repo:NewColorRepoMgo(storage), logger:logger}
}

func (s *Service) QueryColor(colorId string) (*Color, error) {
	color, err := s.repo.QueryColor(colorId)
	if err != nil {
		s.logger.Error(err.Error())
		return nil, err
	}
	return color, nil
}

func (s *Service) QueryColors(colorIds ...string) ([]*Color, error) {
	var length = len(colorIds)
	var colors = make([]*Color, 0, length)
	for i := 0; i < length; i++ {
		color , err := s.repo.QueryColor(colorIds[i])
		if err != nil {
			s.logger.Info(err.Error())
		} else {
			colors = append(colors, color)
		}
	}
	return colors, nil
}

func (s *Service) CreateColor(color *Color) error {
	color.ColorId = utils.SnowFlakeIdString()
	color.CreateTime = times.Timestamp()
	color.UpdateTime = times.Timestamp()
	color.DeleteTime = 0
	err := s.repo.CreateColor(color)
	if err != nil {
		s.logger.Error(err.Error())
		return err
	}
	return nil
}




