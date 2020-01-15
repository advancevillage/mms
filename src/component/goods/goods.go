//author: richard
package goods

import (
	"mms/src/component/brand"
	"mms/src/component/category"
	"mms/src/component/color"
	"mms/src/component/image"
	"mms/src/component/language"
	"mms/src/component/manufacturer"
	"mms/src/component/size"
	"mms/src/component/style"
	"mms/src/component/tag"
)

const (
	Schema = "goods"

	StatusActive  = 0x101
	StatusDeleted = 0x102
	StatusInvalid = 0x199

	DefaultRank   = 9999
)

type IMerchandise interface {
	CreateMerchandise(goods *Goods) error
	UpdateMerchandise(goods *Goods) error
	QueryMerchandise (goodsId string) (*Goods, error)
	QueryMerchandises(where map[string]interface{}, page int, perPage int) ([]Goods, int64, error)
}

type Goods struct {
	Id    string `json:"goodsId"` 	//商品标识
	Title language.Languages `json:"goodsTitle"` //商品标题
	SummaryDescription  language.Languages `json:"goodsSummaryDescription"` //商品概要描述
	DetailedDescription language.Languages `json:"goodsDetailedDescription"` //商品详细描述
	Tags []tag.Tag `json:"goodsTags"` //商品标签
	Keywords []language.Languages `json:"goodsKeywords"` //商品关键字
	Status int `json:"goodsStatus"` //商品状态
	IsIntegral bool  `json:"goodsIsIntegral"` //商品是否是积分商品
	IsReady bool  `json:"goodsIsReady"` //商品是否已经准备好
	Rank int `json:"goodsRank"` //商品排名
	Colors []color.Color `json:"goodsColors"` //商品颜色
	Size  []size.Size `json:"goodsSize"` //商品尺码
	Origin language.Languages `json:"goodsOrigin"` //商品产地
	Material language.Languages `json:"goodsMaterial"` //商品材质
	Manufacturers []manufacturer.Manufacturer `json:"goodsManufacturers"` //商品生产商
	Category []category.Category `json:"goodsCategory"` //商品分类
	BarCode string `json:"goodsBarCode"` //商品条形码
	Brand []brand.Brand `json:"goodsBrand"` //商品品牌
	CostPrice float64 `json:"goodsCostPrice"` //商品成本
	Price float64 `json:"goodsPrice"` //商品标价
	Stock int `json:"goodsStock"` //商品库存
	Images []image.Image `json:"goodsImages"` //商品图片
	Style  []style.Style `json:"goodsStyle"`  //商品款式
	CreateTime int64 `json:"goodsCreateTime"` //商品创建时间
	UpdateTime int64 `json:"goodsUpdateTime"` //商品更新时间
	DeleteTime int64 `json:"goodsDeleteTime"` //商品删除时间
}