//author: richard
package api

type Goods struct {
	Id        string 		`json:"id"` 		//商品标识
	Name   	  *Languages 	`json:"name"`   	//商品概要描述
	Title     *Languages 	`json:"title"` 		//商品标题
	Status    int 	  		`json:"status"` 	//商品状态
	Rank 	  int 	   		`json:"rank"` 		//商品排名
	Origin    *Languages    `json:"origin"`    	//商品产地
	Price     float64   	`json:"price"` 					//商品售价 P = C + E * C
	Purchase  float64		`json:"purchase"`				//商品进价
	NewIn     float64  		`json:"newIn"`					//新品售价
	Sale      float64 		`json:"sale"`					//促销价
	Clearance float64		`json:"clearance"`				//清仓价
	Stocks    []Stock		`json:"stock"`					//库存
}


type Stock struct {
	Id      string `json:"id"`
	GoodsId string `json:"goodsId"`
	ColorId string `json:"colorId"`
	SizeId  string `json:"sizeId"`
	Total   int    `json:"total"`
	CreateTime int64 `json:"createTime"`
	UpdateTime int64 `json:"updateTime"`
	DeleteTime int64 `json:"deleteTime"`
	Version int   `json:"version"`
}
