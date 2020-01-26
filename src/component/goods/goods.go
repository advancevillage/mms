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
	Tags     []tag.Tag `json:"tags"` 				//商品标签
	Keywords []tag.Tag `json:"keywords"` 			//商品关键字
	Material []tag.Tag `json:"material"` 			//商品材质
	Status   int 	   `json:"status"` 				//商品状态
	Rank 	 int 	   `json:"rank"` 				//商品排名
	Colors   []color.Color `json:"colors"`          //商品颜色
	Size     []size.Size   `json:"size"` 			//商品尺码
	Origin   *language.Languages `json:"origin"`    //商品产地
	Category []category.Category `json:"category"`  //商品分类
	BarCode  string         `json:"barCode"`	    //商品条形码
	Brand    []brand.Brand  `json:"brand"`			//商品品牌
	Price     float64   `json:"price"` 				//商品售价 P = C + E * C
	Purchase  float64	`json:"purchase"`			//商品进价
	NewIn     float64  	`json:"newIn"`				//新品售价
	Sale      float64 	`json:"sale"`				//促销价
	Clearance float64	`json:"clearance"`			//清仓价
	Stock     []struct{								//库存
		ColorId string `json:"colorId"`
		SizeId  string `json:"sizeId"`
		Stock   int    `json:"stock"`
	} `json:"stock"`
	Images    []image.Image `json:"images"` 		//商品图片
	CreateTime int64 `json:"createTime"` 			//商品创建时间
	UpdateTime int64 `json:"updateTime"` 			//商品更新时间
	DeleteTime int64 `json:"deleteTime"` 			//商品删除时间
	Manufacturers []manufacturer.Manufacturer `json:"manufacturers"` //商品生产商
	Description *language.Languages `json:"Description"` 		     //商品详细描述
}