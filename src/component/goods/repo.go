//author: richard
package goods

import (
	"encoding/json"
	"github.com/advancevillage/3rd/storages"
)

type RepoMgo struct {
	storage storages.Storage
}

func NewGoodsRepoMgo(storage storages.Storage) *RepoMgo {
	return &RepoMgo{storage:storage}
}

func (s *RepoMgo) CreateGoods(g *Goods) error {
	body, err := json.Marshal(g)
	if err != nil {
		return err
	}
	return s.storage.CreateStorageV2(Schema, g.GoodsId, body)
}

func (s *RepoMgo) DeleteGoods(g ... *Goods) error {
	var key = make([]string, 0, len(g))
	for i := range g {
		key = append(key, g[i].GoodsId)
	}
	return s.storage.DeleteStorageV2(Schema, key ...)
}

func (s *RepoMgo) UpdateGoods(g *Goods) error {
	body, err := json.Marshal(g)
	if err != nil {
		return err
	}
	return s.storage.UpdateStorageV2(Schema, g.GoodsId, body)
}

func (s *RepoMgo) QueryGoods(goodsId string) (*Goods, error) {
	buf, err := s.storage.QueryStorageV2(Schema, goodsId)
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

