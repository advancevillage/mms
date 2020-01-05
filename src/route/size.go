//author: richard
package route

import (
	"encoding/json"
	"github.com/advancevillage/3rd/https"
	"mms/src/config"
	"net/http"
)

//@Summary 创建尺码
//@Produce json
//@Param {} body route.RequestSize true "CreateSize"
//@Success 200 {object} route.HttpOk
//@Failure 400 {object} route.HttpError
//@Failure 404 {object} route.HttpError
//@Failure 500 {object} route.HttpError
//@Router /v1/sizes [post]
func (s *service) CreateSize(ctx *https.Context) {
	body, err := s.body(ctx)
	if err != nil {
		ctx.JsonResponse(http.StatusBadRequest, s.NewHttpError(SizeCode, SizeMsg, RequestBodyErrorCode, RequestBodyErrorMsg))
		return
	}
	param := RequestSize{}
	err = json.Unmarshal(body, &param)
	if err != nil {
		ctx.JsonResponse(http.StatusBadRequest, s.NewHttpError(SizeCode, SizeMsg, BodyStructureErrorCode, BodyStructureErrorMsg))
		return
	}
	err = config.Services().SizeService().CreateSize(param.NameEn)
	if err != nil {
		ctx.JsonResponse(http.StatusBadRequest, s.NewHttpError(SizeCode, SizeMsg, CreateErrorCode, CreateErrorMsg))
		return
	}
	ctx.JsonResponse(http.StatusOK, s.NewHttpOk(http.StatusOK))
}

//@Summary 查询尺码列表
//@Produce json
//@Param page    query int false "页码" default "0"
//@Param perPage query int false "每页条数" default "20"
//@Param status  query int false "状态"
//@Success 200 {object} route.HttpOk
//@Failure 400 {object} route.HttpError
//@Failure 404 {object} route.HttpError
//@Failure 500 {object} route.HttpError
//@Router /v1/sizes [get]
func (s *service) QuerySizes(ctx *https.Context) {
	page    := s.page(ctx)
	perPage := s.perPage(ctx)
	status  := s.status(ctx)
	sizes, err := config.Services().SizeService().QuerySizes(status, page, perPage)
	if err != nil {
		ctx.JsonResponse(http.StatusBadRequest, s.NewHttpError(SizeCode, SizeMsg, QueryErrorCode, QueryErrorMsg))
		return
	}
	ctx.JsonResponse(http.StatusOK, sizes)
}

//@Summary 查询尺码
//@Produce json
//@Success 200 {object} route.HttpOk
//@Failure 400 {object} route.HttpError
//@Failure 404 {object} route.HttpError
//@Failure 500 {object} route.HttpError
//@Router /v1/sizes/{pathId} [get]
func (s *service) QuerySize(ctx *https.Context) {
	sizeId, err := s.pathId(ctx)
	if err != nil {
		ctx.JsonResponse(http.StatusBadRequest, s.NewHttpError(SizeCode, SizeMsg, IDErrorCode, IDErrorMsg))
		return
	}
	size, err := config.Services().SizeService().QuerySizeById(sizeId)
	if err != nil {
		ctx.JsonResponse(http.StatusBadRequest, s.NewHttpError(SizeCode, SizeMsg, QueryErrorCode, QueryErrorMsg))
		return
	}
	ctx.JsonResponse(http.StatusOK, size)
}

//@Summary 更新品牌
//@Produce json
//@Param {} body route.RequestSize true "UpdateSize"
//@Success 200 {object} route.HttpOk
//@Failure 400 {object} route.HttpError
//@Failure 404 {object} route.HttpError
//@Failure 500 {object} route.HttpError
//@Router /v1/sizes/{pathId} [put]
func (s *service) UpdateSize(ctx *https.Context) {
	body, err := s.body(ctx)
	if err != nil {
		ctx.JsonResponse(http.StatusBadRequest, s.NewHttpError(SizeCode, SizeMsg, RequestBodyErrorCode, RequestBodyErrorMsg))
		return
	}
	sizeId, err := s.pathId(ctx)
	if err != nil {
		ctx.JsonResponse(http.StatusBadRequest, s.NewHttpError(SizeCode, SizeMsg, IDErrorCode, IDErrorMsg))
		return
	}
	param := RequestSize{}
	err = json.Unmarshal(body, &param)
	if err != nil {
		ctx.JsonResponse(http.StatusBadRequest, s.NewHttpError(SizeCode, SizeMsg, BodyStructureErrorCode, BodyStructureErrorMsg))
		return
	}
	err = config.Services().SizeService().UpdateSize(sizeId, param.NameEn, param.NameCn, param.Status)
	if err != nil {
		ctx.JsonResponse(http.StatusBadRequest, s.NewHttpError(SizeCode, SizeMsg, UpdateErrorCode, UpdateErrorMsg))
		return
	}
	ctx.JsonResponse(http.StatusOK, s.NewHttpOk(http.StatusOK))
}

//@Summary 删除尺码
//@Produce json
//@Success 200 {object} route.HttpOk
//@Failure 400 {object} route.HttpError
//@Failure 404 {object} route.HttpError
//@Failure 500 {object} route.HttpError
//@Router /v1/sizes/{pathId} [delete]
func (s *service) DeleteSize(ctx *https.Context) {
	sizeId, err := s.pathId(ctx)
	if err != nil {
		ctx.JsonResponse(http.StatusBadRequest, s.NewHttpError(SizeCode, SizeMsg, IDErrorCode, IDErrorMsg))
		return
	}
	err = config.Services().SizeService().DeleteSize(sizeId)
	if err != nil {
		ctx.JsonResponse(http.StatusBadRequest, s.NewHttpError(SizeCode, SizeMsg, DeleteErrorCode, DeleteErrorMsg))
		return
	}
	ctx.JsonResponse(http.StatusOK, s.NewHttpOk(http.StatusOK))
}
