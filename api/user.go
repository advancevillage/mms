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
