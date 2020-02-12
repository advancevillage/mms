//author: richard
package session

import "github.com/advancevillage/3rd/logs"

const (
	ExpireTime = 15 * 24 * 3600 //15day
)

type ISession interface {
	CreateSession(key string, value []byte) error
	QuerySession(key string) ([]byte, error)
	DeleteSession(key ...string) error
}


type Service struct {
	repo   ISession
	logger logs.Logs
}