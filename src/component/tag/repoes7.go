//author: richard
package tag

import (
	"encoding/json"
	"fmt"
	"github.com/advancevillage/3rd/storages"
)

type RepoEs7 struct {
	storage storages.Storage
}

func NewTagRepoEs7(storage storages.Storage) *RepoEs7 {
	return &RepoEs7{storage:storage}
}

func (s *RepoEs7) CreateTag(tag *Tag) error {
	body, err := json.Marshal(tag)
	if err != nil {
		return err
	}
	return s.storage.CreateStorageV2(Schema, fmt.Sprintf("%d", tag.TagId), body)
}

func (s *RepoEs7) DeleteTag(tag ... *Tag) error {
	var key = make([]string, 0, len(tag))
	for i := range tag {
		key = append(key, fmt.Sprintf("%d", tag[i].TagId))
	}
	return s.storage.DeleteStorageV2(Schema, key ...)
}

func (s *RepoEs7) UpdateTag(tag *Tag) error {
	body, err := json.Marshal(tag)
	if err != nil {
		return err
	}
	return s.storage.UpdateStorageV2(Schema, fmt.Sprintf("%d", tag.TagId), body)
}

func (s *RepoEs7) QueryTag(tagId int64) (*Tag, error) {
	buf, err := s.storage.QueryStorageV2(Schema, fmt.Sprintf("%d", tagId))
	if err != nil {
		return nil, err
	}
	tag := Tag{}
	err = json.Unmarshal(buf, &tag)
	if err != nil {
		return nil, err
	}
	return &tag, nil
}

