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
//@Success 200 {object} route.httpOk
//@Failure 400 {object} route.httpError
//@Failure 404 {object} route.httpError
//@Failure 500 {object} route.httpError
//@Router /v1/images [post]
func (s *Service) UploadImage(ctx *https.Context) {
	suffix := utils.SnowFlakeIdString() + utils.RandsNumberString(4)
    length := len(suffix)
	filename := fmt.Sprintf("%s/%s/%s/%s", s.configService.Configure.Upload, suffix[length-4:length-2], suffix[length-2:length], suffix)
	filename, err := ctx.Save(filename)
	for i := len(filename) - 1; i > 0; i-- {
		if filename[i] == '/' {
			filename = filename[i + 1:]
			break
		} else {
			continue
		}
	}
	uri := fmt.Sprintf("%s/%s/%s/%s", "images", suffix[length-4:length-2], suffix[length-2:length], filename)
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
