package commons

import (
	"encoding/json"

	constants "fast-food-api-client/constants"
	helpers "fast-food-api-client/helpers"
	utils "fast-food-api-client/utils"

	"github.com/nicksnyder/go-i18n/v2/i18n"
	"golang.org/x/text/language"
)

func TranslatorMsgCommon(lang string, msgCode string) string {
	bundle := i18n.NewBundle(language.English)
	bundle.RegisterUnmarshalFunc("json", json.Unmarshal)
	bundle.MustLoadMessageFile("resources/languages/en.json")
	bundle.MustLoadMessageFile("resources/languages/vi.json")

	if helpers.IsNil(lang) {
		lang = constants.LangEN
	}

	// locales
	localizer := i18n.NewLocalizer(bundle, lang)

	// map message name
	msgName := utils.MapMsgNameByMsgCode(msgCode)

	translation := localizer.MustLocalize(&i18n.LocalizeConfig{
		MessageID: msgName,
	})

	return translation
}
