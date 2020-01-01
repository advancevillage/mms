//author: richard
package category

import (
	"encoding/json"
	"github.com/advancevillage/3rd/storages"
)

type RepoMgo struct {
	storage storages.Storage
}

func NewCategoryRepoMgo(storage storages.Storage) *RepoMgo {
	return &RepoMgo{storage:storage}
}

func (s *RepoMgo) CreateCategory(cat *Category) error {
	body, err := json.Marshal(cat)
	if err != nil {
		return err
	}
	return s.storage.CreateStorageV2(Schema, cat.CategoryId, body)
}

func (s *RepoMgo) DeleteCategory(cat ... *Category) error {
	var key = make([]string, 0, len(cat))
	for i := range cat {
		key = append(key, cat[i].CategoryId)
	}
	return s.storage.DeleteStorageV2(Schema, key ...)
}

func (s *RepoMgo) UpdateCategory(cat *Category) error {
	body, err := json.Marshal(cat)
	if err != nil {
		return err
	}
	return s.storage.UpdateStorageV2(Schema, cat.CategoryId, body)
}

func (s *RepoMgo) QueryCategory(categoryId string) (*Category, error) {
	buf, err := s.storage.QueryStorageV2(Schema, categoryId)
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


