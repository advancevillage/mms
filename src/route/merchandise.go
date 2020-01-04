//author: richard
package route

import "github.com/advancevillage/3rd/https"

//@Summary 查询商品列表
//@Produce json
//@Param {} body route.RequestCategory true "QueryMerchandises"
//@Success 200 {object} route.HttpOk
//@Failure 400 {object} route.HttpError
//@Failure 404 {object} route.HttpError
//@Failure 500 {object} route.HttpError
//@Router /v1/merchandises [get]
func (s *service) QueryMerchandises(ctx *https.Context) {

}

//@Summary 创建商品
//@Produce json
//@Param {} body route.RequestCategory true "CreateMerchandise"
//@Success 200 {object} route.HttpOk
//@Failure 400 {object} route.HttpError
//@Failure 404 {object} route.HttpError
//@Failure 500 {object} route.HttpError
//@Router /v1/merchandises [post]
func (s *service) CreateMerchandise(ctx *https.Context) {

}

//@Summary 更新商品
//@Produce json
//@Param {} body route.RequestCategory true "UpdateMerchandise"
//@Success 200 {object} route.HttpOk
//@Failure 400 {object} route.HttpError
//@Failure 404 {object} route.HttpError
//@Failure 500 {object} route.HttpError
//@Router /v1/merchandises/{goodsId} [put]
func (s *service) UpdateMerchandise(ctx *https.Context) {

}

//@Summary 查询商品
//@Produce json
//@Param {} body route.RequestCategory true "QueryMerchandise"
//@Success 200 {object} route.HttpOk
//@Failure 400 {object} route.HttpError
//@Failure 404 {object} route.HttpError
//@Failure 500 {object} route.HttpError
//@Router /v1/merchandises/{goodsId} [get]
func (s *service) QueryMerchandise(ctx *https.Context) {

}

//@Summary 删除商品
//@Produce json
//@Param {} body route.RequestCategory true "DeleteMerchandise"
//@Success 200 {object} route.HttpOk
//@Failure 400 {object} route.HttpError
//@Failure 404 {object} route.HttpError
//@Failure 500 {object} route.HttpError
//@Router /v1/merchandises/{goodsId} [delete]
func (s *service) DeleteMerchandise(ctx *https.Context) {

}