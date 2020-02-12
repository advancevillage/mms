//author: richard
package api


type Cart struct {
	Id        string `json:"id,omitempty"`
	GoodsId   string `json:"goodsId,omitempty"`
	ColorId   string `json:"colorId,omitempty"`
	SizeId    string `json:"sizeId,omitempty"`
	SizeValue string `json:"sizeValue,omitempty"`
	ColorName *Languages `json:"colorName,omitempty"`
	GoodsName *Languages `json:"goodsName,omitempty"`
	GoodsStatus int      `json:"goodsStatus,omitempty"`
	GoodsPrice float64   `json:"goodsPrice,omitempty"`
	Count     int 	  `json:"count,omitempty"`
	CreateTime int64  `json:"createTime,omitempty"`
	UpdateTime int64  `json:"updateTime,omitempty"`
	DeleteTime int64  `json:"deleteTime,omitempty"`
	FrontImage string `json:"frontImage,omitempty"`
}
