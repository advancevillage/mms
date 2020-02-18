//author: richard
package api


type User struct {
	Id			string  `json:"id,omitempty"`
	Username	string	`json:"username"`	//email
	Password 	string 	`json:"password"`	//sha1(password)
	Gender		int 	`json:"gender"` 	// 0 women 1 man
	CreateTime  int64 	`json:"createTime,omitempty"`
	UpdateTime  int64 	`json:"updateTime,omitempty"`
	DeleteTime  int64   `json:"deleteTime,omitempty"`
	Credit      []CreditCard  `json:"credit,omitempty"`
	Addr        []Address     `json:"addr,omitempty"`
}

type Login struct {
	Username	string	`json:"username"`	//email
	Password 	string 	`json:"password"`	//sha1(password)
	Timestamp   string   `json:"timestamp"`
	Token 		string  `json:"token"`
	Sign 		string  `json:"sign"`
}


type Register struct {
	Username	string	`json:"username"`	//email
	Password 	string 	`json:"password"`	//sha1(password)
	Timestamp   string  `json:"timestamp"`
	Token 		string  `json:"token"`
	Sign 		string  `json:"sign"`		//签名
	Gender		int 	`json:"gender"` 	//0 women 1 man
}

type Token struct {
	Username	string	`json:"username,omitempty"` //email
	Category    int 	`json:"category,omitempty"` //token 分类
}

type CreditCard struct {
	Number string `json:"number"`
	Expire string `json:"expire"`
	Custom string `json:"custom"`
	ImageUrl string `json:"imageUrl"`
	CVV      string `json:"cvv"`
	IsDefault  bool  `json:"isDefault"`
	CreateTime int64 `json:"createTime"`
	UpdateTime int64 `json:"updateTime"`
	DeleteTime int64 `json:"deleteTime"`
}

type Address struct {
	Country  string `json:"country"`   //国家
	Province string `json:"province"`  //省/州
	City 	 string `json:"city"`	   //城市
	Area     string `json:"area"`	   //区
	Street   string `json:"street"`    //街道及门牌号
	ZipCode  string `json:"zipCode"`   //邮编
	People   string `json:"people"`	   //收货人
	Email    string `json:"email"`	   //收货人邮箱
	Phone    string `json:"phone"`	   //收货人电话
	IsDefault  bool  `json:"isDefault"`
	CreateTime int64 `json:"createTime"`
	UpdateTime int64 `json:"updateTime"`
	DeleteTime int64 `json:"deleteTime"`
}
