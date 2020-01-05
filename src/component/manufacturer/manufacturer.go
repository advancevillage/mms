//author: richard
package manufacturer

import "mms/src/component/language"

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
	QueryManufacturers(where map[string]interface{}, page int, perPage int) ([]Manufacturer, error)
}

type Manufacturer struct {
	Id      string `json:"manufacturerId"`
	Contact string `json:"manufacturerContact"` //生产商联系人
	ContactPhone string `json:"contactPhone"`	//联系人联系电话
	ContactEmail string `json:"contactEmail"` 	//联系人邮箱
	Status  int `json:"manufacturerStatus"` 	//生产商状态
	CreateTime int64 `json:"manufacturerCreateTime"` //生产商录入记录时间
	UpdateTime int64 `json:"manufacturerUpdateTime"` //生产商更新时间
	DeleteTime int64 `json:"manufacturerDeleteTime"` //生产商移除系统时间
	Name    language.Languages `json:"manufacturerName"`	//生产商名称
	Address language.Languages `json:"manufacturerAddress"` //生产商地址
}