//author: richard
package user

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/advancevillage/3rd/storages"
	"mms/api"
	"strconv"
)

type Mongo struct {
	storage storages.StorageExd
}

func NewRepoMongo(storage storages.StorageExd) *Mongo {
	return &Mongo{storage:storage}
}

func (r *Mongo) QueryUserByName(username string) (*api.User, error) {
	//校验邮箱
	where := make(map[string]interface{})
	where["username"] = username
	items, total, err := r.storage.QueryStorageV3(Schema, where,1, 0, nil)
	if err != nil {
		return nil, err
	}
	if total > 1 || total <= 0 {
		return nil, errors.New(fmt.Sprintf("repeat %s", username))
	}
	user := api.User{}
	err = json.Unmarshal(items[total-1], &user)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *Mongo) CreateUser(user *api.User) error {
	buf, err := json.Marshal(user)
	if err != nil {
		return err
	}
	return r.storage.CreateStorageV2(Schema, user.Id, buf)
}

func (r *Mongo) CreateCart(user *api.User, cart *api.Cart) error {
	buf, err := json.Marshal(cart)
	if err != nil {
		return err
	}
	return r.storage.CreateStorageV2Exd(CartSchema, user.Id, cart.Id, buf)
}

func (r *Mongo) QueryCart(user *api.User) ([]api.Cart, int64, error) {
	where := make(map[string]interface{})
	where["deleteTime"] = map[string]interface{}{
		"$eq": 0,
	}
	sort  := make(map[string]interface{})
	sort["createTime"] = -1
	items, total, err := r.storage.SearchStorageV2Exd(CartSchema, user.Id, where, 99, 0, sort)
	if err != nil {
		return nil, 0, err
	}
	carts := make([]api.Cart, 0, total)
	for i := 0; i < len(items); i++ {
		cart := api.Cart{}
		err = json.Unmarshal(items[i], &cart)
		if err != nil {
			return nil, 0, err
		} else {
			carts = append(carts, cart)
			continue
		}
	}
	return carts, total, nil
}

func (r *Mongo) UpdateCart(user *api.User, cart *api.Cart) error {
	buf, err := json.Marshal(cart)
	if err != nil {
		return err
	}
	return r.storage.UpdateStorageV2Exd(CartSchema, user.Id, cart.Id, buf)
}

func (r *Mongo) QueryOneCart(user *api.User, cartId string) (*api.Cart, error) {
	buf, err := r.storage.QueryStorageV2Exd(CartSchema, user.Id, cartId)
	if err != nil {
		return nil, err
	}
	cart := api.Cart{}
	err = json.Unmarshal(buf, &cart)
	if err != nil {
		return nil, err
	}
	return &cart, nil
}

func (r *Mongo) QueryAddress(user *api.User) ([]api.Address, int64, error) {
	where := make(map[string]interface{})
	where["deleteTime"] = map[string]interface{}{
		"$eq": 0,
	}
	sort  := make(map[string]interface{})
	sort["createTime"] = -1
	items, total, err := r.storage.SearchStorageV2Exd(AddressSchema, user.Id, where, 99, 0, sort)
	if err != nil {
		return nil, 0, err
	}
	address := make([]api.Address, 0, total)
	for i := 0; i < len(address); i++ {
		addr := api.Address{}
		err = json.Unmarshal(items[i], &addr)
		if err != nil {
			return nil, 0, err
		} else {
			address = append(address, addr)
			continue
		}
	}
	return address, total, nil
}

func (r *Mongo) CreateAddress(user *api.User, address *api.Address) error {
	buf, err := json.Marshal(address)
	if err != nil {
		return err
	}
	return r.storage.CreateStorageV2Exd(AddressSchema, user.Id, strconv.FormatInt(address.Id, 10),  buf)
}