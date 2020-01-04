//author: richard
package size

import "mms/src/component/language"

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
	SizeStatus int 		`json:"sizeStatus"`
	CreateTime int64 `json:"sizeCreateTime"`
	UpdateTime int64 `json:"sizeUpdateTime"`
	DeleteTime int64 `json:"sizeDeleteTime"`
	SizeName language.Languages `json:"sizeName"`
}
