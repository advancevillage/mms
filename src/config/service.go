//author: richard
package config

import (
	"encoding/xml"
	"fmt"
	"github.com/advancevillage/3rd/caches"
	"github.com/advancevillage/3rd/files"
	"github.com/advancevillage/3rd/logs"
	"github.com/advancevillage/3rd/storages"
	"mms/src/component/brand"
	"mms/src/component/category"
	"mms/src/component/color"
	"mms/src/component/goods"
	"mms/src/component/image"
	"mms/src/component/manufacturer"
	"mms/src/component/session"
	"mms/src/component/size"
	"mms/src/component/tag"
	"os"
)

func Services() *Service {
	return &service
}

func LoadArgs(commit, version, buildTime string, mode string) error {
	var args = os.Args
	var length = len(args)
	configure.mode 	    = mode
	configure.commit    = commit
	configure.version   = version
	configure.buildTime = buildTime
	configure.File = "./etc/mms.xml"
	for i := 0; i < length; i += 2 {
		switch args[i] {
		case "--config", "-c":
			if j := i+1; j < length {
				configure.File = args[j]
			}
		case "--version", "-v":
			ExitWithInfo("commit=%s, version=%s, buildTime=%s", commit, version, buildTime)
			os.Exit(0)
		default:
			continue
		}
	}
	return nil
}

func LoadConfigure() error {
	buf, err := files.NewXMLFile().ReadFile(configure.File)
	if err != nil {
		return err
	}
	err = xml.Unmarshal(buf, &configure)
	if err != nil {
		return err
	}
	return nil
}

func LoadServices() error {
	var err error
	service.logger, err = logs.NewTxtLogger(configure.Log, 1024, 4)
	if err != nil {
		return err
	}
	//@mongo
	service.mgo, err = storages.NewMongoDB(configure.Mongo, service.logger)
	if err != nil {
		return err
	}
	//@cache
	service.cache, err = caches.NewRedisStorage(configure.Redis.Host, configure.Redis.Port, configure.Redis.Auth, configure.Redis.Schema, service.logger)
	if err != nil {
		return err
	}
	service.manufacturer = manufacturer.NewManufacturerService(service.mgo, service.logger)
	service.category = category.NewCategoryService(service.mgo, service.logger)
	service.brand    = brand.NewBrandService(service.mgo, service.logger)
	service.tag      = tag.NewTagService(service.mgo, service.logger)
	service.color    = color.NewColorService(service.mgo, service.logger)
	service.image    = image.NewImageService(service.mgo, service.logger)
	service.goods    = goods.NewGoodsService(service.mgo, service.logger)
	service.size     = size.NewSizeService(service.mgo, service.logger)
	service.session  = session.NewSessionService(service.cache, service.logger)
	return err
}

func ExitWithInfo(format string, a ...interface{}) {
	fmt.Printf(format + "\n\n\t", a ...)
	os.Exit(0)
}

func (s *Service) GetHttpHost() string {
	return configure.HttpHost
}

func (s *Service) GetHttpPort() int {
	return configure.HttpPort
}

func (s *Service) GetMode() string {
	return configure.mode
}

func (s *Service) Version() string {
	return fmt.Sprintf("commit=%s version=%s buildTime=%s mode=%s", configure.commit, configure.version, configure.buildTime, configure.mode)
}

func (s *Service) LogService() logs.Logs {
	return s.logger
}

func (s *Service) CategoryService() *category.Service {
	return s.category
}