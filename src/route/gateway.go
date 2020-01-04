//author: richard
package route

const (
	//品牌 brands
	BrandCode = 1100
	BrandMsg  = "brand"
	//尺码 sizes
	SizeCode  = 1200
	SizeMsg   = "size"
	//款式 styles
	StyleCode = 1300
	StyleMsg  = "style"

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

type RequestSize struct {
	SizeNameEn string `json:"sizeNameEn"`
	SizeNameCn string `json:"sizeNameCn,omitempty"`
	Status 		int   `json:"sizeStatus,omitempty"`
}

type RequestStyle struct {
	StyleNameEn string `json:"sizeNameEn"`
	StyleNameCn  string `json:"sizeNameCn,omitempty"`
	StyleDescriptionEn string `json:"styleDescriptionEn"`
	StyleDescriptionCn string `json:"styleDescriptionCn"`
	Status 		int   `json:"sizeStatus,omitempty"`
}