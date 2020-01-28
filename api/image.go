//author: richard
package api

type Image struct {
	Id 	      string `json:"id,omitempty"`
	Url	      string `json:"url"`
	Sequence  int 	 `json:"sequence"`
	Direction int 	 `json:"direction"`
	CreateTime int64 `json:"createTime,omitempty"`
	UpdateTime int64 `json:"updateTime,omitempty"`
	DeleteTime int64 `json:"deleteTime,omitempty"`
}
