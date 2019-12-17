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
	return &Service{repo:NewTagRepoEs7(storage), logger:logger}
}

func (s *Service) QueryTag(tagId int64) (*Tag, error) {
	tag, err := s.repo.QueryTag(tagId)
	if err != nil {
		s.logger.Error(err.Error())
		return nil, err
	}
	return tag, nil
}

func (s *Service) QueryTags(tagId ...int64) ([]*Tag, error) {
	var length = len(tagId)
	var tags = make([]*Tag, 0, length)
	for i := 0; i < length; i++ {
		brd , err := s.repo.QueryTag(tagId[i])
		if err != nil {
			s.logger.Info(err.Error())
		} else {
			tags = append(tags, brd)
		}
	}
	return tags, nil
}

func (s *Service) CreateTag(tag *Tag) error {
	tag.TagId = utils.SnowFlakeId()
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


