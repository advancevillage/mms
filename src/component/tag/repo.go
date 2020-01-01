//author: richard
package tag

import (
	"encoding/json"
	"fmt"
	"github.com/advancevillage/3rd/storages"
)

type RepoMgo struct {
	storage storages.Storage
}

func NewTagRepoMgo(storage storages.Storage) *RepoMgo {
	return &RepoMgo{storage:storage}
}

func (s *RepoMgo) CreateTag(tag *Tag) error {
	body, err := json.Marshal(tag)
	if err != nil {
		return err
	}
	return s.storage.CreateStorageV2(Schema, tag.TagId, body)
}

func (s *RepoMgo) DeleteTag(tag ... *Tag) error {
	var key = make([]string, 0, len(tag))
	for i := range tag {
		key = append(key, fmt.Sprintf("%d", tag[i].TagId))
	}
	return s.storage.DeleteStorageV2(Schema, key ...)
}

func (s *RepoMgo) UpdateTag(tag *Tag) error {
	body, err := json.Marshal(tag)
	if err != nil {
		return err
	}
	return s.storage.UpdateStorageV2(Schema, tag.TagId, body)
}

func (s *RepoMgo) QueryTag(tagId string) (*Tag, error) {
	buf, err := s.storage.QueryStorageV2(Schema, tagId)
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

