//author: richard
package color

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
	ColorName   string 	`json:"colorName"`
	ColorStatus int 	`json:"colorStatus"`
	ColorValue  string  `json:"colorValue"` //色值 eg: #rgba(255,255,25,0)
	ColorCreateTime int64 `json:"colorCreateTime"`
	ColorUpdateTime int64 `json:"colorUpdateTime"`
	ColorDeleteTime int64 `json:"colorDeleteTime"`
}