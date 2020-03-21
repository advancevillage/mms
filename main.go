//author: richard
// @title ums
// @version 0.0.1
// @description 用户中心
// @contact.name richard sun
// @contact.email cugriver@163.com
// @license.name Apache 2.0
// @license.url  http://www.apache.org/licenses/LICENSE-2.0.html
// @host localhost:13172
// @BasePath /
// @schemes http https
package main

import (
	"mms/config"
	"mms/language"
	"mms/order"
	"mms/route"
	"mms/session"
)

var (
	mode    = "http"
	commit  = "000000"
	buildTime = "2006-01-03 16:05:06"
)

func main() {
	//init configure service
	configService := config.NewService()
	//1: load args
	err := configService.LoadArgs(commit, buildTime)
	if err != nil {
		configService.ExitWithInfo("load args fail")
		return
	}
	//2: load configure
	err = configService.LoadConfigure()
	if err != nil {
		configService.ExitWithInfo("load configure fail")
		return
	}
	//3: init basic service
	err = configService.LoadServices()
	if err != nil {
		configService.ExitWithInfo("load basic service fail")
		return
	}

	//init lang service
	langService := language.NewService()

	//init order service
	orderService := order.NewService(configService.Mongo, configService.Pay, configService.Logger)

	//init session service
	sessionService := session.NewService(configService.Cache, configService.Logger)

	//init route service
	routeService := route.NewService(configService, langService, orderService, sessionService)

	//start route service
	err = routeService.StartRouter(mode)
	if err != nil {
		configService.ExitWithInfo("start route fail: %s", err.Error())
		return
	}
	return
}
