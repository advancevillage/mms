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
	"mms/src/component/size"
	"mms/src/component/tag"
	"os"
)

func GetMMSObject() *MMS {
	return &defaultMMS
}

func LoadArgs(commit, version, buildTime string) error {
	var args = os.Args
	var length = len(args)
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
	defaultMMS.logger, err = logs.NewTxtLogger(defaultConfigure.Log, 4096, 4)
	if err != nil {
		return err
	}
	defaultMMS.es7, err = storages.NewTES(defaultConfigure.Es7.DSN, defaultMMS.logger)
	if err != nil {
		return err
	}
	defaultMMS.cache, err = caches.NewRedis(defaultConfigure.Redis.Host, defaultConfigure.Redis.Port, defaultConfigure.Redis.Auth, defaultConfigure.Redis.Schema, defaultMMS.logger, defaultMMS.es7)
	if err != nil {
		return err
	}
	defaultMMS.manufacturer = manufacturer.NewManufacturerService(defaultMMS.es7, defaultMMS.logger)
	defaultMMS.category = category.NewCategoryService(defaultMMS.es7, defaultMMS.logger)
	defaultMMS.brand    = brand.NewBrandService(defaultMMS.es7, defaultMMS.logger)
	defaultMMS.tag      = tag.NewTagService(defaultMMS.es7, defaultMMS.logger)
	defaultMMS.color    = color.NewColorService(defaultMMS.es7, defaultMMS.logger)
	defaultMMS.image    = image.NewImageService(defaultMMS.es7, defaultMMS.logger)
	defaultMMS.goods    = goods.NewGoodsService(defaultMMS.es7, defaultMMS.logger)
	defaultMMS.size     = size.NewSizeService(defaultMMS.es7, defaultMMS.logger)
	return nil
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
	return fmt.Sprintf("commit=%s version=%s buildTime=%s", defaultConfigure.commit, defaultConfigure.version, defaultConfigure.buildTime)
}

func (mms *MMS) GetLogger() logs.Logs {
	return mms.logger
}