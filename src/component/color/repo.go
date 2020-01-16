//author: richard
package color

import (
	"encoding/json"
	"github.com/advancevillage/3rd/storages"
)

type RepoMgo struct {
	storage storages.Storage
}

func NewColorRepoMgo(storage storages.Storage) *RepoMgo {
	return &RepoMgo{storage:storage}
}

func (s *RepoMgo) CreateColor(value *Color) error {
	body, err := json.Marshal(value)
	if err != nil {
		return err
	}
	return s.storage.CreateStorageV2(Schema, value.Id, body)
}

func (s *RepoMgo) UpdateColor(value *Color) error {
	body, err := json.Marshal(value)
	if err != nil {
		return err
	}
	return s.storage.UpdateStorageV2(Schema, value.Id, body)
}

func (s *RepoMgo) QueryColor(id string) (*Color, error) {
	buf, err := s.storage.QueryStorageV2(Schema, id)
	if err != nil {
		return nil, err
	}
	value := Color{}
	err = json.Unmarshal(buf, &value)
	if err != nil {
		return nil, err
	}
	return &value, nil
}

func (s *RepoMgo) QueryColors(where map[string]interface{}, page int, perPage int, sort map[string]interface{}) ([]Color, int64, error) {
	items, total, err := s.storage.QueryStorageV3(Schema, where, perPage, page * perPage, sort)
	if err != nil {
		return nil, 0, err
	}
	values := make([]Color, 0, len(items))
	for i := range items {
		buf := items[i]
		value := Color{}
		err = json.Unmarshal(buf, &value)
		if err != nil {
			return nil, 0, err
		} else {
			values = append(values, value)
		}
	}
	return values, total, nil
}

