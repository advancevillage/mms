//author: richard
package image

import "mms/src/component/language"

const (
	Schema = "images"

	StatusActive  = 0x401
	StatusDeleted = 0x402
	StatusInvalid = 0x499
)

type IImage interface {
	CreateImage(image *Image) error
	UpdateImage(image *Image) error
	QueryImage(imageId string) (*Image, error)
	QueryImages(where map[string]interface{}, page int, perPage int) ([]Image, int64, error)
}

type Image struct {
	Id 	    string 	`json:"imageId"`
	Url	    string 	`json:"imageUrl"`
	IsDefault bool  `json:"imageIsDefault"`
	Status     int  `json:"imageStatus"`
	CustomSize string   `json:"imageCustomSize"`
	CustomType string   `json:"imageCustomType"`
	CustomDirection int `json:"imageCustomDirection"`
	CreateTime int64 `json:"imageCreateTime"`
	UpdateTime int64 `json:"imageUpdateTime"`
	DeleteTime int64 `json:"imageDeleteTime"`
	Description language.Languages `json:"imageDescription"`
}
