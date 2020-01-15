//author: richard
package route

import (
	"errors"
	"fmt"
	"github.com/advancevillage/3rd/https"
	"mms/src/config"
	"net/http"
	"strconv"
)

type service struct {}

func NewApiService() *service {
	return &service{}
}

//@Summary 显示当前服务的版本和代码版本号
//@Produce json
//@Success 200 {object} route.ResponseVersion
//@Failure 400 {object} route.HttpError
//@Failure 404 {object} route.HttpError
//@Failure 500 {object} route.HttpError
//@Router /v1/services/version [get]
func (s *service) Version(ctx *https.Context) {
	version := ResponseVersion{}
	version.Info = config.Services().Version()
	ctx.JsonResponse(http.StatusOK, version)
}

func (s *service) Test(ctx *https.Context) {
	ctx.JsonResponse(http.StatusOK, nil)
}


func (s *service) NewHttpError(apiCode int, apiMsg string, opCode int, opMsg string) *HttpError {
	return &HttpError{
		Code: fmt.Sprintf("%d%d", apiCode, opCode),
		Message: fmt.Sprintf("%s %s", apiMsg, opMsg),
	}
}

func (s *service) NewHttpOk(statusCode int) *HttpOk {
	return &HttpOk{
		Code: statusCode,
		Message: OperateSuccess,
	}
}

func (s *service) page(ctx *https.Context) int {
	value := ctx.Param("page")
	page, err := strconv.Atoi(value)
	if err != nil || page < 0 || page > 50 {
		page = 0
	}
	return page
}

func (s *service) perPage(ctx *https.Context) int {
	value := ctx.Param("perPage")
	perPage, err := strconv.Atoi(value)
	if err != nil || perPage < 0 || perPage > 100 {
		perPage = 20
	}
	return perPage
}

func (s *service) level(ctx *https.Context) int {
	value := ctx.Param("level")
	level, err := strconv.Atoi(value)
	if err != nil || level <= 0 || level > 3 {
		level = 1
	}
	return level
}

func (s *service) status(ctx *https.Context) int {
	value := ctx.Param("status")
	status, err := strconv.Atoi(value)
	if err != nil {
		status = -1
	}
	return status
}

func (s *service) pathId(ctx *https.Context) (string, error) {
	id := ctx.Param("pathId")
	if len(id) != SnowFlakeIdLength {
		return "", errors.New(IDErrorMsg)
	}
	return id, nil
}

func (s *service) body(ctx *https.Context) ([]byte, error) {
	return ctx.Body()
}

func (s *service) response(items interface{}, total int64) interface{} {
	response := make(map[string]interface{})
	response["items"] = items
	response["total"] = total
	return response
}