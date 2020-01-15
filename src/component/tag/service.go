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

func (s *Service) QueryTagById(tagId string) (*Tag, error) {
	tag, err := s.repo.QueryTag(tagId)
	if err != nil {
		s.logger.Error(err.Error())
		return nil, err
	}
	return tag, nil
}

func (s *Service) QueryTags(status int, page int, perPage int) ([]Tag, int64, error) {
	where := make(map[string]interface{})
	where["tagStatus"] = s.Status(status)
	tags, total, err := s.repo.QueryTags(where, page, perPage)
	if err != nil {
		s.logger.Error(err.Error())
		return nil, 0, err
	}
	return tags, total, nil
}

func (s *Service) CreateTag(nameEn string) error {
	value := &Tag{}
	value.Id = utils.SnowFlakeIdString()
	value.Name.English = nameEn
	value.Status = StatusActive
	value.CreateTime = times.Timestamp()
	value.UpdateTime = times.Timestamp()
	value.DeleteTime = 0
	err := s.repo.CreateTag(value)
	if err != nil {
		s.logger.Error(err.Error())
		return err
	}
	return nil
}

func (s *Service) UpdateTag(id string, nameEn, nameCn string, status int) error {
	value, err := s.QueryTagById(id)
	if err != nil {
		s.logger.Error(err.Error())
		return err
	}
	status = s.Status(status)
	value.Name.English = nameEn
	value.Name.Chinese = nameCn
	value.UpdateTime   = times.Timestamp()
	value.Status = s.Status(status)
	err = s.repo.UpdateTag(value)
	if err != nil {
		s.logger.Error(err.Error())
		return err
	}
	return nil
}

func (s *Service) DeleteTag(id string) error {
	value, err := s.QueryTagById(id)
	if err != nil {
		s.logger.Error(err.Error())
		return err
	}
	value.Status = StatusDeleted
	value.UpdateTime  = times.Timestamp()
	value.DeleteTime  = times.Timestamp()
	err = s.repo.UpdateTag(value)
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
