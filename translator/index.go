package translator

import (
	"encoding/json"
	constants "fast-food-api-client/constants"
	coreLogger "fast-food-api-client/core/logger"
	helpers "fast-food-api-client/helpers"
	utils "fast-food-api-client/utils"

	"github.com/nicksnyder/go-i18n/v2/i18n"
	"golang.org/x/text/language"
)

var logTranslator = coreLogger.Logger(constants.AppName, constants.StructTranslator)

func TranslateMessage(lang string, msgCode string) string {
	logTranslator.Info("BEGIN TranslateMessage....")
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

	logTranslator.Info("END TranslateMessage with translation=", "["+translation+"]")

	return translation
}
