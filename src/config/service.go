//author: richard
package config

import (
	"encoding/xml"
	"fmt"
	"github.com/advancevillage/3rd/files"
	"github.com/advancevillage/3rd/logs"
	"github.com/advancevillage/3rd/storages"
	"mms/src/component/brand"
	"mms/src/component/category"
	"mms/src/component/color"
	"mms/src/component/goods"
	"mms/src/component/image"
	"mms/src/component/manufacturer"
	"mms/src/component/size"
	"mms/src/component/tag"
	"os"
)

func GetMMSObject() *MMS {
	return &defaultMMS
}

func LoadArgs(commit, version, buildTime string, mode string) error {
	var args = os.Args
	var length = len(args)
	defaultConfigure.mode 	   = mode
	defaultConfigure.commit    = commit
	defaultConfigure.version   = version
	defaultConfigure.buildTime = buildTime
	defaultConfigure.File = "./etc/mms.xml"
	for i := 0; i < length; i += 2 {
		switch args[i] {
		case "--config", "-c":
			if j := i+1; j < length {
				defaultConfigure.File = args[j]
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
	buf, err := files.NewXMLFile().ReadFile(defaultConfigure.File)
	if err != nil {
		return err
	}
	err = xml.Unmarshal(buf, &defaultConfigure)
	if err != nil {
		return err
	}
	return nil
}

func LoadObject() error {
	var err error
	defaultMMS.logger, err = logs.NewTxtLogger(defaultConfigure.Log, 1024, 4)
	if err != nil {
		return err
	}
	//@es7
	//defaultMMS.es7, err = storages.NewTES(defaultConfigure.Es7.DSN, defaultMMS.logger)
	//if err != nil {
	//	return err
	//}
	//@cache
	//defaultMMS.cache, err = caches.NewRedis(defaultConfigure.Redis.Host, defaultConfigure.Redis.Port, defaultConfigure.Redis.Auth, defaultConfigure.Redis.Schema, defaultMMS.logger, defaultMMS.es7)
	//if err != nil {
	//	return err
	//}
	//@mongo
	defaultMMS.mgo, err = storages.NewMongoDB(defaultConfigure.Mongo, defaultMMS.logger)
	if err != nil {
		return err
	}
	defaultMMS.manufacturer = manufacturer.NewManufacturerService(defaultMMS.mgo, defaultMMS.logger)
	defaultMMS.category = category.NewCategoryService(defaultMMS.mgo, defaultMMS.logger)
	defaultMMS.brand    = brand.NewBrandService(defaultMMS.mgo, defaultMMS.logger)
	defaultMMS.tag      = tag.NewTagService(defaultMMS.mgo, defaultMMS.logger)
	defaultMMS.color    = color.NewColorService(defaultMMS.mgo, defaultMMS.logger)
	defaultMMS.image    = image.NewImageService(defaultMMS.mgo, defaultMMS.logger)
	defaultMMS.goods    = goods.NewGoodsService(defaultMMS.mgo, defaultMMS.logger)
	defaultMMS.size     = size.NewSizeService(defaultMMS.mgo, defaultMMS.logger)
	return err
}

func ExitWithInfo(format string, a ...interface{}) {
	fmt.Printf(format + "\n\n\t", a ...)
	os.Exit(0)
}

func (mms *MMS) GetHttpHost() string {
	return defaultConfigure.HttpHost
}

func (mms *MMS) GetHttpPort() int {
	return defaultConfigure.HttpPort
}

func (mms *MMS) GetCategoryService() *category.Service {
	return mms.category
}

func (mms *MMS) GetVersion() string {
	return fmt.Sprintf("commit=%s version=%s buildTime=%s mode=%s", defaultConfigure.commit, defaultConfigure.version, defaultConfigure.buildTime, defaultConfigure.mode)
}

func (mms *MMS) GetMode() string {
	return defaultConfigure.mode
}

func (mms *MMS) GetLogger() logs.Logs {
	return mms.logger
}