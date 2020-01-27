//author: richard
package language

import (
	"github.com/advancevillage/3rd/translate"
	"mms/api"
	"strings"
)

func NewService() *Service {
	t := translate.NewBaiDuTranslate("20200116000375924", "9T6lwR7uIXbvef_O7Wd3")
	return &Service{t:t}
}

func (s *Service) I18n(value *api.Languages, lang string) error {
	var err error
	switch strings.ToLower(lang) {
	case "english":
		value.Chinese, err = s.t.Translate(value.English, "en", "zh")
	case "chinese":
		value.English, err = s.t.Translate(value.Chinese, "zh", "en")
	}
	return err
}
