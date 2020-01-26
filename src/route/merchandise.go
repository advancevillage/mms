//author: richard
package route

import (
	"encoding/json"
	"github.com/advancevillage/3rd/https"
	"mms/src/config"
	"net/http"
)

//@Summary 新增商品
//@Produce json
//@Param language header string false "语言" default "chinese"
//@Param {} body route.RequestGoods true "CreateGoods"
//@Success 200 {object} route.HttpOk
//@Failure 400 {object} route.HttpError
//@Failure 404 {object} route.HttpError
//@Failure 500 {object} route.HttpError
//@Router /v1/merchandises [post]
func (s *service) CreateMerchandise(ctx *https.Context) {
	lang := s.language(ctx)
	body, err := s.body(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, s.NewHttpError(MerchandiseCode, MerchandiseMsg, RequestBodyErrorCode, RequestBodyErrorMsg))
		return
	}
	param := RequestGoods{}
	err = json.Unmarshal(body, &param)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, s.NewHttpError(MerchandiseCode, MerchandiseMsg, BodyStructureErrorCode, BodyStructureErrorMsg))
		return
	}
	//名称
	param.Name.Multi(lang, config.Services().TranslateService(), config.Services().LogService())
	//标题
	param.Title.Multi(lang, config.Services().TranslateService(), config.Services().LogService())
	//描述
	param.Description.Multi(lang, config.Services().TranslateService(), config.Services().LogService())
	//关键字
	for i := 0; i < len(param.Keywords); i++ {
		param.Keywords[i].Multi(lang, config.Services().TranslateService(), config.Services().LogService())
	}
	//标签
	for i := 0; i < len(param.Tags); i++ {
		param.Tags[i].Multi(lang, config.Services().TranslateService(), config.Services().LogService())
	}
	//产地
	param.Origin.Multi(lang, config.Services().TranslateService(), config.Services().LogService())
	//材质
	for i := 0; i < len(param.Materials); i++ {
		param.Materials[i].Multi(lang, config.Services().TranslateService(), config.Services().LogService())
	}
	//
	err = config.Services().MerchandiseService().CreateManufacturer(&param.Title, &param.Description, param.CostPrice)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, s.NewHttpError(MerchandiseCode, MerchandiseMsg, CreateErrorCode, CreateErrorMsg))
		return
	}
	ctx.JSON(http.StatusOK, s.NewHttpOk(http.StatusOK))
}

//@Summary 查询商品列表
//@Produce json
//@Param language header string false "语言" default "chinese"
//@Param page    query int false "页码" default "0"
//@Param perPage query int false "每页条数" default "20"
//@Param status  query int false "状态"
//@Success 200 {object} route.HttpOk
//@Failure 400 {object} route.HttpError
//@Failure 404 {object} route.HttpError
//@Failure 500 {object} route.HttpError
//@Router /v1/merchandises [get]
func (s *service) QueryMerchandises(ctx *https.Context) {
	page    := s.page(ctx)
	perPage := s.perPage(ctx)
	status  := s.status(ctx)
	goods, total, err := config.Services().MerchandiseService().QueryManufacturers(status, page, perPage)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, s.NewHttpError(MerchandiseCode, MerchandiseMsg, QueryErrorCode, QueryErrorMsg))
		return
	}
	ctx.JSON(http.StatusOK, s.response(goods, total))
}

//@Summary 查询商品
//@Produce json
//@Param language header string false "语言" default "chinese"
//@Success 200 {object} route.HttpOk
//@Failure 400 {object} route.HttpError
//@Failure 404 {object} route.HttpError
//@Failure 500 {object} route.HttpError
//@Router /v1/merchandises/{pathId} [get]
func (s *service) QueryMerchandise(ctx *https.Context) {
	goodsId, err := s.pathId(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, s.NewHttpError(MerchandiseCode, MerchandiseMsg, IDErrorCode, IDErrorMsg))
		return
	}
	goods, err := config.Services().MerchandiseService().QueryManufacturerById(goodsId)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, s.NewHttpError(MerchandiseCode, MerchandiseMsg, QueryErrorCode, QueryErrorMsg))
		return
	}
	ctx.JSON(http.StatusOK, goods)
}

//@Summary 删除商品
//@Produce json
//@Param language header string false "语言" default "chinese"
//@Success 200 {object} route.HttpOk
//@Failure 400 {object} route.HttpError
//@Failure 404 {object} route.HttpError
//@Failure 500 {object} route.HttpError
//@Router /v1/merchandises/{pathId} [delete]
func (s *service) DeleteMerchandise(ctx *https.Context) {
	goodsId, err := s.pathId(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, s.NewHttpError(MerchandiseCode, MerchandiseMsg, IDErrorCode, IDErrorMsg))
		return
	}
	err = config.Services().MerchandiseService().DeleteManufacturer(goodsId)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, s.NewHttpError(MerchandiseCode, MerchandiseMsg, DeleteErrorCode, DeleteErrorMsg))
		return
	}
	ctx.JSON(http.StatusOK, s.NewHttpOk(http.StatusOK))
}