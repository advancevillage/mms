//author: richard
package session

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/advancevillage/3rd/caches"
	"github.com/advancevillage/3rd/logs"
	"github.com/advancevillage/3rd/times"
	"github.com/advancevillage/3rd/utils"
	"mms/api"
)

func NewService(cache caches.ICache, logger logs.Logs) *Service {
	return &Service{
		repo: NewSessionRepo(cache),
		logger:logger,
	}
}

func (s *Service) CreateUserSession(user *api.User) (string, error) {
	if user == nil {
		return "", errors.New("user is nil")
	}
	key := s.CreateUserSessionKey(user.Username)
	value, err := json.Marshal(user)
	if err != nil {
		s.logger.Info(err.Error())
		return "", err
	}
	err = s.repo.CreateSession(key, value, ExpireTime)
	if err != nil {
		s.logger.Error(err.Error())
		return "", err
	}
	return key, nil
}

func (s *Service) QueryUserSession(key string) (*api.User, error) {
	buf, err := s.repo.QuerySession(key)
	if err != nil {
		s.logger.Info(err.Error())
		return nil, err
	}
	user := api.User{}
	err = json.Unmarshal(buf, &user)
	if err != nil {
		s.logger.Error(err.Error())
		return nil, err
	}
	return &user, nil
}

func (s *Service) CreateTidSession(user *api.User, value []byte) (string, error) {
	key := s.CreateTidSessionKey(user.Username)
    err := s.repo.CreateSession(key, value, 7 * 24 * 3600)
	if err != nil {
		s.logger.Error(err.Error())
		return "", err
	}
	return key, nil
}

func (s *Service) QueryTidSession(key string) (string, error) {
	buf, err := s.repo.QuerySession(key)
	if err != nil {
		s.logger.Info(err.Error())
		return "", err
	}
	return string(buf), nil
}

func (s *Service) DeleteTidSession(key string) error {
	err := s.repo.DeleteSession(key)
	if err != nil {
		s.logger.Warning(err.Error())
		return err
	}
	return nil
}

func (s *Service) UpdateTidSession(key string, value []byte) error {
	err := s.repo.UpdateSession(key, value, 7 * 24 * 3600)
	if err != nil {
		s.logger.Error(err.Error())
		return err
	}
	return nil
}

func (s *Service) CreateTidSessionKey(username string) string {
	return fmt.Sprintf("%s:%s:%s:%s", "orderpage", username, times.TimeFormatString(times.YYYYMMddHHmmss), utils.RandsNumberString(9))
}

func (s *Service) CreateUserSessionKey(username string) string {
	return fmt.Sprintf("%s:%s:%s:%s", "session", username, times.TimeFormatString(times.YYYYMMddHHmmss), utils.RandsNumberString(6))
}