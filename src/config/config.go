package config

import (
	"encoding/xml"
	"github.com/advancevillage/3rd/caches"
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
	"mms/src/component/style"
	"mms/src/component/tag"
)

var service  Service
var configure Configure
//初始化地址和端口
//初始化存储配置
//初始化缓存配置
//初始化日志配置
type Service struct {
	cache  	   caches.ICache
	logger 	   logs.Logs
	mgo 	   storages.Storage
	category   *category.Service
	brand 	   *brand.Service
	tag 	   *tag.Service
	color 	   *color.Service
	image 	   *image.Service
	manufacturer *manufacturer.Service
	size 	   *size.Service
	goods 	   *goods.Service
	session    *session.Service
	style      *style.Service
}

type Configure struct {
	XMLName 	xml.Name `xml:"mms"`
	HttpHost 	string   `xml:"httpHost"`
	HttpPort	int 	 `xml:"httpPort"`
	Es7   struct{
		DSN []string	 `xml:"dsn"`
	}	`xml:"es7"`
	Redis struct{
		Host string `xml:"host"`
		Port int 	`xml:"port"`
		Auth string `xml:"auth"`
		Schema int `xml:"schema"`
	}	`xml:"redis"`
	Mongo string	`xml:"mongo"`
	Log 		string   `xml:"log"`
	File 	  	string 	 `xml:"-"`
	commit 		string   `xml:"-"`
	version 	string 	 `xml:"-"`
	buildTime 	string   `xml:"-"`
	mode 		string 	 `xml:"-"`
}

