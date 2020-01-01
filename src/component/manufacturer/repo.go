//author: richard
package manufacturer

import (
	"encoding/json"
	"github.com/advancevillage/3rd/storages"
)

type RepoMgo struct {
	storage storages.Storage
}

func NewManufacturerRepoMgo(storage storages.Storage) *RepoMgo {
	return &RepoMgo{storage:storage}
}

func (s *RepoMgo) CreateManufacturer(mf *Manufacturer) error {
	body, err := json.Marshal(mf)
	if err != nil {
		return err
	}
	return s.storage.CreateStorageV2(Schema, mf.ManufacturerId, body)
}

func (s *RepoMgo) DeleteManufacturer(mf ... *Manufacturer) error {
	var key = make([]string, 0, len(mf))
	for i := range mf {
		key = append(key, mf[i].ManufacturerId)
	}
	return s.storage.DeleteStorageV2(Schema, key ...)
}

func (s *RepoMgo) UpdateManufacturer(mf *Manufacturer) error {
	body, err := json.Marshal(mf)
	if err != nil {
		return err
	}
	return s.storage.UpdateStorageV2(Schema, mf.ManufacturerId, body)
}

func (s *RepoMgo) QueryManufacturer(mfId string) (*Manufacturer, error) {
	buf, err := s.storage.QueryStorageV2(Schema, mfId)
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
