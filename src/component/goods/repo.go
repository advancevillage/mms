//author: richard
package goods

import (
	"encoding/json"
	"github.com/advancevillage/3rd/storages"
)

type RepoMgo struct {
	storage storages.Storage
}

func NewGoodsRepoMgo(storage storages.Storage) *RepoMgo {
	return &RepoMgo{storage:storage}
}

func (s *RepoMgo) CreateMerchandise(value *Goods) error {
	body, err := json.Marshal(value)
	if err != nil {
		return err
	}
	return s.storage.CreateStorageV2(Schema, value.Id, body)
}

func (s *RepoMgo) UpdateMerchandise(value *Goods) error {
	body, err := json.Marshal(value)
	if err != nil {
		return err
	}
	return s.storage.UpdateStorageV2(Schema, value.Id, body)
}

func (s *RepoMgo) QueryMerchandise(id string) (*Goods, error) {
	buf, err := s.storage.QueryStorageV2(Schema, id)
	if err != nil {
		return nil, err
	}
	value := Goods{}
	err = json.Unmarshal(buf, &value)
	if err != nil {
		return nil, err
	}
	return &value, nil
}

func (s *RepoMgo) QueryMerchandises(where map[string]interface{}, page int, perPage int, sort map[string]interface{}) ([]Goods, int64, error) {
	items, total, err := s.storage.QueryStorageV3(Schema, where, perPage, page * perPage, sort)
	if err != nil {
		return nil, 0, err
	}
	values := make([]Goods, 0, len(items))
	for i := range items {
		buf := items[i]
		value := Goods{}
		err = json.Unmarshal(buf, &value)
		if err != nil {
			return nil, 0, err
		} else {
			values = append(values, value)
		}
	}
	return values, total, nil
}

