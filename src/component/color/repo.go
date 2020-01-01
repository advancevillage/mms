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

func (s *RepoMgo) CreateColor(color *Color) error {
	body, err := json.Marshal(color)
	if err != nil {
		return err
	}
	return s.storage.CreateStorageV2(Schema, color.ColorId, body)
}

func (s *RepoMgo) DeleteColor(color ... *Color) error {
	var key = make([]string, 0, len(color))
	for i := range color {
		key = append(key, color[i].ColorId)
	}
	return s.storage.DeleteStorageV2(Schema, key ...)
}

func (s *RepoMgo) UpdateColor(color *Color) error {
	body, err := json.Marshal(color)
	if err != nil {
		return err
	}
	return s.storage.UpdateStorageV2(Schema, color.ColorId, body)
}

func (s *RepoMgo) QueryColor(colorId string) (*Color, error) {
	buf, err := s.storage.QueryStorageV2(Schema, colorId)
	if err != nil {
		return nil, err
	}
	color := Color{}
	err = json.Unmarshal(buf, &color)
	if err != nil {
		return nil, err
	}
	return &color, nil
}

