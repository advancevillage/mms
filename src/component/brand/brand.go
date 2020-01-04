//author: richard
package brand

import "mms/src/component/language"

const (
	Schema = "brands"

	StatusActive  = 0x701
	StatusDeleted = 0x702
	StatusInvalid = 0x799
)

type IBrand interface {
	CreateBrand(*Brand) error
	UpdateBrand(*Brand) error
	QueryBrand(string) (*Brand, error)
	DeleteBrand(...string) error
	QueryBrands(where map[string]interface{}, page int, perPage int) ([]Brand, error)
}

type Brand struct {
	BrandId 	string 	`json:"brandId"`
	BrandStatus int 	`json:"brandStatus"`
	CreateTime int64 `json:"brandCreateTime"`
	UpdateTime int64 `json:"brandUpdateTime"`
	DeleteTime int64 `json:"brandDeleteTime"`
	BrandName 	language.Languages 	`json:"brandName"`
}
