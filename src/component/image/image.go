//author: richard
package image

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
	QueryImages(where map[string]interface{}, page int, perPage int, sort map[string]interface{}) ([]Image, int64, error)
}

type Image struct {
	Id 	      string `json:"id"`
	Url	      string `json:"url"`
	Sequence  string `json:"sequence"`
	Direction string `json:"direction"`
	Status     int   `json:"status"`
	CreateTime int64 `json:"createTime"`
	UpdateTime int64 `json:"updateTime"`
	DeleteTime int64 `json:"deleteTime"`
}
