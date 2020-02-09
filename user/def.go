//author: richard
package user


import (
	"github.com/advancevillage/3rd/caches"
	"github.com/advancevillage/3rd/logs"
	"mms/api"
)

const (
	Schema = "users"

	SHA1   = 160

	LoginTimeout    = 15 * 60 // 15'min
	RegisterTimeout = 30 * 60 // 30'min
)

type IUser interface {
	CreateUser(u *api.User) error
	QueryUserByName(username string) (*api.User, error)
	//QueryUserById(id string) (*api.User, error)
}


type Service struct {
	repo   IUser
	logger logs.Logs
	cache  caches.ICache
}
