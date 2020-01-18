//author: richard
//@link: https://my.oschina.net/u/2410867/blog/1647801
package category

import "mms/src/language"

const (
	Schema = "categories"
	SnowFlakeIdLength = 18

	StatusActive  = 0x201
	StatusDeleted = 0x202
	StatusInvalid = 0x299
)

type ICategory interface {
	CreateCategory(category *Category) error
	UpdateCategory(category *Category) error
	QueryCategory(categoryId string) (*Category, error)
	QueryCategories(where map[string]interface{}, page int, perPage int, sort map[string]interface{}) ([]Category, int64, error)
}

type Category struct {
	Id 	   string  `json:"id"`	  		//分类标识
	Status int     `json:"status"`  	//分类状态
	Level  int 	   `json:"level"`		//分类层级
	Child  []string  `json:"child"`		//子分类
	Parent []string  `json:"parent"`	//父分类
	CreateTime int64 `json:"createTime"`  	 //分类创建时间
	UpdateTime int64 `json:"updateTime"`  	 //分类更新时间
	DeleteTime int64 `json:"deleteTime"`  	 //分类删除时间
	Name   *language.Languages `json:"name"` //分类名称
}
