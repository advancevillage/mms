//author: richard
package image

import (
	"encoding/json"
	"github.com/advancevillage/3rd/storages"
)

type RepoMgo struct {
	storage storages.Storage
}

func NewImageRepoMgo(storage storages.Storage) *RepoMgo {
	return &RepoMgo{storage:storage}
}

func (s *RepoMgo) CreateImage(img *Image) error {
	body, err := json.Marshal(img)
	if err != nil {
		return err
	}
	return s.storage.CreateStorageV2(Schema, img.ImageId, body)
}

func (s *RepoMgo) DeleteImage(img ... *Image) error {
	var key = make([]string, 0, len(img))
	for i := range img {
		key = append(key, img[i].ImageId)
	}
	return s.storage.DeleteStorageV2(Schema, key ...)
}

func (s *RepoMgo) UpdateImage(img *Image) error {
	body, err := json.Marshal(img)
	if err != nil {
		return err
	}
	return s.storage.UpdateStorageV2(Schema, img.ImageId, body)
}

func (s *RepoMgo) QueryImage(imgId string) (*Image, error) {
	buf, err := s.storage.QueryStorageV2(Schema, imgId)
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


