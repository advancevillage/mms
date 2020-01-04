//author: richard
//@link: https://my.oschina.net/u/2410867/blog/1647801
package category

import "mms/src/component/language"

const (
	Schema = "categories"

	StatusActived = 0x201
	StatusDeleted = 0x202
)

type ICategory interface {
	CreateCategory(category *Category) error
	DeleteCategory(category ...*Category) error
	UpdateCategory(category *Category) error
	QueryCategory(string) (*Category, error)
}

type Category struct {
	CategoryId 	   string  `json:"categoryId"`	  		//分类标识
	CategoryStatus int    `json:"categoryStatus"`  		//分类状态
	CategoryLevel  int 	  `json:"categoryLevel"`		//分类层级
	ChildCategories  []string `json:"childCategories"`	//子分类
	ParentCategories []string `json:"parentCategories"`	//父分类
	CreateTime int64 `json:"categoryCreateTime"`  //分类创建时间
	UpdateTime int64 `json:"categoryUpdateTime"`  //分类更新时间
	DeleteTime int64 `json:"categoryDeleteTime"`  //分类删除时间
	CategoryName   language.Languages `json:"categoryName"`  //分类名称
}
