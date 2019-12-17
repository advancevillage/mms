//author: richard
package size

const (
	Schema = "sizes"

	StatusActived = 0x301
	StatusDeleted = 0x302
)

type ISize interface {
	CreateSize(*Size) error
	DeleteSize(...*Size) error
	UpdateSize(*Size) error
	QuerySize(int64) (*Size, error)
}

type Size struct {
	SizeId 	 int64 	`json:"sizeId"`
	SizeName string 	`json:"sizeName"`
	SizeStatus int 	`json:"sizeStatus"`
	SizeCreateTime int64 `json:"sizeCreateTime"`
	SizeUpdateTime int64 `json:"sizeUpdateTime"`
	SizeDeleteTime int64 `json:"sizeDeleteTime"`
}
