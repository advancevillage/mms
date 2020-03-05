//author: richard
package internal

import (
	"mms/api"
	"mms/brand"
	"mms/category"
	"mms/color"
	"mms/config"
	"mms/goods"
	"mms/language"
	"mms/manufacturer"
	"mms/size"
)

type API interface {
	UpdateStock(args *api.Stocks, result *api.Stocks) error
}

type Service struct {
	configService *config.Service
	langService   *language.Service
	goodsService  *goods.Service
	colorService  *color.Service
	sizeService   *size.Service
	brandService  *brand.Service
	categoryService *category.Service
	manufacturerService *manufacturer.Service
}