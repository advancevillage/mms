//author: richard
package route

const (
	//商品 merchandise
	MerchandiseCode = 1000
	MerchandiseMsg  = "merchandise"
	//品牌 brands
	BrandCode = 1100
	BrandMsg  = "brand"
	//尺码 sizes
	SizeCode  = 1200
	SizeMsg   = "size"
	//款式 styles
	StyleCode = 1300
	StyleMsg  = "style"
	//标签 tags
	TagCode   = 1400
	TagMsg    = "tag"
	//颜色 color
	ColorCode = 1500
	ColorMsg  = "color"
	//分类 category
	CategoryCode = 1600
	CategoryMsg  = "category"
	//图片 image
	ImageCode = 1700
	ImageMsg  = "category"
	//生产商 manufacturer
	ManufacturerCode = 1800
	ManufacturerMsg  = "manufacturer"

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
type Languages struct {
	English string `json:"english"`
	Chinese string `json:"chinese"`
}

type RequestCategory struct {
	NameEn	string	`json:"categoryNameEn"`
	NameCn	string	`json:"categoryNameCn,omitempty"`
	Status  int 	`json:"categoryStatus,omitempty"`
	Level   int 	`json:"categoryLevel,omitempty"`
	Child  []string `json:"childCategories,omitempty"`
	Parent []string `json:"parentCategories,omitempty"`
}

type RequestBrand struct {
	BrandName Languages `json:"brandName"`
	Status int    `json:"brandStatus,omitempty"`
}

type RequestSize struct {
	NameEn string `json:"sizeNameEn"`
	NameCn string `json:"sizeNameCn,omitempty"`
	Status  int   `json:"sizeStatus,omitempty"`
}

type RequestTag struct {
	NameEn string  `json:"tagNameEn"`
	NameCn string  `json:"tagNameCn,omitempty"`
	Status 	 int   `json:"tagStatus,omitempty"`
}

type RequestColor struct {
	NameEn string  `json:"colorNameEn"`
	NameCn string  `json:"colorNameCn,omitempty"`
	Rgba   string  `json:"rgba,omitempty"`
	Status 	 int   `json:"colorStatus,omitempty"`
}

type RequestStyle struct {
	NameEn  string `json:"styleNameEn"`
	NameCn  string `json:"styleNameCn,omitempty"`
	DescriptionEn string `json:"styleDescriptionEn"`
	DescriptionCn string `json:"styleDescriptionCn"`
	Status 	int    `json:"styleStatus,omitempty"`
}

type RequestManufacturer struct {
	Concat string `json:"concat"`
	Phone  string `json:"phone,omitempty"`
	Email  string `json:"email,omitempty"`
	NameEn string `json:"nameEn,omitempty"`
	NameCn string `json:"nameEn,omitempty"`
	AddressEn string `json:"addressEn,omitempty"`
	AddressCn string `json:"addressCn,omitempty"`
	Status 	int    `json:"styleStatus,omitempty"`
}

type RequestImage struct {
	DescEn  string `json:"imageDescriptionEn"`
	DescCn  string `json:"imageDescriptionCn,omitempty"`
	Status 	int    `json:"imageStatus,omitempty"`
	URL 	string `json:"imageUrl,omitempty"`
	Type    string `json:"imageType,omitempty"`
	Direction int  `json:"imageDirection,omitempty"`
	IsDefault bool `json:"isDefault,omitempty"`
}

type RequestGoods struct {
	TitleEn string `json:"titleEn"`
	DescEn  string `json:"descEn"`
	CostPrice float64 `json:"costPrice"`
}