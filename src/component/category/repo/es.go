//author: richard
package repo

import (
	"github.com/advancevillage/3rd/storages"
	"mms/src/component/category"
)

type CategoryRepoEs7 struct {
	storage storages.Storage
}

func NewCategoryRepoEs7(storage storages.Storage) *CategoryRepoEs7 {
	return &CategoryRepoEs7{storage:storage}
}

func (c *CategoryRepoEs7) CreateCategory(cat *category.Category) error {
	return nil
}

func (c *CategoryRepoEs7) DeleteCategory(cat *category.Category) error {
	return nil
}

func (c *CategoryRepoEs7) UpdateCategory(cat *category.Category) error {
	return nil
}

