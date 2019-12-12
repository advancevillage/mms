//author: richard
package route

const (
	CategoryCreateErrorCode = 1001
	CategoryCreateErrorMsg  = "分类创建失败"
	CategoryCreateOkMsg     = "分类创建成功"
)

type HttpError struct {
	Code  int 	`json:"code,omitempty"`
	Message string `json:"message,omitempty"`
}

type HttpOk struct {
	Code  int 	`json:"code,omitempty"`
	Message string `json:"message,omitempty"`
}

type ResponseVersion struct {
	Info string `json:"info,omitempty"`
}

type RequestCategory struct {
	CategoryName	string	`json:"categoryName,omitempty"`
	CategoryStatus  int 	`json:"categoryStatus,omitempty"`
	ChildCategories  []int64 `json:"childCategories,omitempty"`
	ParentCategories []int64 `json:"parentCategories,omitempty"`
}


