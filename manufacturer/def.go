//author: richard
package manufacturer

import (
	"github.com/advancevillage/3rd/logs"
	"mms/api"
)

const (
	Schema = "manufacturers"
)

type IManufacturer interface {
	CreateManufacturer(m *api.Manufacturer) error
	UpdateManufacturer(m *api.Manufacturer) error
	QueryManufacturer(mId string) (*api.Manufacturer, error)
	QueryManufacturers(where map[string]interface{}, page int, perPage int, sort map[string]interface{}) ([]api.Manufacturer, int64, error)
}


type Service struct {
	repo   IManufacturer
	logger logs.Logs
}
