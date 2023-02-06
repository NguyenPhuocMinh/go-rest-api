package commons

import (
	translator "fast-food-api-client/translator"
	utils "fast-food-api-client/utils"

	"github.com/gin-gonic/gin"
)

func TemplateSuccessCommon(c *gin.Context, data interface{}, messageCode string) *ResponseModel {

	lang := c.Request.Header.Get("Lang")

	statusCode, messageName, messageTranslator := buildTemplateCommon(messageCode, lang)

	res := ResponseModel{
		Data:              data,
		Error:             nil,
		StatusCode:        statusCode,
		MessageCode:       messageCode,
		MessageName:       messageName,
		MessageTranslator: messageTranslator,
	}

	return &res
}

func TemplateErrorCommon(c *gin.Context, err error, messageCode string) *ResponseModel {

	lang := c.Request.Header.Get("Lang")

	statusCode, messageName, messageTranslator := buildTemplateCommon(messageCode, lang)

	res := ResponseModel{
		Data:              nil,
		Error:             err.Error(),
		StatusCode:        statusCode,
		MessageCode:       messageCode,
		MessageName:       messageName,
		MessageTranslator: messageTranslator,
		Pagination:        nil,
	}

	return &res
}

func buildTemplateCommon(messageCode string, lang string) (statusCode int, messageName string, messageTranslator string) {
	statusCode = utils.MapStatusByMsgCode(messageCode)
	messageName = utils.MapMsgNameByMsgCode(messageCode)
	messageTranslator = translator.TranslateMessage(lang, messageCode)

	return statusCode, messageName, messageTranslator
}
