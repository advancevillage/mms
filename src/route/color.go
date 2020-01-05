//author: richard
package route

import (
	"encoding/json"
	"github.com/advancevillage/3rd/https"
	"mms/src/config"
	"net/http"
)

//@Summary 新增颜色
//@Produce json
//@Param {} body route.RequestColor true "CreateColor"
//@Success 200 {object} route.HttpOk
//@Failure 400 {object} route.HttpError
//@Failure 404 {object} route.HttpError
//@Failure 500 {object} route.HttpError
//@Router /v1/colors [post]
func (s *service) CreateColor(ctx *https.Context) {
	body, err := s.body(ctx)
	if err != nil {
		ctx.JsonResponse(http.StatusBadRequest, s.NewHttpError(ColorCode, ColorMsg, RequestBodyErrorCode, RequestBodyErrorMsg))
		return
	}
	param := RequestColor{}
	err = json.Unmarshal(body, &param)
	if err != nil {
		ctx.JsonResponse(http.StatusBadRequest, s.NewHttpError(ColorCode, ColorMsg, BodyStructureErrorCode, BodyStructureErrorMsg))
		return
	}
	err = config.Services().ColorService().CreateColor(param.NameEn, param.Rgba)
	if err != nil {
		ctx.JsonResponse(http.StatusBadRequest, s.NewHttpError(ColorCode, ColorMsg, CreateErrorCode, CreateErrorMsg))
		return
	}
	ctx.JsonResponse(http.StatusOK, s.NewHttpOk(http.StatusOK))
}

//@Summary 查询颜色列表
//@Produce json
//@Param page    query int false "页码" default "0"
//@Param perPage query int false "每页条数" default "20"
//@Param status  query int false "状态"
//@Success 200 {object} route.HttpOk
//@Failure 400 {object} route.HttpError
//@Failure 404 {object} route.HttpError
//@Failure 500 {object} route.HttpError
//@Router /v1/colors [get]
func (s *service) QueryColors(ctx *https.Context) {
	page    := s.page(ctx)
	perPage := s.perPage(ctx)
	status  := s.status(ctx)
	colors, err := config.Services().ColorService().QueryColors(status, page, perPage)
	if err != nil {
		ctx.JsonResponse(http.StatusBadRequest, s.NewHttpError(ColorCode, ColorMsg, QueryErrorCode, QueryErrorMsg))
		return
	}
	ctx.JsonResponse(http.StatusOK, colors)
}

//@Summary 查询颜色
//@Produce json
//@Success 200 {object} route.HttpOk
//@Failure 400 {object} route.HttpError
//@Failure 404 {object} route.HttpError
//@Failure 500 {object} route.HttpError
//@Router /v1/colors/{pathId} [get]
func (s *service) QueryColor(ctx *https.Context) {
	colorId, err := s.pathId(ctx)
	if err != nil {
		ctx.JsonResponse(http.StatusBadRequest, s.NewHttpError(ColorCode, ColorMsg, IDErrorCode, IDErrorMsg))
		return
	}
	color, err := config.Services().ColorService().QueryColorById(colorId)
	if err != nil {
		ctx.JsonResponse(http.StatusBadRequest, s.NewHttpError(ColorCode, ColorMsg, QueryErrorCode, QueryErrorMsg))
		return
	}
	ctx.JsonResponse(http.StatusOK, color)
}

//@Summary 更新颜色
//@Produce json
//@Param {} body route.RequestColor true "UpdateColor"
//@Success 200 {object} route.HttpOk
//@Failure 400 {object} route.HttpError
//@Failure 404 {object} route.HttpError
//@Failure 500 {object} route.HttpError
//@Router /v1/colors/{pathId} [put]
func (s *service) UpdateColor(ctx *https.Context) {
	body, err := s.body(ctx)
	if err != nil {
		ctx.JsonResponse(http.StatusBadRequest, s.NewHttpError(ColorCode, ColorMsg, RequestBodyErrorCode, RequestBodyErrorMsg))
		return
	}
	colorId, err := s.pathId(ctx)
	if err != nil {
		ctx.JsonResponse(http.StatusBadRequest, s.NewHttpError(ColorCode, ColorMsg, IDErrorCode, IDErrorMsg))
		return
	}
	param := RequestColor{}
	err = json.Unmarshal(body, &param)
	if err != nil {
		ctx.JsonResponse(http.StatusBadRequest, s.NewHttpError(ColorCode, ColorMsg, BodyStructureErrorCode, BodyStructureErrorMsg))
		return
	}
	err = config.Services().ColorService().UpdateColor(colorId, param.NameEn, param.NameCn, param.Rgba, param.Status)
	if err != nil {
		ctx.JsonResponse(http.StatusBadRequest, s.NewHttpError(ColorCode, ColorMsg, UpdateErrorCode, UpdateErrorMsg))
		return
	}
	ctx.JsonResponse(http.StatusOK, s.NewHttpOk(http.StatusOK))
}

//@Summary 删除颜色
//@Produce json
//@Success 200 {object} route.HttpOk
//@Failure 400 {object} route.HttpError
//@Failure 404 {object} route.HttpError
//@Failure 500 {object} route.HttpError
//@Router /v1/colors/{pathId} [delete]
func (s *service) DeleteColor(ctx *https.Context) {
	colorId, err := s.pathId(ctx)
	if err != nil {
		ctx.JsonResponse(http.StatusBadRequest, s.NewHttpError(ColorCode, ColorMsg, IDErrorCode, IDErrorMsg))
		return
	}
	err = config.Services().ColorService().DeleteColor(colorId)
	if err != nil {
		ctx.JsonResponse(http.StatusBadRequest, s.NewHttpError(ColorCode, ColorMsg, DeleteErrorCode, DeleteErrorMsg))
		return
	}
	ctx.JsonResponse(http.StatusOK, s.NewHttpOk(http.StatusOK))
}
