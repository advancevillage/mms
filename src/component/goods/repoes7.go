//author: richard
package goods

import (
	"encoding/json"
	"fmt"
	"github.com/advancevillage/3rd/storages"
)

type RepoEs7 struct {
	storage storages.Storage
}

func NewGoodsRepoEs7(storage storages.Storage) *RepoEs7 {
	return &RepoEs7{storage:storage}
}

func (s *RepoEs7) CreateGoods(g *Goods) error {
	body, err := json.Marshal(g)
	if err != nil {
		return err
	}
	return s.storage.CreateStorageV2(Schema, fmt.Sprintf("%d", g.GoodsId), body)
}

func (s *RepoEs7) DeleteGoods(g ... *Goods) error {
	var key = make([]string, 0, len(g))
	for i := range g {
		key = append(key, fmt.Sprintf("%d", g[i].GoodsId))
	}
	return s.storage.DeleteStorageV2(Schema, key ...)
}

func (s *RepoEs7) UpdateGoods(g *Goods) error {
	body, err := json.Marshal(g)
	if err != nil {
		return err
	}
	return s.storage.UpdateStorageV2(Schema, fmt.Sprintf("%d", g.GoodsId), body)
}

func (s *RepoEs7) QueryGoods(gId int64) (*Goods, error) {
	buf, err := s.storage.QueryStorageV2(Schema, fmt.Sprintf("%d", gId))
	if err != nil {
		return nil, err
	}
	g := Goods{}
	err = json.Unmarshal(buf, &g)
	if err != nil {
		return nil, err
	}
	return &g, nil
}

