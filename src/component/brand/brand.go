//author: richard
package brand

const (
	Schema = "brands"

	StatusActived = 0x701
	StatusDeleted = 0x702
)

type IBrand interface {
	CreateBrand(*Brand) error
	DeleteBrand(...*Brand) error
	UpdateBrand(*Brand) error
	QueryBrand(int64) (*Brand, error)
}

type Brand struct {
	BrandId 	string 	`json:"brandId"`
	BrandName 	string 	`json:"brandName"`
	BrandStatus int 	`json:"brandStatus"`
	BrandCreateTime int64 `json:"brandCreateTime"`
	BrandUpdateTime int64 `json:"brandUpdateTime"`
	BrandDeleteTime int64 `json:"brandDeleteTime"`
}
