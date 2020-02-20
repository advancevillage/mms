//author: richard
package api


type User struct {
	Id			string  `json:"id"`
	Username	string	`json:"username"`	//email
	Password 	string 	`json:"password"`	//sha1(password)
	Gender		int 	`json:"gender"` 	// 0 women 1 man
	CreateTime  int64 	`json:"createTime,omitempty"`
	UpdateTime  int64 	`json:"updateTime,omitempty"`
	DeleteTime  int64   `json:"deleteTime,omitempty"`
}

type Login struct {
	Username	string	`json:"username"`	//email
	Password 	string 	`json:"password"`	//sha1(password)
	Timestamp   string  `json:"timestamp"`
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
	Id 	   int64    `json:"id"`
	Number string `json:"number"`
	Expire string `json:"expire"`
	ImageUrl string `json:"imageUrl"`
	CVV      string `json:"cvv"`
	IsDefault  bool  `json:"isDefault"`
	CreateTime int64 `json:"createTime"`
	UpdateTime int64 `json:"updateTime"`
	DeleteTime int64 `json:"deleteTime"`
}

type Address struct {
	Id       int64  `json:"id"`
	FullName string `json:"fullName"`
	Country  string `json:"country"`   //国家
	Province string `json:"province"`  //省
	City     string `json:"city"`	   //城市
	Line1    string `json:"line1"`	   //Line 1
	Line2    string `json:"line2"`	   //Line 2
	ZipCode  string `json:"zipCode"`   //邮编
	Phone    string `json:"phone"`	   //收货人电话
	IsDefault  bool  `json:"isDefault"`
	CreateTime int64 `json:"createTime"`
	UpdateTime int64 `json:"updateTime"`
	DeleteTime int64 `json:"deleteTime"`
}
