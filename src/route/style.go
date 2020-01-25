//author: richard
package route

import (
	"encoding/json"
	"github.com/advancevillage/3rd/https"
	"mms/src/config"
	"net/http"
)

//@Summary 创建款式
//@Produce json
//@Param language header string false "语言" default "chinese"
//@Param {} body route.RequestStyle true "CreateStyle"
//@Success 200 {object} route.HttpOk
//@Failure 400 {object} route.HttpError
//@Failure 404 {object} route.HttpError
//@Failure 500 {object} route.HttpError
//@Router /v1/styles [post]
func (s *service) CreateStyle(ctx *https.Context) {
	lang := s.language(ctx)
	body, err := s.body(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, s.NewHttpError(StyleCode, StyleMsg, RequestBodyErrorCode, RequestBodyErrorMsg))
		return
	}
	param := RequestStyle{}
	err = json.Unmarshal(body, &param)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, s.NewHttpError(StyleCode, StyleMsg, BodyStructureErrorCode, BodyStructureErrorMsg))
		return
	}
	param.Name.Multi(lang, config.Services().TranslateService(), config.Services().LogService())
	param.Description.Multi(lang, config.Services().TranslateService(), config.Services().LogService())
	err = config.Services().StyleService().CreateStyle(&param.Name, &param.Description)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, s.NewHttpError(StyleCode, StyleMsg, CreateErrorCode, CreateErrorMsg))
		return
	}
	ctx.JSON(http.StatusOK, s.NewHttpOk(http.StatusOK))
}

//@Summary 查询款式列表
//@Produce json
//@Param language header string false "语言" default "chinese"
//@Param page    query int false "页码" default "0"
//@Param perPage query int false "每页条数" default "20"
//@Param status  query int false "状态"
//@Success 200 {object} route.HttpOk
//@Failure 400 {object} route.HttpError
//@Failure 404 {object} route.HttpError
//@Failure 500 {object} route.HttpError
//@Router /v1/styles [get]
func (s *service) QueryStyles(ctx *https.Context) {
	page    := s.page(ctx)
	perPage := s.perPage(ctx)
	status  := s.status(ctx)
	styles, total, err := config.Services().StyleService().QueryStyles(status, page, perPage)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, s.NewHttpError(StyleCode, StyleMsg, QueryErrorCode, QueryErrorMsg))
		return
	}
	ctx.JSON(http.StatusOK, s.response(styles, total))
}

//@Summary 查询款式
//@Produce json
//@Param language header string false "语言" default "chinese"
//@Success 200 {object} route.HttpOk
//@Failure 400 {object} route.HttpError
//@Failure 404 {object} route.HttpError
//@Failure 500 {object} route.HttpError
//@Router /v1/styles/{pathId} [get]
func (s *service) QueryStyle(ctx *https.Context) {
	styleId, err := s.pathId(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, s.NewHttpError(StyleCode, StyleMsg, IDErrorCode, IDErrorMsg))
		return
	}
	style, err := config.Services().StyleService().QueryStyleById(styleId)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, s.NewHttpError(StyleCode, StyleMsg, QueryErrorCode, QueryErrorMsg))
		return
	}
	ctx.JSON(http.StatusOK, style)
}

//@Summary 更新款式
//@Produce json
//@Param language header string false "语言" default "chinese"
//@Param {} body route.RequestStyle true "UpdateStyle"
//@Success 200 {object} route.HttpOk
//@Failure 400 {object} route.HttpError
//@Failure 404 {object} route.HttpError
//@Failure 500 {object} route.HttpError
//@Router /v1/styles/{pathId} [put]
func (s *service) UpdateStyle(ctx *https.Context) {
	body, err := s.body(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, s.NewHttpError(StyleCode, StyleMsg, RequestBodyErrorCode, RequestBodyErrorMsg))
		return
	}
	styleId, err := s.pathId(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, s.NewHttpError(StyleCode, StyleMsg, IDErrorCode, IDErrorMsg))
		return
	}
	param := RequestStyle{}
	err = json.Unmarshal(body, &param)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, s.NewHttpError(StyleCode, StyleMsg, BodyStructureErrorCode, BodyStructureErrorMsg))
		return
	}
	err = config.Services().StyleService().UpdateStyle(styleId, &param.Name, &param.Description, param.Status)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, s.NewHttpError(StyleCode, StyleMsg, UpdateErrorCode, UpdateErrorMsg))
		return
	}
	ctx.JSON(http.StatusOK, s.NewHttpOk(http.StatusOK))
}

//@Summary 删除款式
//@Produce json
//@Param language header string false "语言" default "chinese"
//@Success 200 {object} route.HttpOk
//@Failure 400 {object} route.HttpError
//@Failure 404 {object} route.HttpError
//@Failure 500 {object} route.HttpError
//@Router /v1/styles/{pathId} [delete]
func (s *service) DeleteStyle(ctx *https.Context) {
	styleId, err := s.pathId(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, s.NewHttpError(StyleCode, StyleMsg, IDErrorCode, IDErrorMsg))
		return
	}
	err = config.Services().StyleService().DeleteStyle(styleId)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, s.NewHttpError(StyleCode, StyleMsg, DeleteErrorCode, DeleteErrorMsg))
		return
	}
	ctx.JSON(http.StatusOK, s.NewHttpOk(http.StatusOK))
}