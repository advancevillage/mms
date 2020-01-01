//author: richard
package size

const (
	Schema = "sizes"

	StatusActived = 0x501
	StatusDeleted = 0x502
)

type ISize interface {
	CreateSize(*Size) error
	DeleteSize(...*Size) error
	UpdateSize(*Size) error
	QuerySize(string) (*Size, error)
}

type Size struct {
	SizeId 	 string 	`json:"sizeId"`
	SizeName string 	`json:"sizeName"`
	SizeStatus int 	`json:"sizeStatus"`
	SizeCreateTime int64 `json:"sizeCreateTime"`
	SizeUpdateTime int64 `json:"sizeUpdateTime"`
	SizeDeleteTime int64 `json:"sizeDeleteTime"`
}
