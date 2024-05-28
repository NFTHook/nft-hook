package ml

import (
	"encoding/json"
	"nfthook/config"

	"github.com/nicksnyder/go-i18n/v2/i18n"
	"golang.org/x/text/language"
)

var bundle *i18n.Bundle

func init() {
	bundle = i18n.NewBundle(language.English)
	bundle.RegisterUnmarshalFunc("json", json.Unmarshal)
	bundle.LoadMessageFile(config.Get().Env.En_i18n)
	bundle.LoadMessageFile(config.Get().Env.Zh_cn_i18n)
}

func getTranslatedText(code, lang string) string {

	if len(lang) == 0 {
		lang = "zh_cn"
	}
	localizer := i18n.NewLocalizer(bundle, lang)
	message, err := localizer.Localize(&i18n.LocalizeConfig{MessageID: code})
	if err != nil {
		return ""
	}
	return message
}

type ResponseData struct {
	Code    string      `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func Succ(lang string, data interface{}) ResponseData {
	return ResponseData{
		Code:    "200",
		Message: getTranslatedText("200", lang),
		Data:    data,
	}
}

func Fail(lang string, code string) ResponseData {
	return ResponseData{
		Code:    code,
		Message: getTranslatedText(code, lang),
		Data:    nil,
	}
}

func Res(lang string, code string, data interface{}) ResponseData {
	return ResponseData{
		Code:    code,
		Message: getTranslatedText(code, lang),
		Data:    data,
	}
}
