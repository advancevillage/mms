//author: richard
package route

import (
	"mms/src/language"
)

const (
	//商品 merchandise
	MerchandiseCode = 1000
	MerchandiseMsg  = "merchandise"
	//品牌 brands
	BrandCode = 1100
	BrandMsg  = "brand"
	//尺码 sizes
	SizeCode  = 1200
	SizeMsg   = "size"
	//款式 styles
	StyleCode = 1300
	StyleMsg  = "style"
	//标签 tags
	TagCode   = 1400
	TagMsg    = "tag"
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
	RequestBodyErrorCode = 11
	RequestBodyErrorMsg  = "request query body error"

	BodyStructureErrorCode = 12
	BodyStructureErrorMsg  = "request body struct format error"

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
)

const (
	OperateSuccess = "operate success"
)

type HttpError struct {
	Code    string `json:"code,omitempty"`
	Message string `json:"message,omitempty"`
}

type HttpOk struct {
	Code    int `json:"code,omitempty"`
	Message string `json:"message,omitempty"`
}

type ResponseVersion struct {
	Info string `json:"info,omitempty"`
}

type RequestCategory struct {
	Status  int 	`json:"status"`
	Level   int 	`json:"level"`
	Child   string `json:"child"`
	Parent  string `json:"parent"`
	Name   language.Languages `json:"name"`
}

type RequestBrand struct {
	Status int    `json:"status"`
	Name   language.Languages `json:"name"`
}

type RequestSize struct {
	Status int   `json:"status"`
	Name   language.Languages `json:"name"`
}

type RequestTag struct {
	Status 	 int   `json:"status"`
	Name   language.Languages `json:"name"`
}

type RequestColor struct {
	Rgba   string  `json:"rgba"`
	Status int     `json:"status"`
	Name   language.Languages `json:"name"`
}

type RequestStyle struct {
	Status 	int    `json:"status"`
	Name   language.Languages `json:"name"`
	Description language.Languages `json:"description"`
}

type RequestManufacturer struct {
	Concat string `json:"concat"`
	Phone  string `json:"phone"`
	Email  string `json:"email"`
	Status 	int   `json:"status"`
	Name    language.Languages `json:"name"`
	Address language.Languages `json:"address"`
}

type RequestImage struct {
	Status 	int    `json:"status"`
	URL 	string `json:"url"`
	Type    string `json:"type"`
	Direction int  `json:"direction"`
	IsDefault bool `json:"isDefault"`
	Description language.Languages `json:"description"`
}

type RequestGoods struct {
	CostPrice float64 `json:"costPrice"`
	Title language.Languages `json:"title"`
	Description language.Languages `json:"description"`
}