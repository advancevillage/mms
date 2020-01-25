//author: richard
package route

import (
	"encoding/json"
	"fmt"
	"github.com/advancevillage/3rd/https"
	"github.com/advancevillage/3rd/utils"
	"mms/src/config"
	"net/http"
)

//@Summary 新增图片
//@Produce json
//@Param language header string false "语言" default "chinese"
//@Param {} body route.RequestImage true "CreateImage"
//@Success 200 {object} route.HttpOk
//@Failure 400 {object} route.HttpError
//@Failure 404 {object} route.HttpError
//@Failure 500 {object} route.HttpError
//@Router /v1/images [post]
func (s *service) UploadImage(ctx *https.Context) {
	filename := utils.SnowFlakeIdString() + utils.RandsNumberString(8)
	uri := fmt.Sprintf("%s/%s/%s/%s", "images", filename[:2], filename[2:4], filename)
	uri, err := ctx.Save(uri)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, s.NewHttpError(ImageCode, ImageMsg, CreateErrorCode, CreateErrorMsg))
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

//@Summary 查询图片列表
//@Produce json
//@Param language header string false "语言" default "chinese"
//@Param page    query int false "页码" default "0"
//@Param perPage query int false "每页条数" default "20"
//@Param status  query int false "状态"
//@Success 200 {object} route.HttpOk
//@Failure 400 {object} route.HttpError
//@Failure 404 {object} route.HttpError
//@Failure 500 {object} route.HttpError
//@Router /v1/images [get]
func (s *service) QueryImages(ctx *https.Context) {
	page    := s.page(ctx)
	perPage := s.perPage(ctx)
	status  := s.status(ctx)
	images, total, err := config.Services().ImageService().QueryImages(status, page, perPage)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, s.NewHttpError(ImageCode, ImageMsg, QueryErrorCode, QueryErrorMsg))
		return
	}
	ctx.JSON(http.StatusOK, s.response(images, total))
}

//@Summary 查询图片
//@Produce json
//@Param language header string false "语言" default "chinese"
//@Success 200 {object} route.HttpOk
//@Failure 400 {object} route.HttpError
//@Failure 404 {object} route.HttpError
//@Failure 500 {object} route.HttpError
//@Router /v1/images/{pathId} [get]
func (s *service) QueryImage(ctx *https.Context) {
	imageId, err := s.pathId(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, s.NewHttpError(ImageCode, ImageMsg, IDErrorCode, IDErrorMsg))
		return
	}
	color, err := config.Services().ImageService().QueryImageById(imageId)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, s.NewHttpError(ImageCode, ImageMsg, QueryErrorCode, QueryErrorMsg))
		return
	}
	ctx.JSON(http.StatusOK, color)
}

//@Summary 更新图片
//@Produce json
//@Param language header string false "语言" default "chinese"
//@Param {} body route.RequestImage true "UpdateImage"
//@Success 200 {object} route.HttpOk
//@Failure 400 {object} route.HttpError
//@Failure 404 {object} route.HttpError
//@Failure 500 {object} route.HttpError
//@Router /v1/images/{pathId} [put]
func (s *service) UpdateImage(ctx *https.Context) {
	body, err := s.body(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, s.NewHttpError(ImageCode, ImageMsg, RequestBodyErrorCode, RequestBodyErrorMsg))
		return
	}
	imageId, err := s.pathId(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, s.NewHttpError(ImageCode, ImageMsg, IDErrorCode, IDErrorMsg))
		return
	}
	param := RequestImage{}
	err = json.Unmarshal(body, &param)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, s.NewHttpError(ImageCode, ImageMsg, BodyStructureErrorCode, BodyStructureErrorMsg))
		return
	}
	err = config.Services().ImageService().UpdateImage(imageId, &param.Description, param.Status)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, s.NewHttpError(ImageCode, ImageMsg, UpdateErrorCode, UpdateErrorMsg))
		return
	}
	ctx.JSON(http.StatusOK, s.NewHttpOk(http.StatusOK))
}

//@Summary 删除图片
//@Produce json
//@Param language header string false "语言" default "chinese"
//@Success 200 {object} route.HttpOk
//@Failure 400 {object} route.HttpError
//@Failure 404 {object} route.HttpError
//@Failure 500 {object} route.HttpError
//@Router /v1/images/{pathId} [delete]
func (s *service) DeleteImage(ctx *https.Context) {
	imageId, err := s.pathId(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, s.NewHttpError(ImageCode, ImageMsg, IDErrorCode, IDErrorMsg))
		return
	}
	err = config.Services().ImageService().DeleteImage(imageId)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, s.NewHttpError(ImageCode, ImageMsg, DeleteErrorCode, DeleteErrorMsg))
		return
	}
	ctx.JSON(http.StatusOK, s.NewHttpOk(http.StatusOK))
}
