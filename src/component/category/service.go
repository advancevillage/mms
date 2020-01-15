//author: richard
package category

//@note:
//@对象单一责任原则: 只需要导入repo && github.com/advancevillage/3rd/xxx

import (
	"github.com/advancevillage/3rd/logs"
	"github.com/advancevillage/3rd/storages"
	"github.com/advancevillage/3rd/times"
	"github.com/advancevillage/3rd/utils"
)

type Service struct {
	repo   ICategory
	logger logs.Logs
}

func NewCategoryService(storage storages.Storage, logger logs.Logs) *Service {
	return &Service{repo:NewCategoryRepoMgo(storage), logger:logger}
}

func (s *Service) QueryCategoryById(id string) (*Category, error) {
	color, err := s.repo.QueryCategory(id)
	if err != nil {
		s.logger.Error(err.Error())
		return nil, err
	}
	return color, nil
}

func (s *Service) QueryCategories(status int, page int, perPage int, level int) ([]Category, int64, error) {
	where := make(map[string]interface{})
	where["categoryStatus"] = s.Status(status)
	where["categoryLevel"] = level
	categories, total, err := s.repo.QueryCategories(where, page, perPage)
	if err != nil {
		s.logger.Error(err.Error())
		return nil, 0, err
	}
	return categories, total, nil
}

func (s *Service) CreateCategory(nameEn string, level int, child, parent []string) error {
	value := &Category{}
	value.Id = utils.SnowFlakeIdString()
	value.Name.English = nameEn
	value.Status = StatusActive
	value.Level  = s.Level(level)
	value.Child  = child
	value.Parent = parent
	value.CreateTime = times.Timestamp()
	value.UpdateTime = times.Timestamp()
	value.DeleteTime = 0
	err := s.repo.CreateCategory(value)
	if err != nil {
		s.logger.Error(err.Error())
		return err
	}
	return nil
}

func (s *Service) UpdateCategory(id string, nameEn, nameCn string, child, parent []string, status, level int) error {
	value, err := s.QueryCategoryById(id)
	if err != nil {
		s.logger.Error(err.Error())
		return err
	}
	value.Level = s.Level(level)
	value.Status = s.Status(status)
	value.Name.English = nameEn
	value.Name.Chinese = nameCn
	value.UpdateTime   = times.Timestamp()
	value.Child  = child
	value.Parent = parent
	err = s.repo.UpdateCategory(value)
	if err != nil {
		s.logger.Error(err.Error())
		return err
	}
	return nil
}

func (s *Service) DeleteCategory(id string) error {
	value, err := s.QueryCategoryById(id)
	if err != nil {
		s.logger.Error(err.Error())
		return err
	}
	value.Status = StatusDeleted
	value.UpdateTime  = times.Timestamp()
	value.DeleteTime  = times.Timestamp()
	err = s.repo.UpdateCategory(value)
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

func (s *Service) Level(level int) int {
	switch level {
	case 1:
		level = 1
	case 2:
		level = 2
	case 3:
		level = 3
	default:
		level = 1
	}
	return level
}
