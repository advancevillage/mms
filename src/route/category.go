//author: richard
package route

import (
	"encoding/json"
	"github.com/advancevillage/3rd/https"
	"mms/src/config"
	"net/http"
)

//@Summary 新增分类
//@Produce json
//@Param {} body route.RequestCategory true "CreateCategory"
//@Success 200 {object} route.HttpOk
//@Failure 400 {object} route.HttpError
//@Failure 404 {object} route.HttpError
//@Failure 500 {object} route.HttpError
//@Router /v1/categories [post]
func (s *service) CreateCategory(ctx *https.Context) {
	body, err := s.body(ctx)
	if err != nil {
		ctx.JsonResponse(http.StatusBadRequest, s.NewHttpError(CategoryCode, CategoryMsg, RequestBodyErrorCode, RequestBodyErrorMsg))
		return
	}
	param := RequestCategory{}
	err = json.Unmarshal(body, &param)
	if err != nil {
		ctx.JsonResponse(http.StatusBadRequest, s.NewHttpError(CategoryCode, CategoryMsg, BodyStructureErrorCode, BodyStructureErrorMsg))
		return
	}
	err = config.Services().CategoryService().CreateCategory(param.NameEn, param.Level, param.Child, param.Parent )
	if err != nil {
		ctx.JsonResponse(http.StatusBadRequest, s.NewHttpError(CategoryCode, CategoryMsg, CreateErrorCode, CreateErrorMsg))
		return
	}
	ctx.JsonResponse(http.StatusOK, s.NewHttpOk(http.StatusOK))
}

//@Summary 查询分类列表
//@Produce json
//@Param page    query int false "页码" default "0"
//@Param perPage query int false "每页条数" default "20"
//@Param status  query int false "状态"
//@Success 200 {object} route.HttpOk
//@Failure 400 {object} route.HttpError
//@Failure 404 {object} route.HttpError
//@Failure 500 {object} route.HttpError
//@Router /v1/categories [get]
func (s *service) QueryCategories(ctx *https.Context) {
	page    := s.page(ctx)
	perPage := s.perPage(ctx)
	status  := s.status(ctx)
	categories, err := config.Services().CategoryService().QueryCategories(status, page, perPage)
	if err != nil {
		ctx.JsonResponse(http.StatusBadRequest, s.NewHttpError(CategoryCode, CategoryMsg, QueryErrorCode, QueryErrorMsg))
		return
	}
	ctx.JsonResponse(http.StatusOK, categories)
}

//@Summary 查询分类
//@Produce json
//@Success 200 {object} route.HttpOk
//@Failure 400 {object} route.HttpError
//@Failure 404 {object} route.HttpError
//@Failure 500 {object} route.HttpError
//@Router /v1/categories/{pathId} [get]
func (s *service) QueryCategory(ctx *https.Context) {
	categoryId, err := s.pathId(ctx)
	if err != nil {
		ctx.JsonResponse(http.StatusBadRequest, s.NewHttpError(CategoryCode, CategoryMsg, IDErrorCode, IDErrorMsg))
		return
	}
	category, err := config.Services().CategoryService().QueryCategoryById(categoryId)
	if err != nil {
		ctx.JsonResponse(http.StatusBadRequest, s.NewHttpError(CategoryCode, CategoryMsg, QueryErrorCode, QueryErrorMsg))
		return
	}
	ctx.JsonResponse(http.StatusOK, category)
}

//@Summary 更新分类
//@Produce json
//@Param {} body route.RequestCategory true "UpdateCategory"
//@Success 200 {object} route.HttpOk
//@Failure 400 {object} route.HttpError
//@Failure 404 {object} route.HttpError
//@Failure 500 {object} route.HttpError
//@Router /v1/categories/{pathId} [put]
func (s *service) UpdateCategory(ctx *https.Context) {
	body, err := s.body(ctx)
	if err != nil {
		ctx.JsonResponse(http.StatusBadRequest, s.NewHttpError(CategoryCode, CategoryMsg, RequestBodyErrorCode, RequestBodyErrorMsg))
		return
	}
	categoryId, err := s.pathId(ctx)
	if err != nil {
		ctx.JsonResponse(http.StatusBadRequest, s.NewHttpError(CategoryCode, CategoryMsg, IDErrorCode, IDErrorMsg))
		return
	}
	param := RequestCategory{}
	err = json.Unmarshal(body, &param)
	if err != nil {
		ctx.JsonResponse(http.StatusBadRequest, s.NewHttpError(CategoryCode, CategoryMsg, BodyStructureErrorCode, BodyStructureErrorMsg))
		return
	}
	err = config.Services().CategoryService().UpdateCategory(categoryId, param.NameEn, param.NameCn, param.Child, param.Parent, param.Status, param.Level)
	if err != nil {
		ctx.JsonResponse(http.StatusBadRequest, s.NewHttpError(CategoryCode, CategoryMsg, UpdateErrorCode, UpdateErrorMsg))
		return
	}
	ctx.JsonResponse(http.StatusOK, s.NewHttpOk(http.StatusOK))
}

//@Summary 删除分类
//@Produce json
//@Success 200 {object} route.HttpOk
//@Failure 400 {object} route.HttpError
//@Failure 404 {object} route.HttpError
//@Failure 500 {object} route.HttpError
//@Router /v1/categories/{pathId} [delete]
func (s *service) DeleteCategory(ctx *https.Context) {
	categoryId, err := s.pathId(ctx)
	if err != nil {
		ctx.JsonResponse(http.StatusBadRequest, s.NewHttpError(CategoryCode, CategoryMsg, IDErrorCode, IDErrorMsg))
		return
	}
	err = config.Services().CategoryService().DeleteCategory(categoryId)
	if err != nil {
		ctx.JsonResponse(http.StatusBadRequest, s.NewHttpError(CategoryCode, CategoryMsg, DeleteErrorCode, DeleteErrorMsg))
		return
	}
	ctx.JsonResponse(http.StatusOK, s.NewHttpOk(http.StatusOK))
}


