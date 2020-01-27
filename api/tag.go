//author: richard
package api

type Tag struct {
	Id 		   string 	`json:"id,omitempty"`
	CreateTime int64 	`json:"createTime,omitempty"`
	UpdateTime int64 	`json:"updateTime,omitempty"`
	DeleteTime int64 	`json:"deleteTime,omitempty"`
	Name *Languages  	`json:"name,omitempty"`
}
