//author: richard
package image

import (
	"github.com/advancevillage/3rd/logs"
	"github.com/advancevillage/3rd/storages"
	"github.com/advancevillage/3rd/times"
	"github.com/advancevillage/3rd/utils"
)

type Service struct {
	repo   IImage
	logger logs.Logs
}

func NewImageService(storage storages.Storage, logger logs.Logs) *Service {
	return &Service{repo:NewImageRepoEs7(storage), logger:logger}
}

func (s *Service) QueryImage(imgId int64) (*Image, error) {
	img, err := s.repo.QueryImage(imgId)
	if err != nil {
		s.logger.Error(err.Error())
		return nil, err
	}
	return img, nil
}

func (s *Service) QueryImages(imgId ...int64) ([]*Image, error) {
	var length = len(imgId)
	var img = make([]*Image, 0, length)
	for i := 0; i < length; i++ {
		brd , err := s.repo.QueryImage(imgId[i])
		if err != nil {
			s.logger.Info(err.Error())
		} else {
			img = append(img, brd)
		}
	}
	return img, nil
}

func (s *Service) CreateImage(img *Image) error {
	img.ImageId = utils.SnowFlakeId()
	img.ImageCreateTime = times.Timestamp()
	img.ImageUpdateTime = times.Timestamp()
	img.ImageDeleteTime = 0
	err := s.repo.CreateImage(img)
	if err != nil {
		s.logger.Error(err.Error())
		return err
	}
	return nil
}


