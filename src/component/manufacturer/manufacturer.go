//author: richard
package manufacturer

const (
	Schema = "manufacturers"

	StatusActived = 0x701
	StatusDeleted = 0x702
)

type IManufacturer interface {
	CreateManufacturer(*Manufacturer) error
	DeleteManufacturer(...*Manufacturer) error
	UpdateManufacturer(*Manufacturer) error
	QueryManufacturer(int64) (*Manufacturer, error)
}

type Manufacturer struct {
	ManufacturerId   int64 `json:"manufacturerId"`
	ManufacturerName string `json:"manufacturerName"`	//生产商名称
	ManufacturerContact string `json:"manufacturerContact"` //生产商联系人
	ContactPhone string `json:"contactPhone"`	//联系人联系电话
	ContactEmail string `json:"contactEmail"` 	//联系人邮箱
	ManufacturerStatus int `json:"manufacturerStatus"`	//生产商状态
	ManufacturerAddress string `json:"manufacturerAddress"` //生产商地址
	ManufacturerCreateTime int64 `json:"manufacturerCreateTime"` //生产商录入记录时间
	ManufacturerUpdateTime int64 `json:"manufacturerUpdateTime"` //生产商编辑时间
	ManufacturerDeleteTime int64 `json:"manufacturerDeleteTime"` //生产商移除系统时间
}