//author: richard
package tag

import (
	"github.com/advancevillage/3rd/logs"
	"github.com/advancevillage/3rd/storages"
	"github.com/advancevillage/3rd/times"
	"github.com/advancevillage/3rd/utils"
)

type Service struct {
	repo   ITag
	logger logs.Logs
}

func NewTagService(storage storages.Storage, logger logs.Logs) *Service {
	return &Service{repo:NewTagRepoMgo(storage), logger:logger}
}

func (s *Service) QueryTag(tagId string) (*Tag, error) {
	tag, err := s.repo.QueryTag(tagId)
	if err != nil {
		s.logger.Error(err.Error())
		return nil, err
	}
	return tag, nil
}

func (s *Service) QueryTags(tagId ...string) ([]*Tag, error) {
	var length = len(tagId)
	var tags = make([]*Tag, 0, length)
	for i := 0; i < length; i++ {
		tag , err := s.repo.QueryTag(tagId[i])
		if err != nil {
			s.logger.Info(err.Error())
		} else {
			tags = append(tags, tag)
		}
	}
	return tags, nil
}

func (s *Service) CreateTag(tag *Tag) error {
	tag.TagId = utils.SnowFlakeIdString()
	tag.TagCreateTime = times.Timestamp()
	tag.TagUpdateTime = times.Timestamp()
	tag.TagDeleteTime = 0
	err := s.repo.CreateTag(tag)
	if err != nil {
		s.logger.Error(err.Error())
		return err
	}
	return nil
}


