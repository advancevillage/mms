//author: richard
package session

import (
	"github.com/advancevillage/3rd/caches"
)

type Storage struct {
	cache caches.ICache
}

func NewSessionRepo(cache caches.ICache) *Storage {
	return &Storage{cache: cache}
}

func (r *Storage) CreateSession(key string, value []byte, expire int) error {
	err := r.cache.CreateCache(key, value, expire)
	if err != nil {
		return err
	}
	return nil
}

func (r *Storage) QuerySession(key string) ([]byte, error) {
	value, err := r.cache.QueryCache(key, ExpireTime)
	if err != nil {
		return nil, err
	}
	return value, err
}

func (r *Storage) DeleteSession(key ...string) error {
	err := r.cache.DeleteCache(key...)
	if err != nil {
		return err
	}
	return nil
}

func (r *Storage) UpdateSession(key string, value []byte, expire int) error {
	err := r.cache.UpdateCache(key, value, expire)
	if err != nil {
		return err
	}
	return nil
}