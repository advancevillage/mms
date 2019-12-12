//author: richard
package route

import (
	"encoding/json"
	"github.com/advancevillage/3rd/https"
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
	version.Info = config.GetMMSObject().GetVersion()
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
		config.GetMMSObject().GetLogger().Warning(err.Error())
		ctx.JsonResponse(http.StatusBadRequest, &HttpError{Code: CategoryCreateErrorCode, Message: CategoryCreateErrorMsg})
		return
	}
	rc := RequestCategory{}
	err = json.Unmarshal(buf, &rc)
	if err != nil {
		config.GetMMSObject().GetLogger().Warning(err.Error())
		ctx.JsonResponse(http.StatusBadRequest, &HttpError{Code: CategoryCreateErrorCode, Message: CategoryCreateErrorMsg})
		return
	}
	err = config.GetMMSObject().GetCategoryService().CreateCategory(rc.CategoryName, rc.CategoryStatus, rc.ChildCategories, rc.ParentCategories)
	if err != nil {
		config.GetMMSObject().GetLogger().Warning(err.Error())
		ctx.JsonResponse(http.StatusBadRequest, &HttpError{Code: CategoryCreateErrorCode, Message: CategoryCreateErrorMsg})
		return
	}
	ctx.JsonResponse(http.StatusOK, &HttpOk{Code: http.StatusOK, Message: CategoryCreateOkMsg})
}