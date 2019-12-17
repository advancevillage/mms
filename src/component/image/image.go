//author: richard
package image

const (
	Schema = "images"

	StatusActived = 0x601
	StatusDeleted = 0x602
)

type IImage interface {
	CreateImage(*Image) error
	DeleteImage(...*Image) error
	UpdateImage(*Image) error
	QueryImage(int64) (*Image, error)
}


type Image struct {
	ImageId 	int64 	`json:"colorId"`
	ImageUrl	string 	`json:"imageUrl"`
	ImageName 	string  `json:"imageName"`
	ImageIsDefault int  `json:"imageIsDefault"`
	ImageStatus    int  `json:"imageStatus"`
	ImageCustomSize string `json:"imageCustomSize"`
	ImageCustomType string `json:"imageCustomType"`
	ImageCustomDirection int `json:"imageCustomDirection"`
	ImageCreateTime int64 `json:"imageCreateTime"`
	ImageUpdateTime int64 `json:"imageUpdateTime"`
	ImageDeleteTime int64 `json:"imageDeleteTime"`
}
