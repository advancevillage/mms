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
	Timestamp   string   `json:"timestamp"`
	Token 		string  `json:"token"`
	Sign 		string  `json:"sign"`		//签名
	Gender		int 	`json:"gender"` 	// 0 women 1 man
}

type Token struct {
	Username	string	`json:"username,omitempty"` //email
	Category    int 	`json:"category,omitempty"` //token 分类
}

