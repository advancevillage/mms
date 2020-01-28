//author: richard
package route

import (
	"encoding/json"
	"github.com/advancevillage/3rd/https"
	"mms/api"
	"net/http"
)

//@Summary 新增颜色
//@Produce json
//@Param x-language header string false "语言" default "chinese"
//@Param {} body api.Color true "CreateColor"
//@Success 200 {object} route.httpOk
//@Failure 400 {object} route.httpError
//@Failure 404 {object} route.httpError
//@Failure 500 {object} route.httpError
//@Router /v1/colors [post]
func (s *Service) CreateColor(ctx *https.Context) {
	lang := s.language(ctx)
	body, err := s.body(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, s.newHttpError(ColorCode, ColorMsg, BodyErrorCode, BodyErrorMsg))
		return
	}
	param := api.Color{}
	err = json.Unmarshal(body, &param)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, s.newHttpError(ColorCode, ColorMsg, ContextErrorCode, ContextErrorMsg))
		return
	}

	err = s.langService.I18n(param.Name, lang)
	if err != nil {
		s.configService.Logger.Warning(err.Error())
	}

	err = s.colorService.CreateColor(&param)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, s.newHttpError(ColorCode, ColorMsg, CreateErrorCode, CreateErrorMsg))
		return
	}
	ctx.JSON(http.StatusOK, s.newHttpOk())
}

//@Summary 查询颜色列表
//@Produce json
//@Param x-language header string false "语言" default "chinese"
//@Param page    query int false "页码" default "0"
//@Param perPage query int false "每页条数" default "20"
//@Success 200 {object} route.httpOk
//@Failure 400 {object} route.httpError
//@Failure 404 {object} route.httpError
//@Failure 500 {object} route.httpError
//@Router /v1/colors [get]
func (s *Service) QueryColors(ctx *https.Context) {
	page    := s.page(ctx)
	perPage := s.perPage(ctx)
	colors, total, err := s.colorService.QueryColors(page, perPage)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, s.newHttpError(ColorCode, ColorMsg, QueryErrorCode, QueryErrorMsg))
		return
	}
	ctx.JSON(http.StatusOK, s.response(colors, total))
}

//@Summary 查询颜色
//@Produce json
//@Param x-language header string false "语言" default "chinese"
//@Success 200 {object} route.httpOk
//@Failure 400 {object} route.httpError
//@Failure 404 {object} route.httpError
//@Failure 500 {object} route.httpError
//@Router /v1/colors/{pathId} [get]
func (s *Service) QueryColor(ctx *https.Context) {
	colorId, err := s.pathId(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, s.newHttpError(ColorCode, ColorMsg, IDErrorCode, IDErrorMsg))
		return
	}
	color, err := s.colorService.QueryColorById(colorId)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, s.newHttpError(ColorCode, ColorMsg, QueryErrorCode, QueryErrorMsg))
		return
	}
	ctx.JSON(http.StatusOK, color)
}

//@Summary 更新颜色
//@Produce json
//@Param x-language header string false "语言" default "chinese"
//@Param {} body api.Color true "UpdateColor"
//@Success 200 {object} route.httpOk
//@Failure 400 {object} route.httpError
//@Failure 404 {object} route.httpError
//@Failure 500 {object} route.httpError
//@Router /v1/colors/{pathId} [put]
func (s *Service) UpdateColor(ctx *https.Context) {
	body, err := s.body(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, s.newHttpError(ColorCode, ColorMsg, BodyErrorCode, BodyErrorMsg))
		return
	}
	colorId, err := s.pathId(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, s.newHttpError(ColorCode, ColorMsg, IDErrorCode, IDErrorMsg))
		return
	}
	param := api.Color{}
	err = json.Unmarshal(body, &param)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, s.newHttpError(ColorCode, ColorMsg, ContextErrorCode, ContextErrorMsg))
		return
	}
	param.Id = colorId
	err = s.colorService.UpdateColor(&param)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, s.newHttpError(ColorCode, ColorMsg, UpdateErrorCode, UpdateErrorMsg))
		return
	}
	ctx.JSON(http.StatusOK, s.newHttpOk())
}

//@Summary 删除颜色
//@Produce json
//@Param x-language header string false "语言" default "chinese"
//@Success 200 {object} route.httpOk
//@Failure 400 {object} route.httpError
//@Failure 404 {object} route.httpError
//@Failure 500 {object} route.httpError
//@Router /v1/colors/{pathId} [delete]
func (s *Service) DeleteColor(ctx *https.Context) {
	colorId, err := s.pathId(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, s.newHttpError(ColorCode, ColorMsg, IDErrorCode, IDErrorMsg))
		return
	}
	err = s.colorService.DeleteColor(colorId)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, s.newHttpError(ColorCode, ColorMsg, DeleteErrorCode, DeleteErrorMsg))
		return
	}
	ctx.JSON(http.StatusOK, s.newHttpOk())
}

