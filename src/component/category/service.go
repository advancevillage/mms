//author: richard
package category

//@note:
//@对象单一责任原则: 只需要导入repo && github.com/advancevillage/3rd/xxx

import (
	"errors"
	"github.com/advancevillage/3rd/logs"
	"github.com/advancevillage/3rd/storages"
	"github.com/advancevillage/3rd/times"
	"github.com/advancevillage/3rd/utils"
	"mms/src/language"
)

type Service struct {
	repo   ICategory
	logger logs.Logs
}

func NewCategoryService(storage storages.Storage, logger logs.Logs) *Service {
	return &Service{repo:NewCategoryRepoMgo(storage), logger:logger}
}

func (s *Service) QueryCategoryById(id string) (*Category, error) {
	if len(id) != SnowFlakeIdLength {
		return nil, errors.New("id format error")
	}
	category, err := s.repo.QueryCategory(id)
	if err != nil {
		s.logger.Error(err.Error())
		return nil, err
	}
	return category, nil
}

func (s *Service) QueryCategories(status int, page int, perPage int, level int) ([]Category, int64, error) {
	where := make(map[string]interface{})
	sort := make(map[string]interface{})
	where["status"] = s.Status(status)
	where["level"] = level
	sort["createTime"] = s.desc()
	categories, total, err := s.repo.QueryCategories(where, page, perPage, sort)
	if err != nil {
		s.logger.Error(err.Error())
		return nil, 0, err
	}
	return categories, total, nil
}

func (s *Service) QueryChildCategories(id string) ([]Category, error) {
	category, err := s.QueryCategoryById(id)
	if err != nil {
		return nil, err
	}
	categories := make([]Category, 0, len(category.Child))
	for i := range category.Child {
		value, err := s.QueryCategoryById(category.Child[i])
		//无效 || 状态不是生效中 || 分类层级不是下一级
		if err != nil || value.Status != StatusActive  || category.Level != value.Level - 1 {
			continue
		} else {
			categories = append(categories, *value)
		}
	}
	return categories, nil
}

func (s *Service) CreateCategory(name *language.Languages, level int, child, parent string) error {
	value := &Category{}
	value.Id = utils.SnowFlakeIdString()
	value.Name = name
	value.Status = StatusActive
	value.Level  = s.Level(level)
	value.CreateTime = times.Timestamp()
	value.UpdateTime = times.Timestamp()
	value.DeleteTime = 0

	children, err := s.QueryCategoryById(child)
	if err != nil {
		value.Child = make([]string, 0)
	} else {
		//value.id's children is child
		value.Child = append(value.Child, child)
		//child's parent is value.id
		children.Parent = append(children.Parent, value.Id)
	}

	parents, err := s.QueryCategoryById(parent)
	if err != nil {
		value.Parent = make([]string, 0)
	} else {
		//value.id's parents is parent
		value.Parent = append(value.Parent, parent)
		//parents's child is value.id
		parents.Child = append(parents.Child, value.Id)
	}
	err = s.repo.CreateCategory(value)
	if err != nil {
		s.logger.Error(err.Error())
		return err
	}

	err = s.repo.UpdateCategory(children)
	if err != nil {
		s.logger.Info(err.Error())
	}

	err = s.repo.UpdateCategory(parents)
	if err != nil {
		s.logger.Info(err.Error())
	}
	return nil
}

func (s *Service) UpdateCategory(id string, name *language.Languages, status int) error {
	value, err := s.QueryCategoryById(id)
	if err != nil {
		s.logger.Error(err.Error())
		return err
	}
	value.Status = s.Status(status)
	value.Name = name
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

func (s *Service) asc() int {
	return 1
}

func (s *Service) desc() int {
	return -1
}
