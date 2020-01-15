//author: richard
//@link: https://my.oschina.net/u/2410867/blog/1647801
package category

import "mms/src/component/language"

const (
	Schema = "categories"

	StatusActive  = 0x201
	StatusDeleted = 0x202
	StatusInvalid = 0x299
)

type ICategory interface {
	CreateCategory(category *Category) error
	UpdateCategory(category *Category) error
	QueryCategory(categoryId string) (*Category, error)
	QueryCategories(where map[string]interface{}, page int, perPage int) ([]Category, int64, error)
}

type Category struct {
	Id 	   string  `json:"categoryId"`	  		//分类标识
	Status int     `json:"categoryStatus"`  	//分类状态
	Level  int 	   `json:"categoryLevel"`		//分类层级
	Child  []string  `json:"childCategories"`	//子分类
	Parent []string  `json:"parentCategories"`	//父分类
	CreateTime int64 `json:"categoryCreateTime"`  //分类创建时间
	UpdateTime int64 `json:"categoryUpdateTime"`  //分类更新时间
	DeleteTime int64 `json:"categoryDeleteTime"`  //分类删除时间
	Name   language.Languages `json:"categoryName"`  //分类名称
}
