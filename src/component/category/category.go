//author: richard
//@link: https://my.oschina.net/u/2410867/blog/1647801
package category

const (
	CategorySchema = "categories"
)

type ICategory interface {
	CreateCategory(*Category) error
	DeleteCategory(*Category) error
	UpdateCategory(*Category) error
}

type Category struct {
	CategoryId 	   int64  `json:"categoryId,omitempty"`	  		//分类标识
	CategoryName   string `json:"categoryName,omitempty"`   	//分类名称
	CategoryStatus int    `json:"categoryStatus,omitempty"`  	//分类状态
	ChildCategories  []int64 `json:"childCategories,omitempty"`	    //子分类
	ParentCategories []int64 `json:"parentCategories,omitempty"`	//父分类
	CategoryCreateTime int64 `json:"categoryCreateTime,omitempty"`	//分类创建时间
	CategoryUpdateTime int64 `json:"categoryUpdateTime,omitempty"`  //分类修复时间
	CategoryDeleteTime int64 `json:"categoryDeleteTime,omitempty"`  //分类删除时间
}
