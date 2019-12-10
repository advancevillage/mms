package init

import (
	"github.com/advancevillage/3rd/caches"
	"github.com/advancevillage/3rd/logs"
	"github.com/advancevillage/3rd/storages"
)
//初始化地址和端口
//初始化存储配置
//初始化缓存配置
//初始化日志配置
type MMS struct {
	Host  string 	`xml:"host"`
	Port  int    	`xml:"port"`
	cache  	   caches.Cache		`xml:"-"`
	logger 	   logs.Logs 		`xml:"-"`
	storage    storages.Storage `xml:"-"`
}




