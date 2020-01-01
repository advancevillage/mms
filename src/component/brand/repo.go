//author: richard
package brand


import (
	"encoding/json"
	"fmt"
	"github.com/advancevillage/3rd/storages"
)

type RepoMgo struct {
	storage storages.Storage
}

func NewBrandRepoMgo(storage storages.Storage) *RepoMgo {
	return &RepoMgo{storage:storage}
}

func (s *RepoMgo) CreateBrand(brd *Brand) error {
	body, err := json.Marshal(brd)
	if err != nil {
		return err
	}
	return s.storage.CreateStorageV2(Schema, brd.BrandId, body)
}

func (s *RepoMgo) DeleteBrand(brd ... *Brand) error {
	var key = make([]string, 0, len(brd))
	for i := range brd {
		key = append(key, fmt.Sprintf("%d", brd[i].BrandId))
	}
	return s.storage.DeleteStorageV2(Schema, key ...)
}

func (s *RepoMgo) UpdateBrand(brd *Brand) error {
	body, err := json.Marshal(brd)
	if err != nil {
		return err
	}
	return s.storage.UpdateStorageV2(Schema, brd.BrandId, body)
}

func (s *RepoMgo) QueryBrand(BrandId int64) (*Brand, error) {
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
