package route

import "github.com/advancevillage/3rd/https"


var router = func (api API) []https.Router{
	return []https.Router {
		{"GET", "/v1/merchandises/version", api.Version},
	}
}

type API interface {
	//merchandises
	Version(ctx *https.Context)
}

func LoadRouter(host string, port int) error {
	server := https.NewServer(host, port, router(NewApiService()))
	err := server.StartServer()
	if err != nil {
		return err
	}
	return nil
}