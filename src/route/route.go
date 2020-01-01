package route

import (
	"github.com/advancevillage/3rd/https"
	"mms/src/config"
)


var router = func (api API) []https.Router{
	return []https.Router {
		{"GET", "/v1/merchandises/version", api.Version},
		{"POST", "/v1/categories", api.CreateCategory},
	}
}

type API interface {
	//merchandises
	Version(ctx *https.Context)
	CreateCategory(ctx *https.Context)
}

func LoadRouter(host string, port int, mode string) error {
	var err error
	switch mode {
	case "lambda":
		config.GetMMSObject().GetLogger().Info("%s", "loadLambdaRouter")
		err = LoadLambdaRouter()
	default:
		config.GetMMSObject().GetLogger().Info("%s", "loadHttpRouter")
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