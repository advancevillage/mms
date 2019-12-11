//author: richard
package category

import (
	"github.com/advancevillage/3rd/storages"
)

type RepoEs7 struct {
	storage storages.Storage
}

func NewCategoryRepoEs7(storage storages.Storage) *RepoEs7 {
	return &RepoEs7{storage:storage}
}

func (c *RepoEs7) CreateCategory(cat *Category) error {
	return nil
}

func (c *RepoEs7) DeleteCategory(cat *Category) error {
	return nil
}

func (c *RepoEs7) UpdateCategory(cat *Category) error {
	return nil
}

