//author: richard
package color

import (
	"encoding/json"
	"fmt"
	"github.com/advancevillage/3rd/storages"
)

type RepoEs7 struct {
	storage storages.Storage
}

func NewColorRepoEs7(storage storages.Storage) *RepoEs7 {
	return &RepoEs7{storage:storage}
}

func (s *RepoEs7) CreateColor(color *Color) error {
	body, err := json.Marshal(color)
	if err != nil {
		return err
	}
	return s.storage.CreateStorageV2(Schema, fmt.Sprintf("%d", color.ColorId), body)
}

func (s *RepoEs7) DeleteColor(color ... *Color) error {
	var key = make([]string, 0, len(color))
	for i := range color {
		key = append(key, fmt.Sprintf("%d", color[i].ColorId))
	}
	return s.storage.DeleteStorageV2(Schema, key ...)
}

func (s *RepoEs7) UpdateColor(color *Color) error {
	body, err := json.Marshal(color)
	if err != nil {
		return err
	}
	return s.storage.UpdateStorageV2(Schema, fmt.Sprintf("%d", color.ColorId), body)
}

func (s *RepoEs7) QueryColor(colorId int64) (*Color, error) {
	buf, err := s.storage.QueryStorageV2(Schema, fmt.Sprintf("%d", colorId))
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


