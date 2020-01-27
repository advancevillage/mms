//author: richard
package color

import (
	"github.com/advancevillage/3rd/logs"
	"mms/api"
)

const (
	Schema = "colors"
)

type IColor interface {
	CreateColor(c *api.Color) error
	UpdateColor(c *api.Color) error
	QueryColor(id string) (*api.Color, error)
	QueryColors(where map[string]interface{}, page int, perPage int, sort map[string]interface{}) ([]api.Color, int64, error)
}

type Service struct {
	repo   IColor
	logger logs.Logs
}