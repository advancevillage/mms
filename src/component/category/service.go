//author: richard
package category

//@note:
//@对象单一责任原则: 只需要导入repo && github.com/advancevillage/3rd/xxx

import (
	"github.com/advancevillage/3rd/logs"
	"github.com/advancevillage/3rd/storages"
)

type Service struct {
	repo   ICategory
	logger logs.Logs
}

func NewCategoryService(storage storages.Storage, logger logs.Logs) *Service {
	return &Service{repo:NewCategoryRepoEs7(storage), logger:logger}
}

func (s *Service) CreateCategory(cat *Category) error {
	err := s.repo.CreateCategory(cat)
	if err != nil {
		s.logger.Error(err.Error())
		return err
	}
	return nil
}
