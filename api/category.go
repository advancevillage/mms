//author: richard
package api

type Category struct {
	Id 	   string    `json:"id,omitempty"`	  		//分类标识
	Level  int 	     `json:"level,omitempty"`		//分类层级
	Child  []string  `json:"child,omitempty"`		//子分类
	Parent []string  `json:"parent,omitempty"`		//父分类
	CreateTime int64 `json:"createTime,omitempty"`  //分类创建时间
	UpdateTime int64 `json:"updateTime,omitempty"`  //分类更新时间
	DeleteTime int64 `json:"deleteTime,omitempty"`  //分类删除时间
	Name  *Languages `json:"name,omitempty"` 		//分类名称
}
