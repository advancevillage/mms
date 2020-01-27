//author: richard
package api

type Size struct {
	Id 	   		string 	`json:"id,omitempty"`
	CreateTime  int64 	`json:"createTime,omitempty"`
	UpdateTime  int64   `json:"updateTime,omitempty"`
	DeleteTime  int64 	`json:"deleteTime,omitempty"`
	Value 		string 	`json:"value,omitempty"`
	Group  *Languages 	`json:"group,omitempty"`
}
