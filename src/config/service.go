//author: richard
package config

import (
	"encoding/xml"
	"fmt"
	"github.com/advancevillage/3rd/files"
	"os"
)

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
			fmt.Printf("commit=%s, version=%s, buildTime=%s", commit, version, buildTime)
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