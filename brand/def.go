//author: richard
package brand

import (
	"github.com/advancevillage/3rd/logs"
	"mms/api"
)

const (
	Schema = "brands"
)

type IBrand interface {
	CreateBrand(b *api.Brand) error
	UpdateBrand(b *api.Brand) error
	QueryBrand(id string) (*api.Brand, error)
	QueryBrands(where map[string]interface{}, page int, perPage int, sort map[string]interface{}) ([]api.Brand, int64, error)
}

type Service struct {
	repo   IBrand
	logger logs.Logs
}
