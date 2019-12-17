//author: richard
package category

import (
	"encoding/json"
	"fmt"
	"github.com/advancevillage/3rd/storages"
)

type RepoEs7 struct {
	storage storages.Storage
}

func NewCategoryRepoEs7(storage storages.Storage) *RepoEs7 {
	return &RepoEs7{storage:storage}
}

func (s *RepoEs7) CreateCategory(cat *Category) error {
	body, err := json.Marshal(cat)
	if err != nil {
		return err
	}
	return s.storage.CreateStorageV2(Schema, fmt.Sprintf("%d", cat.CategoryId), body)
}

func (s *RepoEs7) DeleteCategory(cat ... *Category) error {
	var key = make([]string, 0, len(cat))
	for i := range cat {
		key = append(key, fmt.Sprintf("%d", cat[i].CategoryId))
	}
	return s.storage.DeleteStorageV2(Schema, key ...)
}

func (s *RepoEs7) UpdateCategory(cat *Category) error {
	body, err := json.Marshal(cat)
	if err != nil {
		return err
	}
	return s.storage.UpdateStorageV2(Schema, fmt.Sprintf("%d", cat.CategoryId), body)
}

func (s *RepoEs7) QueryCategory(categoryId int64) (*Category, error) {
	buf, err := s.storage.QueryStorageV2(Schema, fmt.Sprintf("%d", categoryId))
	if err != nil {
		return nil, err
	}
	cat := Category{}
	err = json.Unmarshal(buf, &cat)
	if err != nil {
		return nil, err
	}
	return &cat, nil
}


