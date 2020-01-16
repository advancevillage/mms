//author: richard
package goods

import (
	"mms/src/component/brand"
	"mms/src/component/category"
	"mms/src/component/color"
	"mms/src/component/image"
	"mms/src/component/manufacturer"
	"mms/src/component/size"
	"mms/src/component/style"
	"mms/src/component/tag"
	"mms/src/language"
)

const (
	Schema = "goods"

	StatusActive  = 0x101
	StatusDeleted = 0x102
	StatusInvalid = 0x199

	DefaultRank   = 9999

	Expected = 2.5
)

type IMerchandise interface {
	CreateMerchandise(goods *Goods) error
	UpdateMerchandise(goods *Goods) error
	QueryMerchandise (goodsId string) (*Goods, error)
	QueryMerchandises(where map[string]interface{}, page int, perPage int, sort map[string]interface{}) ([]Goods, int64, error)
}

type Goods struct {
	Id       string `json:"id"` 					//商品标识
	Title    *language.Languages `json:"title"` 	//商品标题
	Summary  *language.Languages `json:"summary"`   //商品概要描述
	Detailed *language.Languages `json:"detailed"`  //商品详细描述
	Tags     []tag.Tag `json:"tags"` 				//商品标签
	Keywords []language.Languages `json:"keywords"` //商品关键字
	Status   int `json:"status"` 					//商品状态
	IsIntegral bool  `json:"isIntegral"` 			//商品是否是积分商品
	IsReady    bool  `json:"isReady"` 				//商品是否已经准备好
	Rank 	   int 	 `json:"rank"` 					//商品排名
	Colors   []color.Color `json:"colors"`          //商品颜色
	Size     []size.Size   `json:"size"` 			//商品尺码
	Origin   *language.Languages `json:"origin"`    //商品产地
	Material *language.Languages `json:"material"`  //商品材质
	Manufacturers []manufacturer.Manufacturer `json:"manufacturers"` //商品生产商
	Category []category.Category `json:"category"`  //商品分类
	BarCode  string  `json:"barCode"`			    //商品条形码
	Brand    []brand.Brand  `json:"brand"`			//商品品牌
	CostPrice float64 `json:"costPrice"` 			//商品进价
	Price     float64 `json:"goodsPrice"` 			//商品售价 P = C + E * C
	Expected  float64 `json:"expected"`				//商品期望
	Stock 	  int     `json:"stock"` 			    //商品库存
	Images    []image.Image `json:"images"` 		//商品图片
	Style     []style.Style `json:"style"` 			//商品款式
	CreateTime int64 `json:"createTime"` 			//商品创建时间
	UpdateTime int64 `json:"updateTime"` 			//商品更新时间
	DeleteTime int64 `json:"deleteTime"` 			//商品删除时间
}