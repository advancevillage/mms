//author: richard
package category

//@note:
//@对象单一责任原则: 只需要导入repo && github.com/advancevillage/3rd/xxx
import (
	"github.com/advancevillage/3rd/storages"
)

type Service struct {
	repo ICategory
}

//TODO 创建分类如何传递配置数据
//eg: 数据库的存储资源,全局变量
func NewCategoryService(storage storages.Storage) *Service {
	return &Service{repo:NewCategoryRepoEs7(storage)}
}
