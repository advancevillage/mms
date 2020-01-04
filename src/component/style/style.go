//author: richard
package style

import "mms/src/component/language"

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
	QueryStyles(where map[string]interface{}, page int, perPage int) ([]Style, error)
}

type Style struct {
	Id 	string 	`json:"styleId"`
	Status int 	`json:"styleStatus"`
	CreateTime  int64 `json:"styleCreateTime"`
	UpdateTime  int64 `json:"styleUpdateTime"`
	DeleteTime  int64 `json:"styleDeleteTime"`
	Name language.Languages `json:"styleName"`
	Description language.Languages `json:"styleDescription"`
}
