//author: richard
package api

type Color struct {
	Id 	   string 	 `json:"id,omitempty"`
	RGB    string    `json:"rgb,omitempty"`
	CreateTime int64 `json:"createTime,omitempty"`
	UpdateTime int64 `json:"updateTime,omitempty"`
	DeleteTime int64 `json:"deleteTime,omitempty"`
	Name  *Languages `json:"name,omitempty"`
}
