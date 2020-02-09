//author: richard
package user

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/advancevillage/3rd/storages"
	"mms/api"
)

type Mongo struct {
	storage storages.Storage
}

func NewRepoMongo(storage storages.Storage) *Mongo {
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