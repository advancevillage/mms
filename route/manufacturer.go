//author: richard
package route

import (
	"encoding/json"
	"github.com/advancevillage/3rd/https"
	"mms/api"
	"net/http"
)

//@Summary 新增生产商
//@Produce json
//@Param x-language header string false "语言" default "chinese"
//@Param {} body api.Manufacturer true "CreateManufacturer"
//@Success 200 {object} route.httpOk
//@Failure 400 {object} route.httpError
//@Failure 404 {object} route.httpError
//@Failure 500 {object} route.httpError
//@Router /v1/manufacturers [post]
func (s *Service) CreateManufacturer(ctx *https.Context) {
	lang := s.language(ctx)
	body, err := s.body(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, s.newHttpError(ManufacturerCode, ManufacturerMsg, BodyErrorCode, BodyErrorMsg))
		return
	}
	param := api.Manufacturer{}
	err = json.Unmarshal(body, &param)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, s.newHttpError(ManufacturerCode, ManufacturerMsg, ContextErrorCode, ContextErrorMsg))
		return
	}

	err = s.langService.I18n(param.Address, lang)
	if err != nil {
		s.configService.Logger.Warning(err.Error())
	}

	err = s.langService.I18n(param.Name, lang)
	if err != nil {
		s.configService.Logger.Warning(err.Error())
	}

	err = s.manufacturerService.CreateManufacturer(&param)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, s.newHttpError(ManufacturerCode, ManufacturerMsg, CreateErrorCode, CreateErrorMsg))
		return
	}
	ctx.JSON(http.StatusOK, s.newHttpOk())
}

//@Summary 查询生产商列表
//@Produce json
//@Param x-language header string false "语言" default "chinese"
//@Param page    query int false "页码" default "0"
//@Param perPage query int false "每页条数" default "20"
//@Param status  query int false "状态"
//@Success 200 {object} route.httpOk
//@Failure 400 {object} route.httpError
//@Failure 404 {object} route.httpError
//@Failure 500 {object} route.httpError
//@Router /v1/manufacturers [get]
func (s *Service) QueryManufacturers(ctx *https.Context) {
	page    := s.page(ctx)
	perPage := s.perPage(ctx)
	manufacturers, total, err := s.manufacturerService.QueryManufacturers(page, perPage)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, s.newHttpError(ManufacturerCode, ManufacturerMsg, QueryErrorCode, QueryErrorMsg))
		return
	}
	ctx.JSON(http.StatusOK, s.response(manufacturers, total))
}

//@Summary 查询生产商
//@Produce json
//@Param x-language header string false "语言" default "chinese"
//@Success 200 {object} route.httpOk
//@Failure 400 {object} route.httpError
//@Failure 404 {object} route.httpError
//@Failure 500 {object} route.httpError
//@Router /v1/manufacturers/{pathId} [get]
func (s *Service) QueryManufacturer(ctx *https.Context) {
	manufacturerId, err := s.pathId(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, s.newHttpError(ManufacturerCode, ManufacturerMsg, IDErrorCode, IDErrorMsg))
		return
	}
	manufacturer, err := s.manufacturerService.QueryManufacturerById(manufacturerId)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, s.newHttpError(ManufacturerCode, ManufacturerMsg, QueryErrorCode, QueryErrorMsg))
		return
	}
	ctx.JSON(http.StatusOK, manufacturer)
}

//@Summary 更新生产商
//@Produce json
//@Param x-language header string false "语言" default "chinese"
//@Param {} body api.Manufacturer true "UpdateManufacturer"
//@Success 200 {object} route.httpOk
//@Failure 400 {object} route.httpError
//@Failure 404 {object} route.httpError
//@Failure 500 {object} route.httpError
//@Router /v1/manufacturers/{pathId} [put]
func (s *Service) UpdateManufacturer(ctx *https.Context) {
	body, err := s.body(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, s.newHttpError(ManufacturerCode, ManufacturerMsg, BodyErrorCode, BodyErrorMsg))
		return
	}
	manufacturerId, err := s.pathId(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, s.newHttpError(ManufacturerCode, ManufacturerMsg, IDErrorCode, IDErrorMsg))
		return
	}
	param := api.Manufacturer{}
	err = json.Unmarshal(body, &param)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, s.newHttpError(ManufacturerCode, ManufacturerMsg, ContextErrorCode, ContextErrorMsg))
		return
	}

	param.Id = manufacturerId

	err = s.manufacturerService.UpdateManufacturer(&param)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, s.newHttpError(ManufacturerCode, ManufacturerMsg, UpdateErrorCode, UpdateErrorMsg))
		return
	}
	ctx.JSON(http.StatusOK, s.newHttpOk())
}

//@Summary 删除生产商
//@Produce json
//@Param x-language header string false "语言" default "chinese"
//@Success 200 {object} route.httpOk
//@Failure 400 {object} route.httpError
//@Failure 404 {object} route.httpError
//@Failure 500 {object} route.httpError
//@Router /v1/manufacturers/{pathId} [delete]
func (s *Service) DeleteManufacturer(ctx *https.Context) {
	manufacturerId, err := s.pathId(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, s.newHttpError(ManufacturerCode, ManufacturerMsg, IDErrorCode, IDErrorMsg))
		return
	}
	err = s.manufacturerService.DeleteManufacturer(manufacturerId)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, s.newHttpError(ManufacturerCode, ManufacturerMsg, DeleteErrorCode, DeleteErrorMsg))
		return
	}
	ctx.JSON(http.StatusOK, s.newHttpOk())
}

