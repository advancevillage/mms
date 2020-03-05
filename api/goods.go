//author: richard
package api

type Goods struct {
	Id        string 		`json:"id,omitempty"` 			//商品标识
	Name   	  *Languages 	`json:"name,omitempty"`   		//商品概要描述
	Title     *Languages 	`json:"title,omitempty"` 		//商品标题
	Tags      []Tag 		`json:"tags,omitempty"` 		//商品标签
	Keywords  []Tag 		`json:"keywords,omitempty"` 	//商品关键字
	Materials []Tag 		`json:"materials,omitempty"` 	//商品材质
	Status    int 	  		`json:"status,omitempty"` 		//商品状态
	Rank 	  int 	   		`json:"rank,omitempty"` 		//商品排名
	Colors    []Color 		`json:"colors,omitempty"`       //商品颜色
	Size      []Size   		`json:"sizes,omitempty"` 		//商品尺码
	Origin    *Languages    `json:"origin,omitempty"`    	//商品产地
	Category  *Category 	`json:"category,omitempty"`  	//商品分类
	Brand     *Brand  		`json:"brand,omitempty"`		//商品品牌
	Price     float64   	`json:"price"` 					//商品售价 P = C + E * C
	Purchase  float64		`json:"purchase"`				//商品进价
	NewIn     float64  		`json:"newIn"`					//新品售价
	Sale      float64 		`json:"sale"`					//促销价
	Clearance float64		`json:"clearance"`				//清仓价
	Stock     []Stocks		`json:"stock"`					//库存
	Images    []Image 		`json:"images,omitempty"` 		//商品图片
	CreateTime int64 		`json:"createTime,omitempty"`	//商品创建时间
	UpdateTime int64 		`json:"updateTime,omitempty"`	//商品更新时间
	DeleteTime int64 		`json:"deleteTime,omitempty"` 	//商品删除时间
	Manufacturer  *Manufacturer `json:"manufacturer,omitempty"`  //商品生产商
	Description   *Languages 	`json:"description,omitempty"`   //商品详细描述
}


type Stocks struct {
	ColorId string `json:"colorId"`
	SizeId  string `json:"sizeId"`
	Stock   int    `json:"stock"`
}
