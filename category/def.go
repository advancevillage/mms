//author: richard
package category

import (
	"github.com/advancevillage/3rd/logs"
	"mms/api"
)

const (
	Schema = "categories"
)

type ICategory interface {
	CreateCategory(category *api.Category) error
	UpdateCategory(category *api.Category) error
	QueryCategory(categoryId string) (*api.Category, error)
	QueryCategories(where map[string]interface{}, page int, perPage int, sort map[string]interface{}) ([]api.Category, int64, error)
}

type Service struct {
	repo   ICategory
	logger logs.Logs
}