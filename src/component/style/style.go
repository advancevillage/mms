//author: richard
package style

import "mms/src/component/language"

const (
	Schema = "styles"

	StatusActived = 0x901
	StatusDeleted = 0x902
)

type IStyle interface {
	CreateStyle(style *Style) error
	DeleteStyle(style ...*Style) error
	UpdateStyle(style *Style) error
	QueryStyle(styleId string) (*Style, error)
}

type Style struct {
	StyleId 	string 	`json:"styleId"`
	StyleStatus int 	`json:"styleStatus"`
	CreateTime  int64 `json:"styleCreateTime"`
	UpdateTime  int64 `json:"styleUpdateTime"`
	DeleteTime  int64 `json:"styleDeleteTime"`
	StyleName language.Languages `json:"styleName"`
	StyleDescription language.Languages `json:"styleDescription"`
}
