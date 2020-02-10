//author: richard
package route

import (
	"github.com/advancevillage/3rd/https"
	"mms/brand"
	"mms/category"
	"mms/color"
	"mms/config"
	"mms/goods"
	"mms/language"
	"mms/manufacturer"
	"mms/size"
)

const (
	//商品 merchandise
	GoodsCode = 1000
	GoodsMsg  = "goods"
	//品牌 brands
	BrandCode = 1100
	BrandMsg  = "brand"
	//尺码 sizes
	SizeCode  = 1200
	SizeMsg   = "size"
	//颜色 color
	ColorCode = 1500
	ColorMsg  = "color"
	//分类 category
	CategoryCode = 1600
	CategoryMsg  = "category"
	//图片 image
	ImageCode = 1700
	ImageMsg  = "category"
	//生产商 manufacturer
	ManufacturerCode = 1800
	ManufacturerMsg  = "manufacturer"

	SnowFlakeIdLength = 18
)

const (
	BodyErrorCode = 11
	BodyErrorMsg  = "request query body error"

	ContextErrorCode = 12
	ContextErrorMsg  = "request body struct format error"

	CreateErrorCode = 13
	CreateErrorMsg  = "create error"
	QueryErrorCode  = 14
	QueryErrorMsg   = "query error"
	UpdateErrorCode = 15
	UpdateErrorMsg  = "update error"
	DeleteErrorCode = 16
	DeleteErrorMsg  = "delete error"
	IDErrorCode     = 20
	IDErrorMsg      = "id error"

	OperateSuccess = "operate success"
)

var router = func (api API) []https.Router {
	return []https.Router {
		{"GET", "/v1/services/version", api.version},
		{"GET", "/v1/services/ping", api.ping},
		//颜色
		{"POST", "/v1/colors", api.CreateColor},
		{"GET", "/v1/colors", api.QueryColors},
		{"GET", "/v1/colors/:pathId", api.QueryColor},
		{"PUT", "/v1/colors/:pathId", api.UpdateColor},
		{"DELETE", "/v1/colors/:pathId", api.DeleteColor},
		{"OPTIONS", "/v1/colors", api.ping},
		{"OPTIONS", "/v1/colors/:pathId", api.ping},
		//尺码
		{"POST", "/v1/sizes", api.CreateSize},
		{"GET", "/v1/sizes", api.QuerySizes},
		{"GET", "/v1/sizes/:pathId", api.QuerySize},
		{"PUT", "/v1/sizes/:pathId", api.UpdateSize},
		{"DELETE", "/v1/sizes/:pathId", api.DeleteSize},
		{"OPTIONS", "/v1/sizes", api.ping},
		{"OPTIONS", "/v1/sizes/:pathId", api.ping},
		//图片
		{"POST", "/v1/images", api.UploadImage},
		{"OPTIONS", "/v1/images", api.ping},
		//品牌
		{"GET", "/v1/brands", api.QueryBrands},
		{"POST", "/v1/brands", api.CreateBrand},
		{"PUT", "/v1/brands/:pathId", api.UpdateBrand},
		{"DELETE", "/v1/brands/:pathId", api.DeleteBrand},
		{"GET", "/v1/brands/:pathId", api.QueryBrand},
		{"OPTIONS", "/v1/brands", api.ping},
		{"OPTIONS", "/v1/brands/:pathId", api.ping},
		//分类
		{"GET", "/v1/categories", api.QueryCategories},
		{"POST", "/v1/categories", api.CreateCategory},
		{"PUT", "/v1/categories/:pathId", api.UpdateCategory},
		{"DELETE", "/v1/categories/:pathId", api.DeleteCategory},
		{"GET", "/v1/categories/:pathId", api.QueryCategory},
		{"GET", "/v1/categories/:pathId/categories", api.QueryChildCategories},
		{"OPTIONS", "/v1/categories", api.ping},
		{"OPTIONS", "/v1/categories/:pathId", api.ping},
		{"OPTIONS", "/v1/categories/:pathId/categories", api.ping},
		//生产商
		{"POST", "/v1/manufacturers", api.CreateManufacturer},
		{"GET", "/v1/manufacturers", api.QueryManufacturers},
		{"GET", "/v1/manufacturers/:pathId", api.QueryManufacturer},
		{"PUT", "/v1/manufacturers/:pathId", api.UpdateManufacturer},
		{"DELETE", "/v1/manufacturers/:pathId", api.DeleteManufacturer},
		{"OPTIONS", "/v1/manufacturers", api.ping},
		{"OPTIONS", "/v1/manufacturers/:pathId", api.ping},
		//商品
		{"POST", "/v1/goods", api.CreateGoods},
		{"GET", "/v1/goods", api.QueryGoods},
		{"GET", "/v1/goods/:pathId", api.QueryOneGoods},
		{"OPTIONS", "/v1/goods", api.ping},
		{"OPTIONS", "/v1/goods/:pathId", api.ping},
	}
}

type API interface {
	ping(ctx *https.Context)
	version(ctx *https.Context)
	//goods
	CreateGoods(ctx *https.Context)
	QueryGoods (ctx *https.Context)
	QueryOneGoods(ctx *https.Context)
	//color
	CreateColor(ctx *https.Context)
	QueryColors(ctx *https.Context)
	QueryColor (ctx *https.Context)
	UpdateColor(ctx *https.Context)
	DeleteColor(ctx *https.Context)
	//sizes
	CreateSize(ctx *https.Context)
	QuerySizes(ctx *https.Context)
	QuerySize (ctx *https.Context)
	UpdateSize(ctx *https.Context)
	DeleteSize(ctx *https.Context)
	//image
	UploadImage(ctx *https.Context)
	//brands
	QueryBrands(ctx *https.Context)
	CreateBrand(ctx *https.Context)
	UpdateBrand(ctx *https.Context)
	DeleteBrand(ctx *https.Context)
	QueryBrand (ctx *https.Context)
	//category
	QueryCategories(ctx *https.Context)
	QueryChildCategories(ctx *https.Context)
	CreateCategory(ctx *https.Context)
	UpdateCategory(ctx *https.Context)
	DeleteCategory(ctx *https.Context)
	QueryCategory (ctx *https.Context)
	//manufacturers
	CreateManufacturer(ctx *https.Context)
	QueryManufacturers(ctx *https.Context)
	QueryManufacturer (ctx *https.Context)
	UpdateManufacturer(ctx *https.Context)
	DeleteManufacturer(ctx *https.Context)
}

type httpError struct {
	Code    string `json:"code,omitempty"`
	Message string `json:"message,omitempty"`
}

type httpOk struct {
	Code    int `json:"code,omitempty"`
	Message string `json:"message,omitempty"`
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
