//author: richard
package session

import (
	"fmt"
	"github.com/advancevillage/3rd/caches"
	"github.com/advancevillage/3rd/logs"
	"github.com/advancevillage/3rd/times"
	"github.com/advancevillage/3rd/utils"
)

type Service struct {
	repo   ISession
	logger logs.Logs
}

func NewSessionService(storage caches.ICache, logger logs.Logs) *Service {
	return &Service{repo:NewSessionRepoRedis(storage), logger:logger}
}

func (s *Service) CreateSession(body []byte) (string, error) {
	session := Session{}
	session.Key  = s.CreateSessionKey()
	session.Body = body
	session.Timeout = DefaultTimeout
	err := s.repo.CreateSession(&session)
	if err != nil {
		s.logger.Error(err.Error())
		return "", ErrorSessionCreateFail
	}
	return session.Key, nil
}

func (s *Service) CreateSessionKey() string {
	return fmt.Sprintf("%s%s%s", "sid", times.TimeFormatString(times.YYYYMMddHHmmss), utils.RandsNumberString(8))
}

