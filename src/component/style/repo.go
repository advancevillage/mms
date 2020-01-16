//author: richard
package style

import (
	"encoding/json"
	"github.com/advancevillage/3rd/storages"
)

type RepoMgo struct {
	storage storages.Storage
}

func NewStyleRepoMgo(storage storages.Storage) *RepoMgo {
	return &RepoMgo{storage:storage}
}

func (s *RepoMgo) CreateStyle(style *Style) error {
	body, err := json.Marshal(style)
	if err != nil {
		return err
	}
	return s.storage.CreateStorageV2(Schema, style.Id, body)
}

func (s *RepoMgo) UpdateStyle(style *Style) error {
	body, err := json.Marshal(style)
	if err != nil {
		return err
	}
	return s.storage.UpdateStorageV2(Schema, style.Id, body)
}

func (s *RepoMgo) QueryStyle(id string) (*Style, error) {
	buf, err := s.storage.QueryStorageV2(Schema, id)
	if err != nil {
		return nil, err
	}
	style := Style{}
	err = json.Unmarshal(buf, &style)
	if err != nil {
		return nil, err
	}
	return &style, nil
}

func (s *RepoMgo) QueryStyles(where map[string]interface{}, page int, perPage int, sort map[string]interface{}) ([]Style, int64, error) {
	items, total, err := s.storage.QueryStorageV3(Schema, where, perPage, page * perPage, sort)
	if err != nil {
		return nil, 0, err
	}
	styles := make([]Style, 0, len(items))
	for i := range items {
		buf := items[i]
		style := Style{}
		err = json.Unmarshal(buf, &style)
		if err != nil {
			return nil, 0, err
		} else {
			styles = append(styles, style)
		}
	}
	return styles, total, nil
}
