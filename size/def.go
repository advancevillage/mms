//author: richard
package size

import (
	"github.com/advancevillage/3rd/logs"
	"mms/api"
)

const (
	Schema = "sizes"
)

type ISize interface {
	CreateSize(size *api.Size) error
	UpdateSize(size *api.Size) error
	QuerySize(sizeId string) (*api.Size, error)
	QuerySizes(where map[string]interface{}, page int, perPage int, sort map[string]interface{}) ([]api.Size, int64, error)
}


type Service struct {
	repo   ISize
	logger logs.Logs
}