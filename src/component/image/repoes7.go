//author: richard
package image

import (
	"encoding/json"
	"fmt"
	"github.com/advancevillage/3rd/storages"
)

type RepoEs7 struct {
	storage storages.Storage
}

func NewImageRepoEs7(storage storages.Storage) *RepoEs7 {
	return &RepoEs7{storage:storage}
}

func (s *RepoEs7) CreateImage(img *Image) error {
	body, err := json.Marshal(img)
	if err != nil {
		return err
	}
	return s.storage.CreateStorageV2(Schema, fmt.Sprintf("%d", img.ImageId), body)
}

func (s *RepoEs7) DeleteImage(img ... *Image) error {
	var key = make([]string, 0, len(img))
	for i := range img {
		key = append(key, fmt.Sprintf("%d", img[i].ImageId))
	}
	return s.storage.DeleteStorageV2(Schema, key ...)
}

func (s *RepoEs7) UpdateImage(img *Image) error {
	body, err := json.Marshal(img)
	if err != nil {
		return err
	}
	return s.storage.UpdateStorageV2(Schema, fmt.Sprintf("%d", img.ImageId), body)
}

func (s *RepoEs7) QueryImage(imgId int64) (*Image, error) {
	buf, err := s.storage.QueryStorageV2(Schema, fmt.Sprintf("%d", imgId))
	if err != nil {
		return nil, err
	}
	img := Image{}
	err = json.Unmarshal(buf, &img)
	if err != nil {
		return nil, err
	}
	return &img, nil
}


