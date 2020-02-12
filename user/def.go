//author: richard
package user


import (
	"github.com/advancevillage/3rd/caches"
	"github.com/advancevillage/3rd/logs"
	"mms/api"
)

const (
	Schema = "users"
	CartSchema = "carts"

	SHA1   = 160

	LoginTimeout    = 15 * 60 // 15'min
	RegisterTimeout = 30 * 60 // 30'min
)

type IUser interface {
	CreateUser(u *api.User) error
	QueryUserByName(username string) (*api.User, error)
	CreateCart(user *api.User, cart *api.Cart) error
	QueryCart(user *api.User) ([]api.Cart, int64, error)
}


type Service struct {
	repo   IUser
	logger logs.Logs
	cache  caches.ICache
}
