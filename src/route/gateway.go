//author: richard
package route

const (
	//品牌 brands
	BrandCode = 1100
	BrandMsg  = "brand"

	SnowFlakeIdLength = 18
)

const (
	RequestBodyErrorCode = 11
	RequestBodyErrorMsg  = "request query body error"

	BodyStructureErrorCode = 12
	BodyStructureErrorMsg  = "request body struct format error"

	CreateErrorCode = 13
	CreateErrorMsg  = "create error"
	QueryErrorCode  = 14
	QueryErrorMsg   = "query error"
	UpdateErrorCode = 15
	UpdateErrorMsg  = "update error"
	DeleteErrorCode = 16
	DeleteErrorMsg  = "delete error"
	IDErrorCode     = 20
	IDErrorMsg      = "id error"
)

const (
	OperateSuccess = "operate success"
)

type HttpError struct {
	Code    string `json:"code,omitempty"`
	Message string `json:"message,omitempty"`
}

type HttpOk struct {
	Code    int `json:"code,omitempty"`
	Message string `json:"message,omitempty"`
}

type ResponseVersion struct {
	Info string `json:"info,omitempty"`
}

type RequestCategory struct {
	CategoryName	string	`json:"categoryName,omitempty"`
	CategoryStatus  int 	`json:"categoryStatus,omitempty"`
	ChildCategories  []string `json:"childCategories,omitempty"`
	ParentCategories []string `json:"parentCategories,omitempty"`
}

type RequestBrand struct {
	BrandNameEn string `json:"brandNameEn"`
	BrandNameCn string `json:"brandNameCn,omitempty"`
	Status 		int    `json:"brandStatus,omitempty"`
}


