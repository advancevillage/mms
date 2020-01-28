//author: richard
package route

import (
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
	body := api.Goods{}
	ctx.JSON(http.StatusOK, body)
}
