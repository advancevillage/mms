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

func (s *Service) QueryCategoryById(categoryId int64) (*Category, error) {
	cat, err := s.repo.QueryCategory(categoryId)
	if err != nil {
		s.logger.Error(err.Error())
		return nil, err
	}
	return cat, nil
}

