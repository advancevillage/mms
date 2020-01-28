//author: richard
package route

import (
	"encoding/json"
	"github.com/advancevillage/3rd/https"
	"mms/api"
	"net/http"
)

//@Summary 创建尺码
//@Produce json
//@Param x-language header string false "语言" default "chinese"
//@Param {} body api.Size true "CreateSize"
//@Success 200 {object} route.httpOk
//@Failure 400 {object} route.httpError
//@Failure 404 {object} route.httpError
//@Failure 500 {object} route.httpError
//@Router /v1/sizes [post]
func (s *Service) CreateSize(ctx *https.Context) {
	body, err := s.body(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, s.newHttpError(SizeCode, SizeMsg, BodyErrorCode, BodyErrorMsg))
		return
	}
	param := api.Size{}
	err = json.Unmarshal(body, &param)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, s.newHttpError(SizeCode, SizeMsg, ContextErrorCode, ContextErrorMsg))
		return
	}

	err = s.sizeService.CreateSize(&param)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, s.newHttpError(SizeCode, SizeMsg, CreateErrorCode, CreateErrorMsg))
		return
	}
	ctx.JSON(http.StatusOK, s.newHttpOk())
}

//@Summary 查询尺码列表
//@Produce json
//@Param x-language header string false "语言" default "chinese"
//@Param page    query int false "页码" default "0"
//@Param perPage query int false "每页条数" default "20"
//@Param status  query int false "状态"
//@Param group   query string false "组" default "number"
//@Success 200 {object} route.httpOk
//@Failure 400 {object} route.httpError
//@Failure 404 {object} route.httpError
//@Failure 500 {object} route.httpError
//@Router /v1/sizes [get]
func (s *Service) QuerySizes(ctx *https.Context) {
	lang    := s.language(ctx)
	page    := s.page(ctx)
	perPage := s.perPage(ctx)
	group   := s.group(ctx)
	sizes, total,  err := s.sizeService.QuerySizes(page, perPage, group, lang)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, s.newHttpError(SizeCode, SizeMsg, QueryErrorCode, QueryErrorMsg))
		return
	}
	ctx.JSON(http.StatusOK, s.response(sizes, total))
}

//@Summary 查询尺码
//@Produce json
//@Param x-language header string false "语言" default "chinese"
//@Success 200 {object} route.httpOk
//@Failure 400 {object} route.httpError
//@Failure 404 {object} route.httpError
//@Failure 500 {object} route.httpError
//@Router /v1/sizes/{pathId} [get]
func (s *Service) QuerySize(ctx *https.Context) {
	sizeId, err := s.pathId(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, s.newHttpError(SizeCode, SizeMsg, IDErrorCode, IDErrorMsg))
		return
	}
	size, err := s.sizeService.QuerySizeById(sizeId)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, s.newHttpError(SizeCode, SizeMsg, QueryErrorCode, QueryErrorMsg))
		return
	}
	ctx.JSON(http.StatusOK, size)
}

//@Summary 更新品牌
//@Produce json
//@Param x-language header string false "语言" default "chinese"
//@Param {} body api.Size true "UpdateSize"
//@Success 200 {object} route.httpOk
//@Failure 400 {object} route.httpError
//@Failure 404 {object} route.httpError
//@Failure 500 {object} route.httpError
//@Router /v1/sizes/{pathId} [put]
func (s *Service) UpdateSize(ctx *https.Context) {
	body, err := s.body(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, s.newHttpError(SizeCode, SizeMsg, BodyErrorCode, BodyErrorMsg))
		return
	}
	sizeId, err := s.pathId(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, s.newHttpError(SizeCode, SizeMsg, IDErrorCode, IDErrorMsg))
		return
	}
	param := api.Size{}
	err = json.Unmarshal(body, &param)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, s.newHttpError(SizeCode, SizeMsg, ContextErrorCode, ContextErrorMsg))
		return
	}
	param.Id = sizeId
	err = s.sizeService.UpdateSize(&param)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, s.newHttpError(SizeCode, SizeMsg, UpdateErrorCode, UpdateErrorMsg))
		return
	}
	ctx.JSON(http.StatusOK, s.newHttpOk())
}

//@Summary 删除尺码
//@Produce json
//@Param x-language header string false "语言" default "chinese"
//@Success 200 {object} route.httpOk
//@Failure 400 {object} route.httpError
//@Failure 404 {object} route.httpError
//@Failure 500 {object} route.httpError
//@Router /v1/sizes/{pathId} [delete]
func (s *Service) DeleteSize(ctx *https.Context) {
	sizeId, err := s.pathId(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, s.newHttpError(SizeCode, SizeMsg, IDErrorCode, IDErrorMsg))
		return
	}
	err = s.sizeService.DeleteSize(sizeId)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, s.newHttpError(SizeCode, SizeMsg, DeleteErrorCode, DeleteErrorMsg))
		return
	}
	ctx.JSON(http.StatusOK, s.newHttpOk())
}
