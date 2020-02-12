//author: richard
package route

import (
	"encoding/json"
	"errors"
	"github.com/advancevillage/3rd/https"
	"mms/api"
	"net/http"
)

//@Summary 获取令牌
//@Produce json
//@Param x-language header string false "语言" default "chinese"
//@Param {} body api.Token true "LoginToken"
//@Success 200 {object} route.httpOk
//@Failure 400 {object} route.httpError
//@Failure 404 {object} route.httpError
//@Failure 500 {object} route.httpError
//@Router /v1/tokens [post]
func (s *Service) CreateToken(ctx *https.Context) {
	body, err := s.body(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, s.newHttpError(TokenCode, TokenMsg, BodyErrorCode, BodyErrorMsg))
		return
	}
	param := api.Token{}
	err = json.Unmarshal(body, &param)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, s.newHttpError(TokenCode, TokenMsg, ContextErrorCode, ContextErrorMsg))
		return
	}
	response := make(map[string]interface{})
	switch param.Category {
	case 0:
		response["token"], err = s.userService.LoginToken(param.Username)
	case 1:
		response["token"], err = s.userService.RegisterToken(param.Username)
	default:
		err = errors.New("unknown token type")
	}
	if err != nil {
		ctx.JSON(http.StatusBadRequest, s.newHttpError(TokenCode, TokenMsg, CreateErrorCode, err.Error()))
		return
	}
	ctx.JSON(http.StatusOK, response)
}

//@Summary 用户登录
//@Produce json
//@Param x-language header string false "语言" default "chinese"
//@Param username  query string true "用户名"
//@Param password  query string true "密码"
//@Param timestamp query int    true "时间戳"
//@Param token     query string true "令牌"
//@Param sign	   query string true "数字签名"
//@Success 200 {object} route.httpOk
//@Failure 400 {object} route.httpError
//@Failure 404 {object} route.httpError
//@Failure 500 {object} route.httpError
//@Router /v1/users [get]
func (s *Service) QueryUser(ctx *https.Context) {
	login := api.Login{}
	login.Username  = s.username(ctx)
	login.Password  = s.password(ctx)
	login.Token     = s.token(ctx)
	login.Timestamp = s.timestamp(ctx)
	login.Sign      = s.sign(ctx)
	u, err := s.userService.QueryUser(&login)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, s.newHttpError(UserCode, UserMsg, QueryErrorCode, err.Error()))
		return
	}

	//用户信息写入Session
	sid, err := s.sessionService.CreateUserSession(u)
	if err != nil {
		s.configService.Logger.Alert(err.Error())
		ctx.JSON(http.StatusOK, s.newHttpOk())
		return
	}

	//cookie set-cookie
	err = ctx.WriteCookie("sid", sid, "/", "localhost", false, false)
	if err != nil {
		s.configService.Logger.Alert(err.Error())
		ctx.JSON(http.StatusOK, s.newHttpOk())
		return
	}
	ctx.JSON(http.StatusOK, s.newHttpOk())
}

//@Summary 用户注册
//@Produce json
//@Param x-language header string false "语言" default "chinese"
//@Param {} body api.Register true "CreateUser"
//@Success 200 {object} route.httpOk
//@Failure 400 {object} route.httpError
//@Failure 404 {object} route.httpError
//@Failure 500 {object} route.httpError
//@Router /v1/users [post]
func (s *Service) CreateUser(ctx *https.Context) {
	body, err := s.body(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, s.newHttpError(UserCode, UserMsg, BodyErrorCode, BodyErrorMsg))
		return
	}
	param := api.Register{}
	err = json.Unmarshal(body, &param)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, s.newHttpError(UserCode, UserMsg, ContextErrorCode, ContextErrorMsg))
		return
	}
	u, err := s.userService.CreateUser(&param)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, s.newHttpError(UserCode, UserMsg, CreateErrorCode, err.Error()))
		return
	}
	//sessionId 存入cookie
	//用户信息写入Session
	sid, err := s.sessionService.CreateUserSession(u)
	if err != nil {
		s.configService.Logger.Alert(err.Error())
		ctx.JSON(http.StatusOK, s.newHttpOk())
		return
	}

	//cookie set-cookie
	err = ctx.WriteCookie("sid", sid, "/", "localhost", false, false)
	if err != nil {
		s.configService.Logger.Alert(err.Error())
		ctx.JSON(http.StatusOK, s.newHttpOk())
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
func (s *Service) CreateCart(ctx *https.Context) {
	//验证请求
	sid, err := s.sid(ctx)
	if err != nil {
		ctx.JSON(http.StatusForbidden, s.newHttpError(SessionCode, SessionMsg, QueryErrorCode, err.Error()))
		return
	}
	user, err := s.sessionService.QueryUserSession(sid)
	if err != nil {
		ctx.JSON(http.StatusForbidden, s.newHttpError(SessionCode, SessionMsg, QueryErrorCode, err.Error()))
		return
	}
	ctx.JSON(http.StatusOK, user)
}
