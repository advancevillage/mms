//author: richard
package route

import (
	"fmt"
	"github.com/advancevillage/3rd/https"
	"github.com/advancevillage/3rd/utils"
	"net/http"
)

//@Summary 新增图片
//@Produce json
//@Param x-language header string false "语言" default "chinese"
//@Param {} body api.Image true "CreateImage"
//@Success 200 {object} route.HttpOk
//@Failure 400 {object} route.HttpError
//@Failure 404 {object} route.HttpError
//@Failure 500 {object} route.HttpError
//@Router /v1/images [post]
func (s *Service) UploadImage(ctx *https.Context) {
	filename := utils.SnowFlakeIdString() + utils.RandsNumberString(4)
	uri := fmt.Sprintf("%s/%s/%s/%s", s.configService.Configure.Upload, filename[:2], filename[2:4], filename)
	uri, err := ctx.Save(uri)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, s.newHttpError(ImageCode, ImageMsg, CreateErrorCode, CreateErrorMsg))
		return
	}
	var image = struct {
		Name string `json:"name"`
		URI  string `json:"uri"`
	}{
		filename,
		uri,
	}
	ctx.JSON(200, image)
}
