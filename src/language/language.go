//author: richard
package language

import (
	"github.com/advancevillage/3rd/logs"
	"github.com/advancevillage/3rd/translate"
	"strings"
)

type Languages struct {
	English string `json:"english"`
	Chinese string `json:"chinese"`
}

func (s *Languages) Multi(lang string, translate translate.Translate, logger logs.Logs) {
	var err error
	switch strings.ToLower(lang) {
	case "english":
		s.Chinese, err = translate.Translate(s.English, "en", "zh")
	case "chinese":
		s.English, err = translate.Translate(s.Chinese, "zh", "en")
	}
	if err != nil {
		logger.Warning(err.Error())
	}
}