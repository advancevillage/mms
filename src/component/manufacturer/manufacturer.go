//author: richard
package manufacturer

import "mms/src/component/language"

const (
	Schema = "manufacturers"

	StatusActived = 0x801
	StatusDeleted = 0x802
)

type IManufacturer interface {
	CreateManufacturer(*Manufacturer) error
	DeleteManufacturer(...*Manufacturer) error
	UpdateManufacturer(*Manufacturer) error
	QueryManufacturer(string) (*Manufacturer, error)
}

type Manufacturer struct {
	ManufacturerId   string `json:"manufacturerId"`
	ManufacturerContact string `json:"manufacturerContact"` //生产商联系人
	ContactPhone string `json:"contactPhone"`	//联系人联系电话
	ContactEmail string `json:"contactEmail"` 	//联系人邮箱
	ManufacturerStatus int `json:"manufacturerStatus"` //生产商状态
	CreateTime int64 `json:"manufacturerCreateTime"` //生产商录入记录时间
	UpdateTime int64 `json:"manufacturerUpdateTime"` //生产商编辑时间
	DeleteTime int64 `json:"manufacturerDeleteTime"` //生产商移除系统时间
	ManufacturerName language.Languages    `json:"manufacturerName"`	//生产商名称
	ManufacturerAddress language.Languages `json:"manufacturerAddress"` //生产商地址
}