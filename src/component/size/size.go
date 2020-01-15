//author: richard
package size

import "mms/src/component/language"

const (
	Schema = "sizes"

	StatusActive  = 0x501
	StatusDeleted = 0x502
	StatusInvalid = 0x599
)

type ISize interface {
	CreateSize(size *Size) error
	UpdateSize(size *Size) error
	QuerySize(sizeId string) (*Size, error)
	QuerySizes(where map[string]interface{}, page int, perPage int) ([]Size, int64, error)
}

type Size struct {
	Id 	   string 	`json:"sizeId"`
	Status int 		`json:"sizeStatus"`
	CreateTime int64 `json:"sizeCreateTime"`
	UpdateTime int64 `json:"sizeUpdateTime"`
	DeleteTime int64 `json:"sizeDeleteTime"`
	Name   language.Languages `json:"sizeName"`
}
