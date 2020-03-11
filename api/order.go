//author: richard
package api


//订单信息
type Order struct {
	Id  	string 		`json:"id"`					//订单Id
	User    *User  		`json:"user"` 	//用户信息快照
	State   string 	    `json:"state,"`				//订单现态
	NextState string  	`json:"nextState,"`		    //订单次态
	GoodsCount int 	    `json:"goodsCount"` //商品数量
	Address *Address    `json:"address"`	  //收货地址快照
	SubTotal   float64  `json:"subTotal"`   //商品小计
	Shipping   float64  `json:"shipping"`	  //运费
	Tax        float64  `json:"tax"`		  //税费
	Total      float64  `json:"total"`	      //实付金额
	OrderTime  int64 	`json:"orderTime"`  //下单时间
	CancelTime int64 	`json:"cancelTime"` //取消时间
	UpdateTime int64 	`json:"updateTime"` //更新时间
	ShippingTime int64  `json:"shippingTime"` //发货时间
	ReturnTime   int64  `json:"returnTime"`	//退货时间
	PayTime      int64  `json:"payTime"`	//支付时间
	Stocks     []Stock  `json:"stocks"`		//商品库存
	Pay    *CreditCard  `json:"pay"`					//支付卡

	SnapshotOrdered     bool `json:"snapshotOrdered"`     //是否已经下单
	SnapshotPendingPay  bool `json:"snapshotPendingPay"`  //是否待支付
	SnapshotPaying      bool `json:"snapshotPaying"`
	SnapshotPayed       bool `json:"snapshotPayed"`
	SnapshotPendingShip bool `json:"snapshotPendingShip"`
	SnapshotShipping    bool `json:"snapshotShipping"`
	SnapshotShipped     bool `json:"snapshotShipped"`
	SnapshotFinished    bool `json:"snapshotFinished"`
	SnapshotClosed      bool `json:"snapshotClosed"`
	SnapshotCanceled    bool `json:"snapshotCanceled"`
	//TODO
}

//信用卡
type CreditCard struct {
	Id     int      `json:"id"`
	Bin    string   `json:"bin"`  //卡片类型
	Number string 	`json:"number"`
	Expire string   `json:"expire"`
	CVV      string `json:"cvv"`
	ImageUrl string `json:"imageUrl"`
}

//支付交易流水


