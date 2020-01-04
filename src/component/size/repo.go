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

func (s *RepoMgo) CreateSize(size *Size) error {
	body, err := json.Marshal(size)
	if err != nil {
		return err
	}
	return s.storage.CreateStorageV2(Schema, size.SizeId, body)
}

func (s *RepoMgo) DeleteSize(size ... *Size) error {
	var key = make([]string, 0, len(size))
	for i := range size {
		key = append(key, size[i].SizeId)
	}
	return s.storage.DeleteStorageV2(Schema, key ...)
}

func (s *RepoMgo) UpdateSize(size *Size) error {
	body, err := json.Marshal(size)
	if err != nil {
		return err
	}
	return s.storage.UpdateStorageV2(Schema, size.SizeId, body)
}

func (s *RepoMgo) QuerySize(sizeId string) (*Size, error) {
	buf, err := s.storage.QueryStorageV2(Schema, sizeId)
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