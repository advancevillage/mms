//author: richard
package route

import (
	"encoding/json"
	"github.com/advancevillage/3rd/https"
	"mms/src/config"
	"net/http"
)

//@Summary 创建品牌
//@Produce json
//@Param {} body route.RequestBrand true "CreateBrand"
//@Success 200 {object} route.HttpOk
//@Failure 400 {object} route.HttpError
//@Failure 404 {object} route.HttpError
//@Failure 500 {object} route.HttpError
//@Router /v1/brands [post]
func (s *service) CreateBrand(ctx *https.Context) {
	body, err := s.body(ctx)
	if err != nil {
		ctx.JsonResponse(http.StatusBadRequest, s.NewHttpError(BrandCode, BrandMsg, RequestBodyErrorCode, RequestBodyErrorMsg))
		return
	}
	param := RequestBrand{}
	err = json.Unmarshal(body, &param)
	if err != nil {
		ctx.JsonResponse(http.StatusBadRequest, s.NewHttpError(BrandCode, BrandMsg, BodyStructureErrorCode, BodyStructureErrorMsg))
		return
	}
	err = config.Services().BrandService().CreateBrand(param.BrandName.English, param.BrandName.Chinese)
	if err != nil {
		ctx.JsonResponse(http.StatusBadRequest, s.NewHttpError(BrandCode, BrandMsg, CreateErrorCode, CreateErrorMsg))
		return
	}
	ctx.JsonResponse(http.StatusOK, s.NewHttpOk(http.StatusOK))
}

//@Summary 查询品牌列表
//@Produce json
//@Param page    query int false "页码" default "0"
//@Param perPage query int false "每页条数" default "20"
//@Param status  query int false "状态"
//@Success 200 {object} route.HttpOk
//@Failure 400 {object} route.HttpError
//@Failure 404 {object} route.HttpError
//@Failure 500 {object} route.HttpError
//@Router /v1/brands [get]
func (s *service) QueryBrands(ctx *https.Context) {
	page    := s.page(ctx)
	perPage := s.perPage(ctx)
	status  := s.status(ctx)
	brands, total, err := config.Services().BrandService().QueryBrands(status, page, perPage)
	if err != nil {
		ctx.JsonResponse(http.StatusBadRequest, s.NewHttpError(BrandCode, BrandMsg, QueryErrorCode, QueryErrorMsg))
		return
	}
	ctx.JsonResponse(http.StatusOK, s.response(brands, total))
}

//@Summary 查询品牌
//@Produce json
//@Success 200 {object} route.HttpOk
//@Failure 400 {object} route.HttpError
//@Failure 404 {object} route.HttpError
//@Failure 500 {object} route.HttpError
//@Router /v1/brands/{pathId} [get]
func (s *service) QueryBrand(ctx *https.Context) {
	brandId, err := s.pathId(ctx)
	if err != nil {
		ctx.JsonResponse(http.StatusBadRequest, s.NewHttpError(BrandCode, BrandMsg, IDErrorCode, IDErrorMsg))
		return
	}
	brand, err := config.Services().BrandService().QueryBrandById(brandId)
	if err != nil {
		ctx.JsonResponse(http.StatusBadRequest, s.NewHttpError(BrandCode, BrandMsg, QueryErrorCode, QueryErrorMsg))
		return
	}
	ctx.JsonResponse(http.StatusOK, brand)
}

//@Summary 更新品牌
//@Produce json
//@Param {} body route.RequestBrand true "UpdateBrand"
//@Success 200 {object} route.HttpOk
//@Failure 400 {object} route.HttpError
//@Failure 404 {object} route.HttpError
//@Failure 500 {object} route.HttpError
//@Router /v1/brands/{pathId} [put]
func (s *service) UpdateBrand(ctx *https.Context) {
	body, err := s.body(ctx)
	if err != nil {
		ctx.JsonResponse(http.StatusBadRequest, s.NewHttpError(BrandCode, BrandMsg, RequestBodyErrorCode, RequestBodyErrorMsg))
		return
	}
	brandId, err := s.pathId(ctx)
	if err != nil {
		ctx.JsonResponse(http.StatusBadRequest, s.NewHttpError(BrandCode, BrandMsg, IDErrorCode, IDErrorMsg))
		return
	}
	param := RequestBrand{}
	err = json.Unmarshal(body, &param)
	if err != nil {
		ctx.JsonResponse(http.StatusBadRequest, s.NewHttpError(BrandCode, BrandMsg, BodyStructureErrorCode, BodyStructureErrorMsg))
		return
	}
	err = config.Services().BrandService().UpdateBrand(brandId, param.BrandName.English, param.BrandName.Chinese, param.Status)
	if err != nil {
		ctx.JsonResponse(http.StatusBadRequest, s.NewHttpError(BrandCode, BrandMsg, UpdateErrorCode, UpdateErrorMsg))
		return
	}
	ctx.JsonResponse(http.StatusOK, s.NewHttpOk(http.StatusOK))
}

//@Summary 删除品牌
//@Produce json
//@Success 200 {object} route.HttpOk
//@Failure 400 {object} route.HttpError
//@Failure 404 {object} route.HttpError
//@Failure 500 {object} route.HttpError
//@Router /v1/brands/{pathId} [delete]
func (s *service) DeleteBrand(ctx *https.Context) {
	brandId, err := s.pathId(ctx)
	if err != nil {
		ctx.JsonResponse(http.StatusBadRequest, s.NewHttpError(BrandCode, BrandMsg, IDErrorCode, IDErrorMsg))
		return
	}
	err = config.Services().BrandService().DeleteBrand(brandId)
	if err != nil {
		ctx.JsonResponse(http.StatusBadRequest, s.NewHttpError(BrandCode, BrandMsg, DeleteErrorCode, DeleteErrorMsg))
		return
	}
	ctx.JsonResponse(http.StatusOK, s.NewHttpOk(http.StatusOK))
}