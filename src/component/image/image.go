//author: richard
package image

const (
	Schema = "images"

	StatusActived = 0x401
	StatusDeleted = 0x402
)

type IImage interface {
	CreateImage(*Image) error
	DeleteImage(...*Image) error
	UpdateImage(*Image) error
	QueryImage(string) (*Image, error)
}


type Image struct {
	ImageId 	string 	`json:"imageId"`
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
