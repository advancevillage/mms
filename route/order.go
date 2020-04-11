//author: richard
package route

import (
	"encoding/json"
	"github.com/advancevillage/3rd/https"
	"mms/api"
	"net/http"
)

//@Summary 创建订单
//@Produce json
//@Param x-language header string false "语言" default "chinese"
//@Param {} body api.Order true "CreateOrder"
//@Success 200 {object} route.httpOk
//@Failure 400 {object} route.httpError
//@Failure 404 {object} route.httpError
//@Failure 500 {object} route.httpError
//@Router /v1/orders [post]
func (s *Service) CreateOrder(ctx *https.Context) {
	//接口幂等性 原理 令牌 在结算页访问时下方 Cookie
	body, err := s.body(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, s.newHttpError(OrderCode, OrderMsg, BodyErrorCode, BodyErrorMsg))
		return
	}
	param := api.Order{}
	err = json.Unmarshal(body, &param)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, s.newHttpError(OrderCode, OrderMsg, ContextErrorCode, ContextErrorMsg))
		return
	}
	//验证用户
	sid, err := s.sid(ctx)
	if err != nil {
		ctx.JSON(http.StatusForbidden, s.newHttpError(SessionCode, SessionMsg, QueryErrorCode, err.Error()))
		return
	}
	//获取令牌
	tid, err := s.tid(ctx)
	if err != nil {
		//如果没有结算页tid 202
		ctx.JSON(http.StatusAccepted, s.newHttpError(PayTokenCode, PayTokenMsg, QueryErrorCode, err.Error()))
		return
	}
	//幂等检测
	handle, err  := s.sessionService.QueryTidSession(tid)
	if err != nil {
		ctx.JSON(http.StatusAccepted, s.newHttpError(PayTokenCode, PayTokenMsg, QueryErrorCode, err.Error()))
		return
	}
	if handle == Handling {
		//处理中
		ctx.JSON(http.StatusOK, s.newHttpOk())
		return
	}
	//TODO
	err = s.sessionService.UpdateTidSession(tid, []byte(Handling))
	//查询用户
	user, err := s.sessionService.QueryUserSession(sid)
	if err != nil {
		ctx.JSON(http.StatusForbidden, s.newHttpError(SessionCode, SessionMsg, QueryErrorCode, err.Error()))
		return
	}
	//查询库存
	if len(param.Stocks) == 0 {
		//没有订单商品则返回
		ctx.JSON(http.StatusOK, s.newHttpOk())
		return
	}
	for i := range param.Stocks {
		buy := &param.Stocks[i] 	//下单量
		sell, err := s.orderService.QueryStock(buy) //库存量
		if err != nil {
			//TODO 库存查询错误
			ctx.JSON(http.StatusInternalServerError, s.newHttpError(StockCode, StockMsg, QueryErrorCode, err.Error()))
			return
		}
		//下单量 > 库存量
		if buy.Total > sell.Total {
			ctx.JSON(http.StatusAccepted, s.newHttpError(StockCode, StockMsg, QueryErrorCode, StockNotEnough))
			return
		}
		//重设版本 CAS
		buy.CartId  = buy.Id
		buy.Id      = sell.Id
		buy.Version = sell.Version
	}

	//TODO 校验地址
	if param.Address == nil {
		ctx.JSON(http.StatusAccepted, s.newHttpError(AddressCode, AddressMsg, QueryErrorCode, InvalidAddress))
		return
	}
	err = s.orderService.CreateOrder(user, &param)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, s.newHttpError(OrderCode, OrderMsg, CreateErrorCode, err.Error()))
		return
	}
	ctx.JSON(http.StatusOK, s.newHttpOk())
}

//@Summary 添加购物车
//@Produce json
//@Param x-language header string false "语言" default "chinese"
//@Param {} body api.Cart true "CreateCart"
//@Success 200 {object} route.httpOk
//@Failure 400 {object} route.httpError
//@Failure 404 {object} route.httpError
//@Failure 500 {object} route.httpError
//@Router /v1/carts [post]


//@Summary 查询用户订单列表
//@Produce json
//@Param x-language header string false "语言" default "chinese"
//@Success 200 {object} route.httpOk
//@Failure 400 {object} route.httpError
//@Failure 404 {object} route.httpError
//@Failure 500 {object} route.httpError
//@Router /v1/orders [get]
func (s *Service) QueryOrder(ctx *https.Context) {
	page := s.page(ctx)
	perPage := s.perPage(ctx)
	//验证用户
	sid, err := s.sid(ctx)
	if err != nil {
		ctx.JSON(http.StatusForbidden, s.newHttpError(SessionCode, SessionMsg, QueryErrorCode, err.Error()))
		return
	}
	//查询用户
	user, err := s.sessionService.QueryUserSession(sid)
	if err != nil {
		ctx.JSON(http.StatusForbidden, s.newHttpError(SessionCode, SessionMsg, QueryErrorCode, err.Error()))
		return
	}
	orders, total, err := s.orderService.QueryOrders(user, page, perPage)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, s.newHttpError(OrderCode, OrderMsg, QueryErrorCode, err.Error()))
		return
	}
	ctx.JSON(http.StatusOK, s.response(orders, total))
}

//@Summary 新增地址
//@Produce json
//@Param x-language header string false "语言" default "chinese"
//@Param {} body api.Address true "CreateAddress"
//@Success 200 {object} route.httpOk
//@Failure 400 {object} route.httpError
//@Failure 404 {object} route.httpError
//@Failure 500 {object} route.httpError
//@Router /v1/payToken [post]
func (s *Service) CreatePayToken(ctx *https.Context) {
	//验证用户
	sid, err := s.sid(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, s.newHttpError(SessionCode, SessionMsg, QueryErrorCode, err.Error()))
		return
	}
	//查询用户
	user, err := s.sessionService.QueryUserSession(sid)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, s.newHttpError(SessionCode, SessionMsg, QueryErrorCode, err.Error()))
		return
	}
	token, err := s.orderService.CreatePayToken(user)
	ctx.JSON(http.StatusOK, token)
}

//@Summary 订单令牌
//@Produce json
//@Param x-language header string false "语言" default "chinese"
//@Success 200 {object} route.httpOk
//@Failure 400 {object} route.httpError
//@Failure 404 {object} route.httpError
//@Failure 500 {object} route.httpError
//@Router /v1/orderToken [post]
func (s *Service) CreateOrderToken(ctx *https.Context) {
	//验证用户
	sid, err := s.sid(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, s.newHttpError(SessionCode, SessionMsg, QueryErrorCode, err.Error()))
		return
	}
	//查询用户
	user, err := s.sessionService.QueryUserSession(sid)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, s.newHttpError(SessionCode, SessionMsg, QueryErrorCode, err.Error()))
		return
	}
	tid, err := s.sessionService.CreateTidSession(user, []byte(PendingHandle))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, s.newHttpError(OrderTokenCode, OrderTokenMsg, CreateErrorCode, err.Error()))
		return
	}
	//cookie set-cookie
	err = ctx.WriteCookie("tid", tid, "/", "localhost", false, true)
	if err != nil {
		ctx.JSON(http.StatusOK, s.newHttpOk())
		return
	}
	ctx.JSON(http.StatusOK, s.newHttpOk())
}
