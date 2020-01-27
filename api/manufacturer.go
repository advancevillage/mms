//author: richard
package api

type Manufacturer struct {
	Id      string `json:"id,omitempty"`
	Contact string `json:"contact,omitempty"` 		//生产商联系人
	Phone   string `json:"phone,omitempty"`			//联系人联系电话
	Email   string `json:"email,omitempty"` 		//联系人邮箱
	CreateTime int64 `json:"createTime,omitempty"`  //生产商录入记录时间
	UpdateTime int64 `json:"updateTime,omitempty"`  //生产商更新时间
	DeleteTime int64 `json:"deleteTime,omitempty"`  //生产商移除系统时间
	Name    *Languages `json:"name,omitempty"`		//生产商名称
	Address *Languages `json:"address,omitempty"`	//生产商地址
}
