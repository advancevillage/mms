package route

import (
	"github.com/advancevillage/3rd/https"
)


var router = func (api API) []https.Router{
	return []https.Router {
		//商品路由
		{"GET", "/v1/merchandises", api.QueryMerchandises},
		{"POST", "/v1/merchandises", api.CreateMerchandise},
		{"GET", "/v1/merchandises/:pathId", api.QueryMerchandise},
		{"DELETE", "/v1/merchandises/:pathId", api.DeleteMerchandise},
		//分类路由
		{"GET", "/v1/categories", api.QueryCategories},
		{"POST", "/v1/categories", api.CreateCategory},
		{"PUT", "/v1/categories/:pathId", api.UpdateCategory},
		{"DELETE", "/v1/categories/:pathId", api.DeleteCategory},
		{"GET", "/v1/categories/:pathId", api.QueryCategory},
		//品牌路由
		{"GET", "/v1/brands", api.QueryBrands},
		{"POST", "/v1/brands", api.CreateBrand},
		{"PUT", "/v1/brands/:pathId", api.UpdateBrand},
		{"DELETE", "/v1/brands/:pathId", api.DeleteBrand},
		{"GET", "/v1/brands/:pathId", api.QueryBrand},
		//尺码路由
		{"POST", "/v1/sizes", api.CreateSize},
		{"GET", "/v1/sizes", api.QuerySizes},
		{"GET", "/v1/sizes/:pathId", api.QuerySize},
		{"PUT", "/v1/sizes/:pathId", api.UpdateSize},
		{"DELETE", "/v1/sizes/:pathId", api.DeleteSize},
		//款式路由
		{"POST", "/v1/styles", api.CreateStyle},
		{"GET", "/v1/styles", api.QueryStyles},
		{"GET", "/v1/styles/:pathId", api.QueryStyle},
		{"PUT", "/v1/styles/:pathId", api.UpdateStyle},
		{"DELETE", "/v1/styles/:pathId", api.DeleteStyle},
		//标签路由
		{"POST", "/v1/tags", api.CreateTag},
		{"GET", "/v1/tags", api.QueryTags},
		{"GET", "/v1/tags/:pathId", api.QueryTag},
		{"PUT", "/v1/tags/:pathId", api.UpdateTag},
		{"DELETE", "/v1/tags/:pathId", api.DeleteTag},
		//颜色路由
		{"POST", "/v1/colors", api.CreateColor},
		{"GET", "/v1/colors", api.QueryColors},
		{"GET", "/v1/colors/:pathId", api.QueryColor},
		{"PUT", "/v1/colors/:pathId", api.UpdateColor},
		{"DELETE", "/v1/colors/:pathId", api.DeleteColor},
		//图片路由
		{"POST", "/v1/images", api.CreateImage},
		{"GET", "/v1/images", api.QueryImages},
		{"GET", "/v1/images/:pathId", api.QueryImage},
		{"PUT", "/v1/images/:pathId", api.UpdateImage},
		{"DELETE", "/v1/images/:pathId", api.DeleteImage},
		//生产商路由 manufacturers
		{"POST", "/v1/manufacturers", api.CreateManufacturer},
		{"GET", "/v1/manufacturers", api.QueryManufacturers},
		{"GET", "/v1/manufacturers/:pathId", api.QueryManufacturer},
		{"PUT", "/v1/manufacturers/:pathId", api.UpdateManufacturer},
		{"DELETE", "/v1/manufacturers/:pathId", api.DeleteManufacturer},
		//路由服务
		{"GET", "/v1/services/version", api.Version},
	}
}

type API interface {
	//services
	Version(ctx *https.Context)
	//merchandises
	QueryMerchandises(ctx *https.Context)
	CreateMerchandise(ctx *https.Context)
	QueryMerchandise (ctx *https.Context)
	//UpdateMerchandise(ctx *https.Context)
	DeleteMerchandise(ctx *https.Context)
	//categories
	QueryCategories(ctx *https.Context)
	CreateCategory(ctx *https.Context)
	UpdateCategory(ctx *https.Context)
	DeleteCategory(ctx *https.Context)
	QueryCategory (ctx *https.Context)
	//brands
	QueryBrands(ctx *https.Context)
	CreateBrand(ctx *https.Context)
	UpdateBrand(ctx *https.Context)
	DeleteBrand(ctx *https.Context)
	QueryBrand (ctx *https.Context)
	//sizes
	CreateSize(ctx *https.Context)
	QuerySizes(ctx *https.Context)
	QuerySize (ctx *https.Context)
	UpdateSize(ctx *https.Context)
	DeleteSize(ctx *https.Context)
	//styles
	CreateStyle(ctx *https.Context)
	QueryStyles(ctx *https.Context)
	QueryStyle (ctx *https.Context)
	UpdateStyle(ctx *https.Context)
	DeleteStyle(ctx *https.Context)
	//tags
	CreateTag(ctx *https.Context)
	QueryTags(ctx *https.Context)
	QueryTag (ctx *https.Context)
	UpdateTag(ctx *https.Context)
	DeleteTag(ctx *https.Context)
	//colors
	CreateColor(ctx *https.Context)
	QueryColors(ctx *https.Context)
	QueryColor (ctx *https.Context)
	UpdateColor(ctx *https.Context)
	DeleteColor(ctx *https.Context)
	//images
	CreateImage(ctx *https.Context)
	QueryImages(ctx *https.Context)
	QueryImage (ctx *https.Context)
	UpdateImage(ctx *https.Context)
	DeleteImage(ctx *https.Context)
	//manufacturers
	CreateManufacturer(ctx *https.Context)
	QueryManufacturers(ctx *https.Context)
	QueryManufacturer (ctx *https.Context)
	UpdateManufacturer(ctx *https.Context)
	DeleteManufacturer(ctx *https.Context)
}

func LoadRouter(host string, port int, mode string) error {
	var err error
	switch mode {
	case "lambda":
		err = LoadLambdaRouter()
	default:
		err = LoadHttpRouter(host, port)
	}
	return err
}

func LoadHttpRouter(host string, port int) error {
	server := https.NewServer(host, port, router(NewApiService()))
	err := server.StartServer()
	if err != nil {
		return err
	}
	return nil
}

func LoadLambdaRouter() error {
	server := https.NewAwsApiGatewayLambdaServer(router(NewApiService()))
	err := server.StartServer()
	if err != nil {
		return err
	}
	return nil
}