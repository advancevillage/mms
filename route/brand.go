//author: richard
package route

import (
	"encoding/json"
	"github.com/advancevillage/3rd/https"
	"mms/api"
	"net/http"
)

//@Summary 创建品牌
//@Produce json
//@Param x-language header string false "语言" default "chinese"
//@Param {} body api.Brand true "CreateBrand"
//@Success 200 {object} route.httpOk
//@Failure 400 {object} route.httpError
//@Failure 404 {object} route.httpError
//@Failure 500 {object} route.httpError
//@Router /v1/brands [post]
func (s *Service) CreateBrand(ctx *https.Context) {
	lang := s.language(ctx)
	body, err := s.body(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, s.newHttpError(BrandCode, BrandMsg, BodyErrorCode, BodyErrorMsg))
		return
	}
	param := api.Brand{}
	err = json.Unmarshal(body, &param)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, s.newHttpError(BrandCode, BrandMsg, ContextErrorCode, ContextErrorMsg))
		return
	}

	err = s.langService.I18n(param.Name, lang)
	if err != nil {
		s.configService.Logger.Warning(err.Error())
	}

	err = s.brandService.CreateBrand(&param)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, s.newHttpError(BrandCode, BrandMsg, CreateErrorCode, CreateErrorMsg))
		return
	}
	ctx.JSON(http.StatusOK, s.newHttpOk())
}

//@Summary 查询品牌列表
//@Produce json
//@Param x-language header string false "语言" default "chinese"
//@Param page    query int false "页码" default "0"
//@Param perPage query int false "每页条数" default "20"
//@Param status  query int false "状态"
//@Success 200 {object} route.httpOk
//@Failure 400 {object} route.httpError
//@Failure 404 {object} route.httpError
//@Failure 500 {object} route.httpError
//@Router /v1/brands [get]
func (s *Service) QueryBrands(ctx *https.Context) {
	page    := s.page(ctx)
	perPage := s.perPage(ctx)
	brands, total, err := s.brandService.QueryBrands(page, perPage)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, s.newHttpError(BrandCode, BrandMsg, QueryErrorCode, QueryErrorMsg))
		return
	}
	ctx.JSON(http.StatusOK, s.response(brands, total))
}

//@Summary 查询品牌
//@Produce json
//@Param x-language header string false "语言" default "chinese"
//@Success 200 {object} route.httpOk
//@Failure 400 {object} route.httpError
//@Failure 404 {object} route.httpError
//@Failure 500 {object} route.httpError
//@Router /v1/brands/{pathId} [get]
func (s *Service) QueryBrand(ctx *https.Context) {
	brandId, err := s.pathId(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, s.newHttpError(BrandCode, BrandMsg, IDErrorCode, IDErrorMsg))
		return
	}
	brand, err := s.brandService.QueryBrandById(brandId)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, s.newHttpError(BrandCode, BrandMsg, QueryErrorCode, QueryErrorMsg))
		return
	}
	ctx.JSON(http.StatusOK, brand)
}

//@Summary 更新品牌
//@Produce json
//@Param x-language header string false "语言" default "chinese"
//@Param {} body api.Brand true "UpdateBrand"
//@Success 200 {object} route.httpOk
//@Failure 400 {object} route.httpError
//@Failure 404 {object} route.httpError
//@Failure 500 {object} route.httpError
//@Router /v1/brands/{pathId} [put]
func (s *Service) UpdateBrand(ctx *https.Context) {
	body, err := s.body(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, s.newHttpError(BrandCode, BrandMsg, BodyErrorCode, BodyErrorMsg))
		return
	}
	brandId, err := s.pathId(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, s.newHttpError(BrandCode, BrandMsg, IDErrorCode, IDErrorMsg))
		return
	}
	param := api.Brand{}
	err = json.Unmarshal(body, &param)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, s.newHttpError(BrandCode, BrandMsg, ContextErrorCode, ContextErrorMsg))
		return
	}
	param.Id = brandId
	err = s.brandService.UpdateBrand(&param)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, s.newHttpError(BrandCode, BrandMsg, UpdateErrorCode, UpdateErrorMsg))
		return
	}
	ctx.JSON(http.StatusOK, s.newHttpOk())
}

//@Summary 删除品牌
//@Produce json
//@Param x-language header string false "语言" default "chinese"
//@Success 200 {object} route.httpOk
//@Failure 400 {object} route.httpError
//@Failure 404 {object} route.httpError
//@Failure 500 {object} route.httpError
//@Router /v1/brands/{pathId} [delete]
func (s *Service) DeleteBrand(ctx *https.Context) {
	brandId, err := s.pathId(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, s.newHttpError(BrandCode, BrandMsg, IDErrorCode, IDErrorMsg))
		return
	}
	err = s.brandService.DeleteBrand(brandId)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, s.newHttpError(BrandCode, BrandMsg, DeleteErrorCode, DeleteErrorMsg))
		return
	}
	ctx.JSON(http.StatusOK, s.newHttpOk())
}
