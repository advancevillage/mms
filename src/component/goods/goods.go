//author: richard
package goods

import (
	"mms/src/component/brand"
	"mms/src/component/category"
	"mms/src/component/color"
	"mms/src/component/image"
	"mms/src/component/manufacturer"
	"mms/src/component/size"
	"mms/src/component/tag"
)

const (
	Schema = "goods"

	StatusActived = 0x101
	StatusDeleted = 0x102
)

type IGoods interface {
	CreateGoods(*Goods) error
	DeleteGoods(...*Goods) error
	UpdateGoods(*Goods) error
	QueryGoods(string) (*Goods, error)
}

type Goods struct {
	GoodsId    string `json:"goodsId"` 	//商品标识
	GoodsTitle string `json:"goodsTitle"` //商品标题
	GoodsSummaryDescription  string `json:"goodsSummaryDescription"` //商品概要描述
	GoodsDetailedDescription string `json:"goodsDetailedDescription"` //商品详细描述
	GoodsTags []*tag.Tag `json:"goodsTags"` //商品标签
	GoodsKeywords []string `json:"goodsKeywords"` //商品关键字
	GoodsStatus int `json:"goodsStatus"` //商品状态
	GoodsIsIntegral bool  `json:"goodsIsIntegral"` //商品是否是积分商品
	GoodsIsReady bool  `json:"goodsIsReady"` //商品是否已经准备好
	GoodsRank int `json:"goodsRank"` //商品排名
	GoodsColors []*color.Color `json:"goodsColors"` //商品颜色
	GoodsSize  []*size.Size `json:"goodsSize"` //商品尺码
	GoodsOrigin string `json:"goodsOrigin"` //商品产地
	GoodsMaterial string `json:"goodsMaterial"` //商品材质
	GoodsManufacturers []*manufacturer.Manufacturer `json:"goodsManufacturers"` //商品生产商
	GoodsCategory []category.Category `json:"goodsCategory"` //商品分类
	GoodsBarCode string `json:"goodsBarCode"` //商品条形码
	GoodsBrand []brand.Brand `json:"goodsBrand"` //商品品牌
	GoodsCostPrice float64 `json:"goodsCostPrice"` //商品成本
	GoodsPrice float64 `json:"goodsPrice"` //商品标价
	GoodsStock int `json:"goodsStock"` //商品库存
	GoodsImages []*image.Image `json:"goodsImages"` //商品图片
	GoodsCreateTime int64 `json:"goodsCreateTime"` //商品创建时间
	GoodsUpdateTime int64 `json:"goodsUpdateTime"` //商品更新时间
	GoodsDeleteTime int64 `json:"goodsDeleteTime"` //商品删除时间
}