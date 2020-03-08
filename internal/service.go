//author: richard
package internal //json-rpc 1.0
import (
	"github.com/advancevillage/3rd/rpcs"
	"mms/api"
	"mms/brand"
	"mms/category"
	"mms/color"
	"mms/config"
	"mms/goods"
	"mms/language"
	"mms/manufacturer"
	"mms/size"
)

func NewService(configService *config.Service, langService *language.Service, goodsService *goods.Service, colorService *color.Service, sizeService *size.Service, brandService *brand.Service, categoryService *category.Service, manufacturerService *manufacturer.Service) *Service {
	return &Service{
		configService: configService,
		langService:   langService,
		goodsService: goodsService,
		colorService: colorService,
		sizeService: sizeService,
		brandService: brandService,
		categoryService: categoryService,
		manufacturerService: manufacturerService,
	}
}

func (s *Service) StartRPCServer() error {
	rcvr := make([]interface{}, 0)
	rcvr = append(rcvr, s)
	server := rpcs.NewServer(s.configService.Configure.RpcHost, s.configService.Configure.RpcPort, s.configService.Logger, rcvr)
	err := server.StartServer()
	if err != nil {
		return err
	}
	return nil
}

func (s *Service) UpdateStock(args *api.Stocks, result *api.Stocks) error {
	return s.goodsService.IncreaseStock(args)
}





