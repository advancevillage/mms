//author: richard
package session

import (
	"errors"
)

const (
	DefaultTimeout	= 3 * 24 * 3600 // 3'day
)

var (
	ErrorSessionCreateFail = errors.New("session create fail")
)

type ISession interface {
	CreateSession(s *Session) error
	UpdateSession(s *Session) error
	DeleteSession(key ...string) error
	QuerySession(key string, timeout int) ([]byte, error)
}


type Session struct {
	Key  string		`json:"key"`
	Body []byte		`json:"body"`
	Timeout int		`json:"timeout"`
}


