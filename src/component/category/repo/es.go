//author: richard
package repo

import (
	"github.com/advancevillage/3rd/storages"
	"mms/src/component/category"
)

type CategoryRepoEs7 struct {
	repo *storages.TES
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

