//author: richard
package manufacturer

import "mms/src/language"

const (
	Schema = "manufacturers"

	StatusActive  = 0x801
	StatusDeleted = 0x802
	StatusInvalid = 0x899
)

type IManufacturer interface {
	CreateManufacturer(m *Manufacturer) error
	UpdateManufacturer(m *Manufacturer) error
	QueryManufacturer(mId string) (*Manufacturer, error)
	QueryManufacturers(where map[string]interface{}, page int, perPage int, sort map[string]interface{}) ([]Manufacturer, int64, error)
}

type Manufacturer struct {
	Id      string `json:"id"`
	Contact string `json:"contact"` 	//生产商联系人
	ContactPhone string `json:"contactPhone"`	//联系人联系电话
	ContactEmail string `json:"contactEmail"` 	//联系人邮箱
	Status  int 	`json:"status"` 	 //生产商状态
	CreateTime int64 `json:"createTime"` //生产商录入记录时间
	UpdateTime int64 `json:"updateTime"` //生产商更新时间
	DeleteTime int64 `json:"deleteTime"` //生产商移除系统时间
	Name    *language.Languages `json:"name"`	//生产商名称
	Address *language.Languages `json:"address"` //生产商地址
}