//author: richard
package style

import "mms/src/language"

const (
	Schema = "styles"

	StatusActive  = 0x901
	StatusDeleted = 0x902
	StatusInvalid = 0x599
)

type IStyle interface {
	CreateStyle(style *Style) error
	UpdateStyle(style *Style) error
	QueryStyle(styleId string) (*Style, error)
	QueryStyles(where map[string]interface{}, page int, perPage int, sort map[string]interface{}) ([]Style, int64, error)
}

type Style struct {
	Id 	string 	`json:"id"`
	Status int 	`json:"status"`
	CreateTime  int64 `json:"createTime"`
	UpdateTime  int64 `json:"updateTime"`
	DeleteTime  int64 `json:"deleteTime"`
	Name *language.Languages `json:"name"`
	Description *language.Languages `json:"description"`
}
