//author: richard
package route

import (
	"errors"
	"fmt"
	"github.com/advancevillage/3rd/https"
	"mms/brand"
	"mms/category"
	"mms/color"
	"mms/config"
	"mms/goods"
	"mms/language"
	"mms/manufacturer"
	"mms/size"
	"net/http"
	"strconv"
)

func NewService(configService *config.Service, langService *language.Service, goodsService *goods.Service, colorService *color.Service, sizeService *size.Service, brandService *brand.Service, categoryService *category.Service, manufacturerService *manufacturer.Service) *Service {
	return &Service{
		configService: configService,
		langService:   langService,
		goodsService: goodsService,
		colorService: colorService,
		sizeService: sizeService,
		brandService: brandService,
		categoryService: categoryService,
		manufacturerService: manufacturerService,
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

func (s *Service) group(ctx *https.Context) string {
	value := ctx.Param("group")
	return value
}

func (s *Service) response(items interface{}, total int64) interface{} {
	response := make(map[string]interface{})
	response["items"] = items
	response["total"] = total
	return response
}