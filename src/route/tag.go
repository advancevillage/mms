//author: richard
package route

import (
	"encoding/json"
	"github.com/advancevillage/3rd/https"
	"mms/src/config"
	"net/http"
)

//@Summary 创建标签
//@Produce json
//@Param {} body route.RequestTag true "CreateTag"
//@Success 200 {object} route.HttpOk
//@Failure 400 {object} route.HttpError
//@Failure 404 {object} route.HttpError
//@Failure 500 {object} route.HttpError
//@Router /v1/tags [post]
func (s *service) CreateTag(ctx *https.Context) {
	body, err := s.body(ctx)
	if err != nil {
		ctx.JsonResponse(http.StatusBadRequest, s.NewHttpError(TagCode, TagMsg, RequestBodyErrorCode, RequestBodyErrorMsg))
		return
	}
	param := RequestTag{}
	err = json.Unmarshal(body, &param)
	if err != nil {
		ctx.JsonResponse(http.StatusBadRequest, s.NewHttpError(TagCode, TagMsg, BodyStructureErrorCode, BodyStructureErrorMsg))
		return
	}
	err = config.Services().TagService().CreateTag(param.NameEn)
	if err != nil {
		ctx.JsonResponse(http.StatusBadRequest, s.NewHttpError(TagCode, TagMsg, CreateErrorCode, CreateErrorMsg))
		return
	}
	ctx.JsonResponse(http.StatusOK, s.NewHttpOk(http.StatusOK))
}

//@Summary 查询标签列表
//@Produce json
//@Param page    query int false "页码" default "0"
//@Param perPage query int false "每页条数" default "20"
//@Param status  query int false "状态"
//@Success 200 {object} route.HttpOk
//@Failure 400 {object} route.HttpError
//@Failure 404 {object} route.HttpError
//@Failure 500 {object} route.HttpError
//@Router /v1/tags [get]
func (s *service) QueryTags(ctx *https.Context) {
	page    := s.page(ctx)
	perPage := s.perPage(ctx)
	status  := s.status(ctx)
	tags, err := config.Services().TagService().QueryTags(status, page, perPage)
	if err != nil {
		ctx.JsonResponse(http.StatusBadRequest, s.NewHttpError(TagCode, TagMsg, QueryErrorCode, QueryErrorMsg))
		return
	}
	ctx.JsonResponse(http.StatusOK, tags)
}

//@Summary 查询标签
//@Produce json
//@Success 200 {object} route.HttpOk
//@Failure 400 {object} route.HttpError
//@Failure 404 {object} route.HttpError
//@Failure 500 {object} route.HttpError
//@Router /v1/tags/{pathId} [get]
func (s *service) QueryTag(ctx *https.Context) {
	tagId, err := s.pathId(ctx)
	if err != nil {
		ctx.JsonResponse(http.StatusBadRequest, s.NewHttpError(TagCode, TagMsg, IDErrorCode, IDErrorMsg))
		return
	}
	tag, err := config.Services().TagService().QueryTagById(tagId)
	if err != nil {
		ctx.JsonResponse(http.StatusBadRequest, s.NewHttpError(TagCode, TagMsg, QueryErrorCode, QueryErrorMsg))
		return
	}
	ctx.JsonResponse(http.StatusOK, tag)
}

//@Summary 更新标签
//@Produce json
//@Param {} body route.RequestTag true "UpdateTag"
//@Success 200 {object} route.HttpOk
//@Failure 400 {object} route.HttpError
//@Failure 404 {object} route.HttpError
//@Failure 500 {object} route.HttpError
//@Router /v1/tags/{pathId} [put]
func (s *service) UpdateTag(ctx *https.Context) {
	body, err := s.body(ctx)
	if err != nil {
		ctx.JsonResponse(http.StatusBadRequest, s.NewHttpError(TagCode, TagMsg, RequestBodyErrorCode, RequestBodyErrorMsg))
		return
	}
	tagId, err := s.pathId(ctx)
	if err != nil {
		ctx.JsonResponse(http.StatusBadRequest, s.NewHttpError(TagCode, TagMsg, IDErrorCode, IDErrorMsg))
		return
	}
	param := RequestTag{}
	err = json.Unmarshal(body, &param)
	if err != nil {
		ctx.JsonResponse(http.StatusBadRequest, s.NewHttpError(TagCode, TagMsg, BodyStructureErrorCode, BodyStructureErrorMsg))
		return
	}
	err = config.Services().TagService().UpdateTag(tagId, param.NameEn, param.NameCn, param.Status)
	if err != nil {
		ctx.JsonResponse(http.StatusBadRequest, s.NewHttpError(TagCode, TagMsg, UpdateErrorCode, UpdateErrorMsg))
		return
	}
	ctx.JsonResponse(http.StatusOK, s.NewHttpOk(http.StatusOK))
}

//@Summary 删除标签
//@Produce json
//@Success 200 {object} route.HttpOk
//@Failure 400 {object} route.HttpError
//@Failure 404 {object} route.HttpError
//@Failure 500 {object} route.HttpError
//@Router /v1/tags/{pathId} [delete]
func (s *service) DeleteTag(ctx *https.Context) {
	tagId, err := s.pathId(ctx)
	if err != nil {
		ctx.JsonResponse(http.StatusBadRequest, s.NewHttpError(TagCode, TagMsg, IDErrorCode, IDErrorMsg))
		return
	}
	err = config.Services().TagService().DeleteTag(tagId)
	if err != nil {
		ctx.JsonResponse(http.StatusBadRequest, s.NewHttpError(TagCode, TagMsg, DeleteErrorCode, DeleteErrorMsg))
		return
	}
	ctx.JsonResponse(http.StatusOK, s.NewHttpOk(http.StatusOK))
}
