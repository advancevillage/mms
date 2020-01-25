//author: richard
package route

import (
	"encoding/json"
	"github.com/advancevillage/3rd/https"
	"mms/src/config"
	"net/http"
)

//@Summary 新增生产商
//@Produce json
//@Param x-language header string false "语言" default "chinese"
//@Param {} body route.RequestManufacturer true "CreateManufacturer"
//@Success 200 {object} route.HttpOk
//@Failure 400 {object} route.HttpError
//@Failure 404 {object} route.HttpError
//@Failure 500 {object} route.HttpError
//@Router /v1/manufacturers [post]
func (s *service) CreateManufacturer(ctx *https.Context) {
	lang := s.language(ctx)
	body, err := s.body(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, s.NewHttpError(ManufacturerCode, ManufacturerMsg, RequestBodyErrorCode, RequestBodyErrorMsg))
		return
	}
	param := RequestManufacturer{}
	err = json.Unmarshal(body, &param)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, s.NewHttpError(ManufacturerCode, ManufacturerMsg, BodyStructureErrorCode, BodyStructureErrorMsg))
		return
	}
	param.Name.Multi(lang, config.Services().TranslateService(), config.Services().LogService())
	param.Address.Multi(lang, config.Services().TranslateService(), config.Services().LogService())
	err = config.Services().ManufacturerService().CreateManufacturer(param.Concat, param.Phone, param.Email, &param.Name, &param.Address)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, s.NewHttpError(ManufacturerCode, ManufacturerMsg, CreateErrorCode, CreateErrorMsg))
		return
	}
	ctx.JSON(http.StatusOK, s.NewHttpOk(http.StatusOK))
}

//@Summary 查询生产商列表
//@Produce json
//@Param x-language header string false "语言" default "chinese"
//@Param page    query int false "页码" default "0"
//@Param perPage query int false "每页条数" default "20"
//@Param status  query int false "状态"
//@Success 200 {object} route.HttpOk
//@Failure 400 {object} route.HttpError
//@Failure 404 {object} route.HttpError
//@Failure 500 {object} route.HttpError
//@Router /v1/manufacturers [get]
func (s *service) QueryManufacturers(ctx *https.Context) {
	page    := s.page(ctx)
	perPage := s.perPage(ctx)
	status  := s.status(ctx)
	manufacturers, total, err := config.Services().ManufacturerService().QueryManufacturers(status, page, perPage)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, s.NewHttpError(ManufacturerCode, ManufacturerMsg, QueryErrorCode, QueryErrorMsg))
		return
	}
	ctx.JSON(http.StatusOK, s.response(manufacturers, total))
}

//@Summary 查询生产商
//@Produce json
//@Param x-language header string false "语言" default "chinese"
//@Success 200 {object} route.HttpOk
//@Failure 400 {object} route.HttpError
//@Failure 404 {object} route.HttpError
//@Failure 500 {object} route.HttpError
//@Router /v1/manufacturers/{pathId} [get]
func (s *service) QueryManufacturer(ctx *https.Context) {
	manufacturerId, err := s.pathId(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, s.NewHttpError(ManufacturerCode, ManufacturerMsg, IDErrorCode, IDErrorMsg))
		return
	}
	manufacturer, err := config.Services().ManufacturerService().QueryManufacturerById(manufacturerId)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, s.NewHttpError(ManufacturerCode, ManufacturerMsg, QueryErrorCode, QueryErrorMsg))
		return
	}
	ctx.JSON(http.StatusOK, manufacturer)
}

//@Summary 更新生产商
//@Produce json
//@Param x-language header string false "语言" default "chinese"
//@Param {} body route.RequestManufacturer true "UpdateManufacturer"
//@Success 200 {object} route.HttpOk
//@Failure 400 {object} route.HttpError
//@Failure 404 {object} route.HttpError
//@Failure 500 {object} route.HttpError
//@Router /v1/manufacturers/{pathId} [put]
func (s *service) UpdateManufacturer(ctx *https.Context) {
	body, err := s.body(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, s.NewHttpError(ManufacturerCode, ManufacturerMsg, RequestBodyErrorCode, RequestBodyErrorMsg))
		return
	}
	manufacturerId, err := s.pathId(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, s.NewHttpError(ManufacturerCode, ManufacturerMsg, IDErrorCode, IDErrorMsg))
		return
	}
	param := RequestManufacturer{}
	err = json.Unmarshal(body, &param)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, s.NewHttpError(ManufacturerCode, ManufacturerMsg, BodyStructureErrorCode, BodyStructureErrorMsg))
		return
	}
	err = config.Services().ManufacturerService().UpdateManufacturer(manufacturerId, param.Phone, param.Email, param.Concat, param.Status)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, s.NewHttpError(ManufacturerCode, ManufacturerMsg, UpdateErrorCode, UpdateErrorMsg))
		return
	}
	ctx.JSON(http.StatusOK, s.NewHttpOk(http.StatusOK))
}

//@Summary 删除生产商
//@Produce json
//@Param x-language header string false "语言" default "chinese"
//@Success 200 {object} route.HttpOk
//@Failure 400 {object} route.HttpError
//@Failure 404 {object} route.HttpError
//@Failure 500 {object} route.HttpError
//@Router /v1/manufacturers/{pathId} [delete]
func (s *service) DeleteManufacturer(ctx *https.Context) {
	manufacturerId, err := s.pathId(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, s.NewHttpError(ManufacturerCode, ManufacturerMsg, IDErrorCode, IDErrorMsg))
		return
	}
	err = config.Services().ManufacturerService().DeleteManufacturer(manufacturerId)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, s.NewHttpError(ManufacturerCode, ManufacturerMsg, DeleteErrorCode, DeleteErrorMsg))
		return
	}
	ctx.JSON(http.StatusOK, s.NewHttpOk(http.StatusOK))
}