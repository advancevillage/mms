//author: richard
package route

import "github.com/advancevillage/3rd/https"

//@Summary 创建分类
//@Produce json
//@Param {} body route.RequestCategory true "CreateCategory"
//@Success 200 {object} route.HttpOk
//@Failure 400 {object} route.HttpError
//@Failure 404 {object} route.HttpError
//@Failure 500 {object} route.HttpError
//@Router /v1/categories [post]
func (s *service) CreateCategory(ctx *https.Context) {

}


