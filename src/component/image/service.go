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
	return &Service{repo:NewImageRepoMgo(storage), logger:logger}
}

func (s *Service) QueryImage(imgId string) (*Image, error) {
	img, err := s.repo.QueryImage(imgId)
	if err != nil {
		s.logger.Error(err.Error())
		return nil, err
	}
	return img, nil
}

func (s *Service) QueryImages(imgIds ...string) ([]*Image, error) {
	var length = len(imgIds)
	var images = make([]*Image, 0, length)
	for i := 0; i < length; i++ {
		image , err := s.repo.QueryImage(imgIds[i])
		if err != nil {
			s.logger.Info(err.Error())
		} else {
			images = append(images, image)
		}
	}
	return images, nil
}

func (s *Service) CreateImage(img *Image) error {
	img.ImageId = utils.SnowFlakeIdString()
	img.CreateTime = times.Timestamp()
	img.UpdateTime = times.Timestamp()
	img.DeleteTime = 0
	err := s.repo.CreateImage(img)
	if err != nil {
		s.logger.Error(err.Error())
		return err
	}
	return nil
}


