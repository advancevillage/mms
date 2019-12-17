//author: richard
package size

import (
	"encoding/json"
	"fmt"
	"github.com/advancevillage/3rd/storages"
)

type RepoEs7 struct {
	storage storages.Storage
}

func NewSizeRepoEs7(storage storages.Storage) *RepoEs7 {
	return &RepoEs7{storage:storage}
}

func (s *RepoEs7) CreateSize(size *Size) error {
	body, err := json.Marshal(size)
	if err != nil {
		return err
	}
	return s.storage.CreateStorageV2(Schema, fmt.Sprintf("%d", size.SizeId), body)
}

func (s *RepoEs7) DeleteSize(size ... *Size) error {
	var key = make([]string, 0, len(size))
	for i := range size {
		key = append(key, fmt.Sprintf("%d", size[i].SizeId))
	}
	return s.storage.DeleteStorageV2(Schema, key ...)
}

func (s *RepoEs7) UpdateSize(size *Size) error {
	body, err := json.Marshal(size)
	if err != nil {
		return err
	}
	return s.storage.UpdateStorageV2(Schema, fmt.Sprintf("%d", size.SizeId), body)
}

func (s *RepoEs7) QuerySize(sizeId int64) (*Size, error) {
	buf, err := s.storage.QueryStorageV2(Schema, fmt.Sprintf("%d", sizeId))
	if err != nil {
		return nil, err
	}
	size := Size{}
	err = json.Unmarshal(buf, &size)
	if err != nil {
		return nil, err
	}
	return &size, nil
}


