package config

import (
	"encoding/xml"
	"github.com/advancevillage/3rd/caches"
	"github.com/advancevillage/3rd/logs"
	"github.com/advancevillage/3rd/storages"
	"mms/src/component/category"
)

var defaultConfigure Configure
//初始化地址和端口
//初始化存储配置
//初始化缓存配置
//初始化日志配置
type MMS struct {
	host  string
	port  int
	cache  	   caches.Cache
	logger 	   logs.Logs
	es7    	   storages.Storage
	category   *category.Service
}

type Configure struct {
	XMLName 	xml.Name `xml:"mms"`
	HttpHost 	string   `xml:"httpHost"`
	HttpPort	int 	 `xml:"httpPort"`
	Es7   struct{
		DSN []string	 `xml:"dsn"`
	}	`xml:"es7"`
	File 	  	string 	 `xml:"-"`
	commit 		string   `xml:"-"`
	version 	string 	 `xml:"-"`
	buildTime 	string   `xml:"-"`
}




