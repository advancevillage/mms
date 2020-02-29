//author: richard
package api


//订单信息
type Order struct {
	Id  	string 		`json:"id"`					//订单Id
	User    *User  		`json:"user,omitempty"` 	//用户信息快照
	State   string 	    `json:"state,"`				//订单现态
	NextState string  	`json:"nextState,"`		    //订单次态
	GoodsCount int 	    `json:"goodsCount,omitempty"` //商品数量
	Address *Address    `json:"address,omitempty"`	  //收货地址快照
	SubTotal   float64  `json:"subTotal,omitempty"`   //商品小计
	Shipping   float64  `json:"shipping,omitempty"`	  //运费
	Tax        float64  `json:"tax,omitempty"`		  //税费
	Total      float64  `json:"total,omitempty"`	  //实付金额
	OrderTime  int64 	`json:"orderTime,omitempty"`  //下单时间
	CancelTime int64 	`json:"cancelTime,omitempty"` //取消时间
	UpdateTime int64 	`json:"updateTime,omitempty"` //更新时间
	ShippingTime int64  `json:"shippingTime,omitempty"` //发货时间
	ReturnTime   int64  `json:"returnTime,omitempty"`	//退货时间
	PayTime      int64  `json:"payTime,omitempty"`		//支付时间
	Goods      []Goods  `json:"goods,omitempty"`		//商品明细
	Pay    *CreditCard  `json:"pay"`					//支付卡

	SnapshotOrdered     bool `json:"snapshotOrdered,omitempty"`     //是否已经下单
	SnapshotPendingPay  bool `json:"snapshotPendingPay,omitempty"`  //是否待支付
	SnapshotPaying      bool `json:"snapshotPaying,omitempty"`
	SnapshotPayed       bool `json:"snapshotPayed,omitempty"`
	SnapshotPendingShip bool `json:"snapshotPendingShip,omitempty"`
	SnapshotShipping    bool `json:"snapshotShipping,omitempty"`
	SnapshotShipped     bool `json:"snapshotShipped,omitempty"`
	SnapshotFinished    bool `json:"snapshotFinished,omitempty"`
	SnapshotClosed      bool `json:"snapshotClosed,omitempty"`
	SnapshotCanceled    bool `json:"snapshotCanceled,omitempty"`
	//TODO
}

//信用卡
type CreditCard struct {
	Bin    string   `json:"bin,omitempty"`  //卡片类型
	Number string 	`json:"number"`
	Expire string   `json:"expire"`
	CVV      string `json:"cvv"`
	ImageUrl string `json:"imageUrl"`
}

//支付交易流水


