//author: richard
package size

import "mms/src/language"

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
	QuerySizes(where map[string]interface{}, page int, perPage int, sort map[string]interface{}) ([]Size, int64, error)
}

type Size struct {
	Id 	   string 	`json:"id"`
	Status int 		`json:"status"`
	CreateTime int64 `json:"createTime"`
	UpdateTime int64 `json:"updateTime"`
	DeleteTime int64 `json:"deleteTime"`
	Name   *language.Languages `json:"name"`
}
