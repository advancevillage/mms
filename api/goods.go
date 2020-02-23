//author: richard
package api


type Goods struct {
	Id        string `json:"id"`
	GoodsId   string `json:"goodsId"`
	ColorId   string `json:"colorId"`
	SizeId    string `json:"sizeId"`
	SizeValue string `json:"sizeValue"`
	ColorName *Languages `json:"colorName"`
	GoodsName *Languages `json:"goodsName"`
	GoodsStatus int      `json:"goodsStatus"`
	GoodsPrice float64   `json:"goodsPrice"`
	Count      int 	  `json:"count"`
	FrontImage string `json:"frontImage"`
}
