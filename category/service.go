//author: richard
package category

import (
	"github.com/advancevillage/3rd/logs"
	"github.com/advancevillage/3rd/storages"
	"github.com/advancevillage/3rd/times"
	"github.com/advancevillage/3rd/utils"
	"mms/api"
)

func NewService(storage storages.Storage, logger logs.Logs) *Service {
	return &Service{repo:NewRepoMongo(storage), logger:logger}
}

func (s *Service) QueryCategoryById(id string) (*api.Category, error) {
	category, err := s.repo.QueryCategory(id)
	if err != nil {
		s.logger.Error(err.Error())
		return nil, err
	}
	return category, nil
}

func (s *Service) QueryCategories(page int, perPage int, level int) ([]api.Category, int64, error) {
	where := make(map[string]interface{})
	sort := make(map[string]interface{})
	where["level"] = level
	sort["createTime"] = s.desc()
	categories, total, err := s.repo.QueryCategories(where, page, perPage, sort)
	if err != nil {
		s.logger.Error(err.Error())
		return nil, 0, err
	}
	return categories, total, nil
}

func (s *Service) QueryChildCategories(id string) ([]api.Category, error) {
	category, err := s.QueryCategoryById(id)
	if err != nil {
		return nil, err
	}
	categories := make([]api.Category, 0, len(category.Child))
	for i := range category.Child {
		value, err := s.QueryCategoryById(category.Child[i])
		//无效 || 分类层级不是下一级
		if err != nil || category.Level != value.Level - 1 {
			continue
		} else {
			categories = append(categories, *value)
		}
	}
	return categories, nil
}

func (s *Service) CreateCategory(value *api.Category) error {
	value.Id = utils.SnowFlakeIdString()
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

func (s *Service) UpdateCategory(category *api.Category) error {
	value, err := s.QueryCategoryById(category.Id)
	if err != nil {
		s.logger.Error(err.Error())
		return err
	}
	value.Name = category.Name
	value.UpdateTime   = times.Timestamp()
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
	value.UpdateTime  = times.Timestamp()
	value.DeleteTime  = times.Timestamp()
	err = s.repo.UpdateCategory(value)
	if err != nil {
		s.logger.Error(err.Error())
		return err
	}
	return nil
}

func (s *Service) asc() int {
	return 1
}

func (s *Service) desc() int {
	return -1
}
