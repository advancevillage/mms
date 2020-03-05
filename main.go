//author: richard
// @title mms
// @version 0.0.1
// @description 商品中心
// @contact.name richard sun
// @contact.email cugriver@163.com
// @license.name Apache 2.0
// @license.url  http://www.apache.org/licenses/LICENSE-2.0.html
// @host localhost:13147
// @BasePath /
// @schemes http https
package main

import (
	"mms/brand"
	"mms/category"
	"mms/color"
	"mms/config"
	"mms/goods"
	"mms/internal"
	"mms/language"
	"mms/manufacturer"
	"mms/route"
	"mms/size"
	"sync"
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

	//init goods service
	goodsService := goods.NewService(configService.Mongo, configService.Logger)

	//init color service
	colorService := color.NewService(configService.Mongo, configService.Logger)

	//init size service
	sizeService  := size.NewService(configService.Mongo, configService.Logger)

	//init brand service
	brandService := brand.NewService(configService.Mongo, configService.Logger)

	//init category service
	categoryService := category.NewService(configService.Mongo, configService.Logger)

	//init manufacturer service
	manufacturerService := manufacturer.NewService(configService.Mongo, configService.Logger)

	//init http service
	routeService := route.NewService(configService, langService, goodsService, colorService, sizeService, brandService, categoryService, manufacturerService)

	//init rcp service
	rpcService := internal.NewService(configService, langService, goodsService, colorService, sizeService, brandService, categoryService, manufacturerService)

	wg := sync.WaitGroup{}
	wg.Add(2)
	go func() {
		//start http server
		defer wg.Done()
		err = routeService.StartHttpServer(mode)
	}()
	go func() {
		defer wg.Done()
		err = rpcService.StartRPCServer()
	}()
	wg.Wait()
	return
}
