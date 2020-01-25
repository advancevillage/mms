//author: richard
package color

import "mms/src/language"

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
	QueryColors(where map[string]interface{}, page int, perPage int, sort map[string]interface{}) ([]Color, int64, error)
}

type Color struct {
	Id 	   string 	`json:"id"`
	Status int 		`json:"status"`
	Value  string   `json:"rgb"` //色值 eg: #rgba(255,255,25,0)
	CreateTime int64 `json:"createTime"`
	UpdateTime int64 `json:"updateTime"`
	DeleteTime int64 `json:"deleteTime"`
	Name   *language.Languages 	`json:"name"`
}
