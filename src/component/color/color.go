//author: richard
package color

import "mms/src/component/language"

const (
	Schema = "colors"

	StatusActived = 0x301
	StatusDeleted = 0x302
)

type IColor interface {
	CreateColor(*Color) error
	DeleteColor(...*Color) error
	UpdateColor(*Color) error
	QueryColor(string) (*Color, error)
}

type Color struct {
	ColorId 	string 	`json:"colorId"`
	ColorStatus int 	`json:"colorStatus"`
	ColorValue  string  `json:"colorValue"` //色值 eg: #rgba(255,255,25,0)
	CreateTime int64 `json:"colorCreateTime"`
	UpdateTime int64 `json:"colorUpdateTime"`
	DeleteTime int64 `json:"colorDeleteTime"`
	ColorName   language.Languages 	`json:"colorName"`
}
