//author: richard
package brand

const (
	Schema = "brands"

	StatusActived = 0x201
	StatusDeleted = 0x202
)

type IBrand interface {
	CreateBrand(*Brand) error
	DeleteBrand(...*Brand) error
	UpdateBrand(*Brand) error
	QueryBrand(int64) (*Brand, error)
}


type Brand struct {
	BrandId 	int64 	`json:"brandId"`
	BrandName 	string 	`json:"brandName"`
	BrandStatus int 	`json:"brandStatus"`
	BrandCreateTime int64 `json:"brandCreateTime"`
	BrandUpdateTime int64 `json:"brandUpdateTime"`
	BrandDeleteTime int64 `json:"brandDeleteTime"`
}
