//author: richard
package api


type Cart struct {
	Id        string `json:"id"`
	GoodsId   string `json:"goodsId"`
	ColorId   string `json:"colorId"`
	SizeId    string `json:"sizeId"`
	SizeValue string `json:"sizeValue"`
	ColorName *Languages `json:"colorName"`
	GoodsName *Languages `json:"goodsName"`
	GoodsStatus int      `json:"goodsStatus"`
	GoodsPrice  float64  `json:"goodsPrice"`
	Total       int 	 `json:"total"`
	CreateTime int64  `json:"createTime"`
	UpdateTime int64  `json:"updateTime"`
	DeleteTime int64  `json:"deleteTime"`
	FrontImage string `json:"frontImage"`
}
