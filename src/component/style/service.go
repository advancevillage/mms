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

func (s *Service) QuerySize(styleId string) (*Style, error) {
	style, err := s.repo.QueryStyle(styleId)
	if err != nil {
		s.logger.Error(err.Error())
		return nil, err
	}
	return style, nil
}

func (s *Service) QuerySizes(styleIds ...string) ([]*Style, error) {
	var length = len(styleIds)
	var styles = make([]*Style, 0, length)
	for i := 0; i < length; i++ {
		style , err := s.repo.QueryStyle(styleIds[i])
		if err != nil {
			s.logger.Info(err.Error())
		} else {
			styles = append(styles, style)
		}
	}
	return styles, nil
}

func (s *Service) CreateSize(style *Style) error {
	style.StyleId = utils.SnowFlakeIdString()
	style.CreateTime = times.Timestamp()
	style.UpdateTime = times.Timestamp()
	style.DeleteTime = 0
	err := s.repo.CreateStyle(style)
	if err != nil {
		s.logger.Error(err.Error())
		return err
	}
	return nil
}
