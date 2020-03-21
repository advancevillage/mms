//author: richard
package config

import (
	"encoding/xml"
	"fmt"
	"github.com/advancevillage/3rd/caches"
	"github.com/advancevillage/3rd/files"
	"github.com/advancevillage/3rd/logs"
	"github.com/advancevillage/3rd/pay"
	"github.com/advancevillage/3rd/storages"
	"os"
)


func NewService() *Service {
	return &Service{}
}

func (s *Service) ExitWithInfo(format string, a ...interface{}) {
	fmt.Printf(format + "\n\n\t", a ...)
	os.Exit(0)
}

func (s *Service) LoadArgs(commit, buildTime string) error {
	var args = os.Args
	var length = len(args)
	s.Configure.Commit	   = commit
	s.Configure.BuildTime  = buildTime
	s.Configure.File 	= "./etc/oms.xml"
	for i := 0; i < length; i += 2 {
		switch args[i] {
		case "--config", "-c":
			if j := i+1; j < length {
				s.Configure.File = args[j]
			}
		case "--version", "-v":
			s.ExitWithInfo("commit=%s, buildTime=%s", commit, buildTime)
		default:
			continue
		}
	}
	return nil
}

func (s *Service) LoadConfigure() error {
	buf, err := files.NewXMLFile().ReadFile(s.Configure.File)
	if err != nil {
		return err
	}
	err = xml.Unmarshal(buf, &s.Configure)
	if err != nil {
		return err
	}
	return nil
}

func (s *Service) LoadServices() error {
	var err error
	s.Logger, err = logs.NewTxtLogger(s.Configure.Log, 1024, 4)
	if err != nil {
		s.ExitWithInfo("init logger fail")
		return err
	}
	//@mongo
	s.Mongo, err = storages.NewMongoDB(s.Configure.Mongo, s.Logger)
	if err != nil {
		s.ExitWithInfo("init mongo fail")
		return err
	}
	//@cache
	s.Cache, err = caches.NewRedisStorage(s.Configure.Redis.Host, s.Configure.Redis.Port, s.Configure.Redis.Auth, s.Configure.Redis.Schema, s.Logger)
	if err != nil {
		s.ExitWithInfo("init redis cache fail")
		return err
	}
	//@pay
	s.Pay = pay.NewBrainTree(s.Configure.Braintree.Url, s.Configure.Braintree.Merchant, s.Configure.Braintree.Public, s.Configure.Braintree.Private, s.Logger)
	return nil
}