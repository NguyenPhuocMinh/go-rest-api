package v1

import (
	"encoding/json"
	"fmt"

	commons "fast-food-api-client/commons"
	constants "fast-food-api-client/constants"
	database "fast-food-api-client/core/database"
	coreLogger "fast-food-api-client/core/logger"
	resources "fast-food-api-client/resources"
	models "fast-food-api-client/src/models"
	utils "fast-food-api-client/utils"

	"github.com/gin-gonic/gin"
)

var logger = coreLogger.Logger(constants.AppName, constants.StructAuthService)

func Test(c *gin.Context) *commons.ResponseModel {
	logger.Info("[BEGIN] Test Service")

	var result *commons.ResponseModel
	var customer models.CustomerModel
	var err error

	err = commons.ValidatorRequestBodyCommon(c, customer)

	if err != nil {
		result = commons.TemplateErrorCommon(c, err, resources.MsgCodeRequestDataInvalid)
		logger.Error("Validate Request Failed With Err = ", "["+err.Error()+"]")
		return result
	}

	createdAt, updatedAt := utils.AuditAttribute("create")

	customer = models.CustomerModel{
		CreatedAt: createdAt,
		UpdatedAt: updatedAt,
	}

	err = commons.ValidatorStructCommon(&customer)
	if err != nil {
		result = commons.TemplateErrorCommon(c, err, resources.MsgCodeRequestDataInvalid)
		logger.Error("Validate Struct Failed With Err = ", "["+err.Error()+"]")
		return result
	}

	data, err := database.CreateOne(constants.CustomerModel, customer)
	if err != nil {
		logger.Error("Failed With Err=", err.Error())
		result = commons.TemplateErrorCommon(c, err, resources.MsgCodeRequestDatabaseError)
		return result
	}

	a, _ := json.MarshalIndent(data, "", " ")
	fmt.Println("hehe", string(a))

	result = commons.TemplateSuccessCommon(c, data, resources.MsgCodeRequestDataIsValid)
	logger.Info("[END] Test Service")
	return result
}
