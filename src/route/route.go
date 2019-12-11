package route

import "github.com/advancevillage/3rd/https"

var api = []https.Router{
	{"GET", "/v1/merchandises", nil},
}

func InitRoute(host string, port int)  {
	server := https.NewServer(host, port, api)
	err := server.StartServer()
	if err != nil {

	}
}