//author: richard
package session

import (
	"github.com/advancevillage/3rd/caches"
)

type RepoRedis struct {
	storage caches.ICache
}

func NewSessionRepoRedis(storage caches.ICache) *RepoRedis {
	return &RepoRedis{storage:storage}
}

func (s *RepoRedis) CreateSession(o *Session) error {
	return s.storage.CreateCache(o.Key, o.Body, o.Timeout)
}

func (s *RepoRedis) DeleteSession(key ...string) error {
	return s.storage.DeleteCache(key ...)
}

func (s *RepoRedis) UpdateSession(o *Session) error {
	return s.storage.UpdateCache(o.Key, o.Body, o.Timeout)
}

func (s *RepoRedis) QuerySession(key string, timeout int) ([]byte, error) {
	buf, err := s.storage.QueryCache(key, timeout)
	if err != nil {
		return nil, err
	}
	return buf, nil
}

