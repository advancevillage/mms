//author: richard
package brand


import (
	"encoding/json"
	"fmt"
	"github.com/advancevillage/3rd/storages"
)

type RepoEs7 struct {
	storage storages.Storage
}

func NewBrandRepoEs7(storage storages.Storage) *RepoEs7 {
	return &RepoEs7{storage:storage}
}

func (s *RepoEs7) CreateBrand(brd *Brand) error {
	body, err := json.Marshal(brd)
	if err != nil {
		return err
	}
	return s.storage.CreateStorageV2(Schema, fmt.Sprintf("%d", brd.BrandId), body)
}

func (s *RepoEs7) DeleteBrand(brd ... *Brand) error {
	var key = make([]string, 0, len(brd))
	for i := range brd {
		key = append(key, fmt.Sprintf("%d", brd[i].BrandId))
	}
	return s.storage.DeleteStorageV2(Schema, key ...)
}

func (s *RepoEs7) UpdateBrand(brd *Brand) error {
	body, err := json.Marshal(brd)
	if err != nil {
		return err
	}
	return s.storage.UpdateStorageV2(Schema, fmt.Sprintf("%d", brd.BrandId), body)
}

func (s *RepoEs7) QueryBrand(BrandId int64) (*Brand, error) {
	buf, err := s.storage.QueryStorageV2(Schema, fmt.Sprintf("%d", BrandId))
	if err != nil {
		return nil, err
	}
	cat := Brand{}
	err = json.Unmarshal(buf, &cat)
	if err != nil {
		return nil, err
	}
	return &cat, nil
}
