//author: richard
package color

import "mms/src/component/language"

const (
	Schema = "colors"

	StatusActive  = 0x301
	StatusDeleted = 0x302
	StatusInvalid = 0x399
)

type IColor interface {
	CreateColor(*Color) error
	UpdateColor(*Color) error
	QueryColor(string) (*Color, error)
	QueryColors(where map[string]interface{}, page int, perPage int) ([]Color, error)
}

type Color struct {
	Id 	   string 	`json:"colorId"`
	Status int 		`json:"colorStatus"`
	Value  string   `json:"colorValue"` //色值 eg: #rgba(255,255,25,0)
	CreateTime int64 `json:"colorCreateTime"`
	UpdateTime int64 `json:"colorUpdateTime"`
	DeleteTime int64 `json:"colorDeleteTime"`
	Name   language.Languages 	`json:"colorName"`
}
