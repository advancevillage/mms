//author: richard
package brand

import (
	"encoding/json"
	"github.com/advancevillage/3rd/storages"
	"mms/api"
)

type Mongo struct {
	storage storages.Storage
}

func NewRepoMongo(storage storages.Storage) *Mongo {
	return &Mongo{storage:storage}
}

func (s *Mongo) CreateBrand(brand *api.Brand) error {
	body, err := json.Marshal(brand)
	if err != nil {
		return err
	}
	return s.storage.CreateStorageV2(Schema, brand.Id, body)
}

func (s *Mongo) UpdateBrand(brand *api.Brand) error {
	body, err := json.Marshal(brand)
	if err != nil {
		return err
	}
	return s.storage.UpdateStorageV2(Schema, brand.Id, body)
}

func (s *Mongo) QueryBrand(brandId string) (*api.Brand, error) {
	buf, err := s.storage.QueryStorageV2(Schema, brandId)
	if err != nil {
		return nil, err
	}
	brd := api.Brand{}
	err = json.Unmarshal(buf, &brd)
	if err != nil {
		return nil, err
	}
	return &brd, nil
}

func (s *Mongo) QueryBrands(where map[string]interface{}, page int, perPage int, sort map[string]interface{}) ([]api.Brand, int64, error) {
	items, total, err := s.storage.QueryStorageV3(Schema, where, perPage, page * perPage, sort)
	if err != nil {
		return nil, 0, err
	}
	brands := make([]api.Brand, 0, len(items))
	for i := range items {
		buf := items[i]
		brd := api.Brand{}
		err = json.Unmarshal(buf, &brd)
		if err != nil {
			return nil, 0, err
		} else {
			brands = append(brands, brd)
		}
	}
	return brands, total, nil
}
