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
	AddressSchema = "address"
	CreditSchema  = "credit"

	SHA1   = 160

	LoginTimeout    = 15 * 60 // 15'min
	RegisterTimeout = 30 * 60 // 30'min
)

type IUser interface {
	CreateUser(u *api.User) error
	QueryUserByName(username string) (*api.User, error)
	CreateCart(user *api.User, cart *api.Cart) error
	QueryCart(user *api.User) ([]api.Cart, int64, error)
	UpdateCart(user *api.User, cart *api.Cart) error
	QueryOneCart(user *api.User, cartId string) (*api.Cart, error)
	IAddress
	ICreditCard
}

type IAddress interface {
	QueryAddress(user *api.User) ([]api.Address, int64, error)
	CreateAddress(user *api.User, address *api.Address) error
}

type ICreditCard interface {
	QueryCreditCard(user *api.User) ([]api.CreditCard, int64, error)
	CreateCreditCard(user *api.User, credit *api.CreditCard) error
}


type Service struct {
	repo   IUser
	logger logs.Logs
	cache  caches.ICache
}
