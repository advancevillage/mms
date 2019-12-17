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
	return &Service{repo:NewColorRepoEs7(storage), logger:logger}
}

func (s *Service) QueryColor(colorId int64) (*Color, error) {
	color, err := s.repo.QueryColor(colorId)
	if err != nil {
		s.logger.Error(err.Error())
		return nil, err
	}
	return color, nil
}

func (s *Service) QueryColors(colorId ...int64) ([]*Color, error) {
	var length = len(colorId)
	var colors = make([]*Color, 0, length)
	for i := 0; i < length; i++ {
		brd , err := s.repo.QueryColor(colorId[i])
		if err != nil {
			s.logger.Info(err.Error())
		} else {
			colors = append(colors, brd)
		}
	}
	return colors, nil
}

func (s *Service) CreateColor(colorName string, colorStatus int, colorValue string) error {
	color := &Color{}
	color.ColorId = utils.SnowFlakeId()
	color.ColorName = colorName
	color.ColorStatus = colorStatus
	color.ColorValue  = colorValue
	color.ColorCreateTime = times.Timestamp()
	color.ColorUpdateTime = times.Timestamp()
	color.ColorDeleteTime = 0
	err := s.repo.CreateColor(color)
	if err != nil {
		s.logger.Error(err.Error())
		return err
	}
	return nil
}




