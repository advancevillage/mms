//author: richard
package size

import (
	"encoding/json"
	"github.com/advancevillage/3rd/storages"
)

type RepoMgo struct {
	storage storages.Storage
}

func NewSizeRepoMgo(storage storages.Storage) *RepoMgo {
	return &RepoMgo{storage:storage}
}

func (s *RepoMgo) CreateSize(value *Size) error {
	body, err := json.Marshal(value)
	if err != nil {
		return err
	}
	return s.storage.CreateStorageV2(Schema, value.Id, body)
}

func (s *RepoMgo) UpdateSize(value *Size) error {
	body, err := json.Marshal(value)
	if err != nil {
		return err
	}
	return s.storage.UpdateStorageV2(Schema, value.Id, body)
}

func (s *RepoMgo) QuerySize(id string) (*Size, error) {
	buf, err := s.storage.QueryStorageV2(Schema, id)
	if err != nil {
		return nil, err
	}
	value := Size{}
	err = json.Unmarshal(buf, &value)
	if err != nil {
		return nil, err
	}
	return &value, nil
}

func (s *RepoMgo) QuerySizes(where map[string]interface{}, page int, perPage int) ([]Size, error) {
	items, err := s.storage.QueryStorageV3(Schema, where, perPage, page * perPage)
	if err != nil {
		return nil, err
	}
	values := make([]Size, 0, len(items))
	for i := range items {
		buf := items[i]
		value := Size{}
		err = json.Unmarshal(buf, &value)
		if err != nil {
			return nil, err
		} else {
			values = append(values, value)
		}
	}
	return values, nil
}