//author: richard
package tag

import (
	"encoding/json"
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
	return s.storage.CreateStorageV2(Schema, tag.Id, body)
}

func (s *RepoMgo) UpdateTag(tag *Tag) error {
	body, err := json.Marshal(tag)
	if err != nil {
		return err
	}
	return s.storage.UpdateStorageV2(Schema, tag.Id, body)
}

func (s *RepoMgo) QueryTag(id string) (*Tag, error) {
	buf, err := s.storage.QueryStorageV2(Schema, id)
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

func (s *RepoMgo) QueryTags(where map[string]interface{}, page int, perPage int) ([]Tag, int64, error) {
	items, total, err := s.storage.QueryStorageV3(Schema, where, perPage, page * perPage)
	if err != nil {
		return nil, 0, err
	}
	tags := make([]Tag, 0, len(items))
	for i := range items {
		buf := items[i]
		tag := Tag{}
		err = json.Unmarshal(buf, &tag)
		if err != nil {
			return nil, 0, err
		} else {
			tags = append(tags, tag)
		}
	}
	return tags, total, nil
}
