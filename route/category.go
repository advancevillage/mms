//author: richard
package route

import (
	"encoding/json"
	"github.com/advancevillage/3rd/https"
	"mms/api"
	"net/http"
)

//@Summary 新增分类
//@Produce json
//@Param x-language header string false "语言" default "chinese"
//@Param {} body api.Category true "CreateCategory"
//@Success 200 {object} route.HttpOk
//@Failure 400 {object} route.HttpError
//@Failure 404 {object} route.HttpError
//@Failure 500 {object} route.HttpError
//@Router /v1/categories [post]
func (s *Service) CreateCategory(ctx *https.Context) {
	lang := s.language(ctx)
	body, err := s.body(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, s.newHttpError(CategoryCode, CategoryMsg, BodyErrorCode, BodyErrorMsg))
		return
	}
	param := api.Category{}
	err = json.Unmarshal(body, &param)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, s.newHttpError(CategoryCode, CategoryMsg, ContextErrorCode, ContextErrorMsg))
		return
	}

	err = s.langService.I18n(param.Name, lang)

	err = s.categoryService.CreateCategory(&param)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, s.newHttpError(CategoryCode, CategoryMsg, CreateErrorCode, CreateErrorMsg))
		return
	}
	ctx.JSON(http.StatusOK, s.newHttpOk())
}

//@Summary 查询分类列表
//@Produce json
//@Param x-language header string false "语言" default "chinese"
//@Param page    query int false "页码" default "0"
//@Param perPage query int false "每页条数" default "20"
//@Param status  query int false "状态"
//@Param level   query int false "层级" default "1"
//@Success 200 {object} route.HttpOk
//@Failure 400 {object} route.HttpError
//@Failure 404 {object} route.HttpError
//@Failure 500 {object} route.HttpError
//@Router /v1/categories [get]
func (s *Service) QueryCategories(ctx *https.Context) {
	page    := s.page(ctx)
	perPage := s.perPage(ctx)
	level   := s.level(ctx)
	categories, total, err := s.categoryService.QueryCategories(page, perPage, level)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, s.newHttpError(CategoryCode, CategoryMsg, QueryErrorCode, QueryErrorMsg))
		return
	}
	ctx.JSON(http.StatusOK, s.response(categories, total))
}

//@Summary 查询子分类列表
//@Produce json
//@Param x-language header string false "语言" default "chinese"
//@Param page    query int false "页码" default "0"
//@Param perPage query int false "每页条数" default "20"
//@Param status  query int false "状态"
//@Param level   query int false "层级" default "1"
//@Success 200 {object} route.HttpOk
//@Failure 400 {object} route.HttpError
//@Failure 404 {object} route.HttpError
//@Failure 500 {object} route.HttpError
//@Router /v1/categories/{pathId}/categories [get]
func (s *Service) QueryChildCategories(ctx *https.Context) {
	categoryId, err := s.pathId(ctx)
	categories, err := s.categoryService.QueryChildCategories(categoryId)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, s.newHttpError(CategoryCode, CategoryMsg, QueryErrorCode, QueryErrorMsg))
		return
	}
	ctx.JSON(http.StatusOK, categories)
}

//@Summary 查询分类
//@Produce json
//@Param x-language header string false "语言" default "chinese"
//@Success 200 {object} route.HttpOk
//@Failure 400 {object} route.HttpError
//@Failure 404 {object} route.HttpError
//@Failure 500 {object} route.HttpError
//@Router /v1/categories/{pathId} [get]
func (s *Service) QueryCategory(ctx *https.Context) {
	categoryId, err := s.pathId(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, s.newHttpError(CategoryCode, CategoryMsg, IDErrorCode, IDErrorMsg))
		return
	}
	category, err := s.categoryService.QueryCategoryById(categoryId)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, s.newHttpError(CategoryCode, CategoryMsg, QueryErrorCode, QueryErrorMsg))
		return
	}
	ctx.JSON(http.StatusOK, category)
}

//@Summary 更新分类
//@Produce json
//@Param x-language header string false "语言" default "chinese"
//@Param {} body route.RequestCategory true "UpdateCategory"
//@Success 200 {object} route.HttpOk
//@Failure 400 {object} route.HttpError
//@Failure 404 {object} route.HttpError
//@Failure 500 {object} route.HttpError
//@Router /v1/categories/{pathId} [put]
func (s *Service) UpdateCategory(ctx *https.Context) {
	body, err := s.body(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, s.newHttpError(CategoryCode, CategoryMsg, BodyErrorCode, BodyErrorMsg))
		return
	}
	categoryId, err := s.pathId(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, s.newHttpError(CategoryCode, CategoryMsg, IDErrorCode, IDErrorMsg))
		return
	}
	param := api.Category{}
	err = json.Unmarshal(body, &param)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, s.newHttpError(CategoryCode, CategoryMsg, ContextErrorCode, ContextErrorMsg))
		return
	}
	param.Id = categoryId
	err = s.categoryService.UpdateCategory(&param)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, s.newHttpError(CategoryCode, CategoryMsg, UpdateErrorCode, UpdateErrorMsg))
		return
	}
	ctx.JSON(http.StatusOK, s.newHttpOk())
}

//@Summary 删除分类
//@Produce json
//@Param x-language header string false "语言" default "chinese"
//@Success 200 {object} route.HttpOk
//@Failure 400 {object} route.HttpError
//@Failure 404 {object} route.HttpError
//@Failure 500 {object} route.HttpError
//@Router /v1/categories/{pathId} [delete]
func (s *Service) DeleteCategory(ctx *https.Context) {
	categoryId, err := s.pathId(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, s.newHttpError(CategoryCode, CategoryMsg, IDErrorCode, IDErrorMsg))
		return
	}
	err = s.categoryService.DeleteCategory(categoryId)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, s.newHttpError(CategoryCode, CategoryMsg, DeleteErrorCode, DeleteErrorMsg))
		return
	}
	ctx.JSON(http.StatusOK, s.newHttpOk())
}
