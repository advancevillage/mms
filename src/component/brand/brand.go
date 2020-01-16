//author: richard
package brand

import "mms/src/language"

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
	QueryBrands(where map[string]interface{}, page int, perPage int, sort map[string]interface{}) ([]Brand, int64, error)
}

type Brand struct {
	Id 	   string 	 `json:"id"`
	Status int 		 `json:"status"`
	CreateTime int64 `json:"createTime"`
	UpdateTime int64 `json:"updateTime"`
	DeleteTime int64 `json:"deleteTime"`
	Name   *language.Languages `json:"name"`
}
