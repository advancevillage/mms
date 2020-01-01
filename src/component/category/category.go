//author: richard
//@link: https://my.oschina.net/u/2410867/blog/1647801
package category

const (
	Schema = "categories"

	StatusActived = 0x201
	StatusDeleted = 0x202
)

type ICategory interface {
	CreateCategory(*Category) error
	DeleteCategory(...*Category) error
	UpdateCategory(*Category) error
	QueryCategory(string) (*Category, error)
}

type Category struct {
	CategoryId 	   string  `json:"categoryId"`	  		//分类标识
	CategoryName   string `json:"categoryName"`   		//分类名称
	CategoryStatus int    `json:"categoryStatus"`  		//分类状态
	CategoryLevel  int 	  `json:"categoryLevel"`		//分类级别
	ChildCategories  []string `json:"childCategories"`	//子分类
	ParentCategories []string `json:"parentCategories"`	//父分类
	CategoryCreateTime int64 `json:"categoryCreateTime"`  //分类创建时间
	CategoryUpdateTime int64 `json:"categoryUpdateTime"`  //分类修复时间
	CategoryDeleteTime int64 `json:"categoryDeleteTime"`  //分类删除时间
}
