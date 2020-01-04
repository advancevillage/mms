//author: richard
package route

import (
	"encoding/json"
	"github.com/advancevillage/3rd/https"
	"mms/src/component/category"
	"mms/src/config"
	"net/http"
)

type service struct {}

func NewApiService() *service {
	return &service{}
}

//@Summary 显示当前服务的版本和代码版本号
//@Produce json
//@Success 200 {object} route.ResponseVersion
//@Failure 400 {object} route.HttpError
//@Failure 404 {object} route.HttpError
//@Failure 500 {object} route.HttpError
//@Router /v1/merchandises/version [get]
func (s *service) Version(ctx *https.Context) {
	version := ResponseVersion{}
	version.Info = config.Services().Version()
	ctx.JsonResponse(http.StatusOK, version)
}

//@Summary 创建分类
//@Produce json
//@Param {} body route.RequestCategory true "CreateCategory"
//@Success 200 {object} route.HttpOk
//@Failure 400 {object} route.HttpError
//@Failure 404 {object} route.HttpError
//@Failure 500 {object} route.HttpError
//@Router /v1/categories [post]
func (s *service) CreateCategory(ctx *https.Context) {
	buf, err := ctx.Body()
	if err != nil {
		config.Services().LogService().Warning(err.Error())
		ctx.JsonResponse(http.StatusBadRequest, &HttpError{Code: CategoryCreateErrorCode, Message: CategoryCreateErrorMsg})
		return
	}
	rc := RequestCategory{}
	err = json.Unmarshal(buf, &rc)
	if err != nil {
		config.Services().LogService().Warning(err.Error())
		ctx.JsonResponse(http.StatusBadRequest, &HttpError{Code: CategoryCreateErrorCode, Message: CategoryCreateErrorMsg})
		return
	}
	cat := category.Category{}
	cat.CategoryStatus = rc.CategoryStatus
	cat.ChildCategories = rc.ChildCategories
	cat.ParentCategories = rc.ParentCategories
	err = config.Services().CategoryService().CreateCategory(&cat)
	if err != nil {
		config.Services().LogService().Warning(err.Error())
		ctx.JsonResponse(http.StatusBadRequest, &HttpError{Code: CategoryCreateErrorCode, Message: CategoryCreateErrorMsg})
		return
	}
	ctx.JsonResponse(http.StatusOK, &HttpOk{Code: http.StatusOK, Message: CategoryCreateOkMsg})
}

//@Summary 查询商品列表
//@Produce json
//@Param {} body route.RequestCategory true "QueryMerchandises"
//@Success 200 {object} route.HttpOk
//@Failure 400 {object} route.HttpError
//@Failure 404 {object} route.HttpError
//@Failure 500 {object} route.HttpError
//@Router /v1/merchandises [get]
func (s *service) QueryMerchandises(ctx *https.Context) {

}

//@Summary 创建商品
//@Produce json
//@Param {} body route.RequestCategory true "CreateMerchandise"
//@Success 200 {object} route.HttpOk
//@Failure 400 {object} route.HttpError
//@Failure 404 {object} route.HttpError
//@Failure 500 {object} route.HttpError
//@Router /v1/merchandises [post]
func (s *service) CreateMerchandise(ctx *https.Context) {

}

//@Summary 更新商品
//@Produce json
//@Param {} body route.RequestCategory true "UpdateMerchandise"
//@Success 200 {object} route.HttpOk
//@Failure 400 {object} route.HttpError
//@Failure 404 {object} route.HttpError
//@Failure 500 {object} route.HttpError
//@Router /v1/merchandises/{goodsId} [put]
func (s *service) UpdateMerchandise(ctx *https.Context) {

}

//@Summary 查询商品
//@Produce json
//@Param {} body route.RequestCategory true "QueryMerchandise"
//@Success 200 {object} route.HttpOk
//@Failure 400 {object} route.HttpError
//@Failure 404 {object} route.HttpError
//@Failure 500 {object} route.HttpError
//@Router /v1/merchandises/{goodsId} [get]
func (s *service) QueryMerchandise(ctx *https.Context) {

}

//@Summary 删除商品
//@Produce json
//@Param {} body route.RequestCategory true "DeleteMerchandise"
//@Success 200 {object} route.HttpOk
//@Failure 400 {object} route.HttpError
//@Failure 404 {object} route.HttpError
//@Failure 500 {object} route.HttpError
//@Router /v1/merchandises/{goodsId} [delete]
func (s *service) DeleteMerchandise(ctx *https.Context) {

}