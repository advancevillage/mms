//author: richard
package route

import (
	"github.com/advancevillage/3rd/https"
	"mms/src/config"
	"net/http"
)

type service struct {}

func NewApiService() *service {
	return &service{}
}

// @Summary 显示当前服务的版本和代码版本号
// @Produce json
// @Success 200 {object} ResponseVersion
// @Failure 400 {object} HttpError
// @Failure 404 {object} HttpError
// @Failure 500 {object} HttpError
// @Router /v1/merchandises/version [get]
func (s *service) Version(ctx *https.Context) {
	version := ResponseVersion{}
	version.Info = config.GetMMSObject().GetVersion()
	ctx.JsonResponse(http.StatusOK, version)
}
