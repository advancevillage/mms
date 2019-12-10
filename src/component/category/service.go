//author: richard
package category

type Service struct {

}

//TODO 创建分类如何传递配置数据
//eg: 数据库的存储资源,全局变量
func NewCategoryService() (*Service, error) {
	s := &Service{}
	return s, nil
}
