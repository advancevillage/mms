//author: richard
package manufacturer

import (
	"encoding/json"
	"fmt"
	"github.com/advancevillage/3rd/storages"
)

type RepoEs7 struct {
	storage storages.Storage
}

func NewManufacturerRepoEs7(storage storages.Storage) *RepoEs7 {
	return &RepoEs7{storage:storage}
}

func (s *RepoEs7) CreateManufacturer(mf *Manufacturer) error {
	body, err := json.Marshal(mf)
	if err != nil {
		return err
	}
	return s.storage.CreateStorageV2(Schema, fmt.Sprintf("%d", mf.ManufacturerId), body)
}

func (s *RepoEs7) DeleteManufacturer(mf ... *Manufacturer) error {
	var key = make([]string, 0, len(mf))
	for i := range mf {
		key = append(key, fmt.Sprintf("%d", mf[i].ManufacturerId))
	}
	return s.storage.DeleteStorageV2(Schema, key ...)
}

func (s *RepoEs7) UpdateManufacturer(mf *Manufacturer) error {
	body, err := json.Marshal(mf)
	if err != nil {
		return err
	}
	return s.storage.UpdateStorageV2(Schema, fmt.Sprintf("%d", mf.ManufacturerId), body)
}

func (s *RepoEs7) QueryManufacturer(mfId int64) (*Manufacturer, error) {
	buf, err := s.storage.QueryStorageV2(Schema, fmt.Sprintf("%d", mfId))
	if err != nil {
		return nil, err
	}
	mf := Manufacturer{}
	err = json.Unmarshal(buf, &mf)
	if err != nil {
		return nil, err
	}
	return &mf, nil
}
