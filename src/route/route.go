package route

import (
	"github.com/advancevillage/3rd/https"
)


var router = func (api API) []https.Router{
	return []https.Router {
		//商品路由
		{"GET", "/v1/merchandises", api.QueryMerchandises},
		{"POST", "/v1/merchandises", api.CreateMerchandise},
		{"PUT", "/v1/merchandises/:goodsId", api.UpdateMerchandise},
		{"GET", "/v1/merchandises/:goodsId", api.QueryMerchandise},
		{"DELETE", "/v1/merchandises/:goodsId", api.DeleteMerchandise},
		//分类路由
		{"POST", "/v1/categories", api.CreateCategory},
	}
}

type API interface {
	//merchandises
	Version(ctx *https.Context)
	QueryMerchandises(ctx *https.Context)
	CreateMerchandise(ctx *https.Context)
	QueryMerchandise (ctx *https.Context)
	UpdateMerchandise(ctx *https.Context)
	DeleteMerchandise(ctx *https.Context)
	//categories
	CreateCategory(ctx *https.Context)
}

func LoadRouter(host string, port int, mode string) error {
	var err error
	switch mode {
	case "lambda":
		err = LoadLambdaRouter()
	default:
		err =  LoadHttpRouter(host, port)
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