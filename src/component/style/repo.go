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
	return s.storage.CreateStorageV2(Schema, style.StyleId, body)
}

func (s *RepoMgo) DeleteStyle(style ... *Style) error {
	var key = make([]string, 0, len(style))
	for i := range style {
		key = append(key, style[i].StyleId)
	}
	return s.storage.DeleteStorageV2(Schema, key ...)
}

func (s *RepoMgo) UpdateStyle(style *Style) error {
	body, err := json.Marshal(style)
	if err != nil {
		return err
	}
	return s.storage.UpdateStorageV2(Schema, style.StyleId, body)
}

func (s *RepoMgo) QueryStyle(styleId string) (*Style, error) {
	buf, err := s.storage.QueryStorageV2(Schema, styleId)
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
