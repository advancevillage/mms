//author: richard
package brand

import (
	"encoding/json"
	"github.com/advancevillage/3rd/storages"
)

type RepoMgo struct {
	storage storages.Storage
}

func NewBrandRepoMgo(storage storages.Storage) *RepoMgo {
	return &RepoMgo{storage:storage}
}

func (s *RepoMgo) CreateBrand(brand *Brand) error {
	body, err := json.Marshal(brand)
	if err != nil {
		return err
	}
	return s.storage.CreateStorageV2(Schema, brand.Id, body)
}

func (s *RepoMgo) UpdateBrand(brand *Brand) error {
	body, err := json.Marshal(brand)
	if err != nil {
		return err
	}
	return s.storage.UpdateStorageV2(Schema, brand.Id, body)
}

func (s *RepoMgo) QueryBrand(brandId string) (*Brand, error) {
	buf, err := s.storage.QueryStorageV2(Schema, brandId)
	if err != nil {
		return nil, err
	}
	brd := Brand{}
	err = json.Unmarshal(buf, &brd)
	if err != nil {
		return nil, err
	}
	return &brd, nil
}

func (s *RepoMgo) QueryBrands(where map[string]interface{}, page int, perPage int) ([]Brand, int64, error) {
	items, total, err := s.storage.QueryStorageV3(Schema, where, perPage, page * perPage)
	if err != nil {
		return nil, 0, err
	}
	brands := make([]Brand, 0, len(items))
	for i := range items {
		buf := items[i]
		brd := Brand{}
		err = json.Unmarshal(buf, &brd)
		if err != nil {
			return nil, 0, err
		} else {
			brands = append(brands, brd)
		}
	}
	return brands, total, nil
}
