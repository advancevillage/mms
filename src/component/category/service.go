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
	return &Service{repo:NewCategoryRepoEs7(storage), logger:logger}
}

func (s *Service) QueryCategory(categoryId int64) (*Category, error) {
	cat, err := s.repo.QueryCategory(categoryId)
	if err != nil {
		s.logger.Error(err.Error())
		return nil, err
	}
	return cat, nil
}

func (s *Service) QueryCategories(categoryId ...int64) ([]*Category, error) {
	var length = len(categoryId)
	var cats = make([]*Category, 0, length)
	for i := 0; i < length; i++ {
		cat , err := s.repo.QueryCategory(categoryId[i])
		if err != nil {
			s.logger.Info(err.Error())
		} else {
			cats = append(cats, cat)
		}
	}
	return cats, nil
}

func (s *Service) CreateCategory(cat *Category) error {
	cat.CategoryId = utils.SnowFlakeId()
	cat.CategoryCreateTime = times.Timestamp()
	cat.CategoryUpdateTime = times.Timestamp()
	cat.CategoryDeleteTime = 0
	err := s.repo.CreateCategory(cat)
	if err != nil {
		s.logger.Error(err.Error())
		return err
	}
	return nil
}


