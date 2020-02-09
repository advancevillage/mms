//author: richard
package route

import (
	"errors"
	"fmt"
	"github.com/advancevillage/3rd/https"
	"mms/config"
	"mms/language"
	"mms/user"
	"net/http"
	"strconv"
	"strings"
)

func NewService(configService *config.Service, langService *language.Service, userService *user.Service) *Service {
	return &Service{
		configService: configService,
		langService:   langService,
		userService:   userService,
	}
}

func (s *Service) ping(ctx *https.Context) {
	ctx.JSON(http.StatusOK, "pong")
}

func (s *Service) version(ctx *https.Context) {
	version := make(map[string]string)
	version["commit"] = s.configService.Configure.Commit
	version["buildTime"] = s.configService.Configure.BuildTime
	ctx.JSON(http.StatusOK, version)
}

func (s *Service) StartRouter(mode string) error {
	var err error
	switch mode {
	case "lambda":
		err = s.startLambdaRouter()
	default:
		err = s.startHttpRouter(s.configService.Configure.HttpHost, s.configService.Configure.HttpPort)
	}
	return err
}

func (s *Service) startHttpRouter(host string, port int) error {
	server := https.NewServer(host, port, router(s), s.headerPlugin)
	err := server.StartServer()
	if err != nil {
		return err
	}
	return nil
}

func (s *Service) startLambdaRouter() error {
	server := https.NewAwsApiGatewayLambdaServer(router(s), s.headerPlugin)
	err := server.StartServer()
	if err != nil {
		return err
	}
	return nil
}

func (s *Service) headerPlugin(ctx *https.Context) {
	ctx.WriteHeader(headers)
	ctx.Next()
}

func (s *Service) newHttpError(apiCode int, apiMsg string, opCode int, opMsg string) *httpError {
	return &httpError{
		Code: fmt.Sprintf("%d%d", apiCode, opCode),
		Message: fmt.Sprintf("%s %s", apiMsg, opMsg),
	}
}

func (s *Service) newHttpOk() *httpOk {
	return &httpOk{
		Code: http.StatusOK,
		Message: OperateSuccess,
	}
}

func (s *Service) pathId(ctx *https.Context) (string, error) {
	id := ctx.Param("pathId")
	if len(id) != SnowFlakeIdLength {
		return "", errors.New(IDErrorMsg)
	}
	return id, nil
}

func (s *Service) body(ctx *https.Context) ([]byte, error) {
	return ctx.Body()
}

func (s *Service) language(ctx *https.Context) string {
	value := ctx.ReadHeader("x-language")
	return value
}

func (s *Service) page(ctx *https.Context) int {
	value := ctx.Param("page")
	page, err := strconv.Atoi(value)
	if err != nil || page < 0 || page > 50 {
		page = 0
	}
	return page
}

func (s *Service) perPage(ctx *https.Context) int {
	value := ctx.Param("perPage")
	perPage, err := strconv.Atoi(value)
	if err != nil || perPage < 0 || perPage > 100 {
		perPage = 20
	}
	return perPage
}

func (s *Service) level(ctx *https.Context) int {
	value := ctx.Param("level")
	level, err := strconv.Atoi(value)
	if err != nil || level <= 0 || level > 3 {
		level = 1
	}
	return level
}

func (s *Service) username(ctx *https.Context) string {
	value := ctx.Param("username")
	return value
}

func (s *Service) password(ctx *https.Context) string {
	value := ctx.Param("password")
	return strings.ToLower(value)
}

func (s *Service) token(ctx *https.Context) string {
	value := ctx.Param("token")
	return value
}

func (s *Service) sign(ctx *https.Context) string {
	value := ctx.Param("sign")
	return value
}

func (s *Service) timestamp(ctx *https.Context) int64 {
	value := ctx.Param("timestamp")
	timestamp, err := strconv.ParseInt(value, 10, 64)
	if err != nil {
		timestamp = 0
	}
	return timestamp
}

func (s *Service) response(items interface{}, total int64) interface{} {
	response := make(map[string]interface{})
	response["items"] = items
	response["total"] = total
	return response
}