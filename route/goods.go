//author: richard
package route

import (
	"encoding/json"
	"github.com/advancevillage/3rd/https"
	"mms/api"
	"net/http"
)

//@Summary 新增商品
//@Produce json
//@Param language header string false "语言" default "chinese"
//@Param {} body api.Goods true "CreateGoods"
//@Success 200 {object} route.httpOk
//@Failure 400 {object} route.httpError
//@Failure 404 {object} route.httpError
//@Failure 500 {object} route.httpError
//@Router /v1/merchandises [post]
func (s *Service) CreateGoods(ctx *https.Context) {
	lang := s.language(ctx)
	body, err := s.body(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, s.newHttpError(GoodsCode, GoodsMsg, BodyErrorCode, BodyErrorMsg))
		return
	}
	param := api.Goods{}
	err = json.Unmarshal(body, &param)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, s.newHttpError(GoodsCode, GoodsMsg, ContextErrorCode, ContextErrorMsg))
		return
	}

	err = s.langService.I18n(param.Name, lang)
	if err != nil {
		s.configService.Logger.Warning(err.Error())
	}

	err = s.langService.I18n(param.Title, lang)
	if err != nil {
		s.configService.Logger.Warning(err.Error())
	}

	err = s.langService.I18n(param.Description, lang)
	if err != nil {
		s.configService.Logger.Warning(err.Error())
	}

	for i := 0; i < len(param.Tags); i++ {
		err = s.langService.I18n(param.Tags[i].Name, lang)
		if err != nil {
			s.configService.Logger.Warning(err.Error())
		}
	}

	for i := 0; i < len(param.Keywords); i++ {
		err = s.langService.I18n(param.Keywords[i].Name, lang)
		if err != nil {
			s.configService.Logger.Warning(err.Error())
		}
	}

	for i := 0; i < len(param.Materials); i++ {
		err = s.langService.I18n(param.Materials[i].Name, lang)
		if err != nil {
			s.configService.Logger.Warning(err.Error())
		}
	}

	err = s.goodsService.CreateGoods(&param)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, s.newHttpError(GoodsCode, GoodsMsg, CreateErrorCode, CreateErrorMsg))
		return
	}
	ctx.JSON(http.StatusOK, s.newHttpOk())
}
