//author: richard
package category

import (
	"encoding/json"
	"github.com/advancevillage/3rd/storages"
	"mms/api"
)

type Mongo struct {
	storage storages.Storage
}

func NewRepoMongo(storage storages.Storage) *Mongo {
	return &Mongo{storage:storage}
}

func (s *Mongo) CreateCategory(value *api.Category) error {
	body, err := json.Marshal(value)
	if err != nil {
		return err
	}
	return s.storage.CreateStorageV2(Schema, value.Id, body)
}

func (s *Mongo) UpdateCategory(value *api.Category) error {
	body, err := json.Marshal(value)
	if err != nil {
		return err
	}
	return s.storage.UpdateStorageV2(Schema, value.Id, body)
}

func (s *Mongo) QueryCategory(id string) (*api.Category, error) {
	buf, err := s.storage.QueryStorageV2(Schema, id)
	if err != nil {
		return nil, err
	}
	value := api.Category{}
	err = json.Unmarshal(buf, &value)
	if err != nil {
		return nil, err
	}
	return &value, nil
}

func (s *Mongo) QueryCategories(where map[string]interface{}, page int, perPage int, sort map[string]interface{}) ([]api.Category, int64, error) {
	items, total, err := s.storage.QueryStorageV3(Schema, where, perPage, page * perPage, sort)
	if err != nil {
		return nil, 0, err
	}
	values := make([]api.Category, 0, len(items))
	for i := range items {
		buf := items[i]
		value := api.Category{}
		err = json.Unmarshal(buf, &value)
		if err != nil {
			return nil, 0, err
		} else {
			values = append(values, value)
		}
	}
	return values, total, nil
}

