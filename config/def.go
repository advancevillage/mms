//author: richard
package config

import (
	"encoding/xml"
	"github.com/advancevillage/3rd/caches"
	"github.com/advancevillage/3rd/logs"
	"github.com/advancevillage/3rd/storages"
)

type configure struct {
	XMLName 	xml.Name `xml:"mms"`
	HttpHost 	string   `xml:"httpHost"`
	HttpPort	int 	 `xml:"httpPort"`
	Redis struct{
		Host string `xml:"host"`
		Port int 	`xml:"port"`
		Auth string `xml:"auth"`
		Schema int  `xml:"schema"`
	} `xml:"redis"`
	Mongo string	`xml:"mongo"`
	Log 		string   `xml:"log"`
	Upload 		string 	 `xml:"upload"`
	File 	  	string 	 `xml:"-"`
	Commit 		string   `xml:"-"`
	BuildTime 	string   `xml:"-"`
}

type 	Service struct {
	Cache     caches.ICache
	Logger    logs.Logs
	Mongo     storages.Storage
	Configure configure
}
