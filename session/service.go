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
	key := s.CreateKey(user.Username)
	value, err := json.Marshal(user)
	if err != nil {
		s.logger.Info(err.Error())
		return "", err
	}
	err = s.repo.CreateSession(key, value)
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

func (s *Service) CreateKey(username string) string {
	return fmt.Sprintf("%s:%s:%s:%s", "session", username, times.TimeFormatString(times.YYYYMMddHHmmss), utils.RandsNumberString(6))
}