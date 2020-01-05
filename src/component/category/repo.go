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

func (s *RepoMgo) CreateCategory(value *Category) error {
	body, err := json.Marshal(value)
	if err != nil {
		return err
	}
	return s.storage.CreateStorageV2(Schema, value.Id, body)
}

func (s *RepoMgo) UpdateCategory(value *Category) error {
	body, err := json.Marshal(value)
	if err != nil {
		return err
	}
	return s.storage.UpdateStorageV2(Schema, value.Id, body)
}

func (s *RepoMgo) QueryCategory(id string) (*Category, error) {
	buf, err := s.storage.QueryStorageV2(Schema, id)
	if err != nil {
		return nil, err
	}
	value := Category{}
	err = json.Unmarshal(buf, &value)
	if err != nil {
		return nil, err
	}
	return &value, nil
}

func (s *RepoMgo) QueryCategories(where map[string]interface{}, page int, perPage int) ([]Category, error) {
	items, err := s.storage.QueryStorageV3(Schema, where, perPage, page * perPage)
	if err != nil {
		return nil, err
	}
	values := make([]Category, 0, len(items))
	for i := range items {
		buf := items[i]
		value := Category{}
		err = json.Unmarshal(buf, &value)
		if err != nil {
			return nil, err
		} else {
			values = append(values, value)
		}
	}
	return values, nil
}


