package v1

import (
	"encoding/json"
	"fmt"

	commons "fast-food-api-client/commons"
	constants "fast-food-api-client/constants"
	database "fast-food-api-client/core/database"
	coreLogger "fast-food-api-client/core/logger"
	resources "fast-food-api-client/resources"
	mappers "fast-food-api-client/src/mappers"
	models "fast-food-api-client/src/models"
	utils "fast-food-api-client/utils"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)

var logAuth = coreLogger.Logger(constants.AppName, constants.StructAuthService)

var cusMapper = mappers.CustomerMapper{}

func AuthLogin(c *gin.Context) *commons.ResponseModel {
	logAuth.Info("[BEGIN] AuthLogin Service")
	// customer := &models.CustomerModel{
	// 	Name: "Test",
	// }
	// data, err := database.CreateOne("CustomerModel", customer)

	res := commons.TemplateSuccessCommon(c, "hehe", resources.MsgCodeRequestDataIsValid)
	logAuth.Info("[END] AuthLogin Service...")
	return res
}

func AuthRegister(c *gin.Context) *commons.ResponseModel {
	logAuth.Info("[BEGIN] AuthRegister Service...")

	var result *commons.ResponseModel
	var customer models.CustomerModel
	var err error

	// validate request
	err = commons.ValidatorRequestBodyCommon(c, &customer)
	if err != nil {
		result = commons.TemplateErrorCommon(c, err, resources.MsgCodeRequestDataInvalid)
		logAuth.Error("Validate Register Request Failed With Err = ", "["+err.Error()+"]")
		return result
	}

	dataReq, _ := json.Marshal(customer)
	logAuth.Data("Data Customer Request = ", string(dataReq))

	// validate email
	err = commons.ValidatorEmailCommon(customer.Email)
	if err != nil {
		result = commons.TemplateErrorCommon(c, err, resources.MsgCodeRequestDataInvalid)
		logAuth.Error("Validate Email Request Failed With Err = ", "["+err.Error()+"]")
		return result
	}

	// parse slug from last name and first name
	fullName := utils.JoinString([]string{customer.LastName, customer.FirstName})
	slug := utils.SlugifyString(fullName)

	// validate duplicated
	filter := bson.D{
		{"$or", bson.A{
			bson.D{{"email", customer.Email}},
			bson.D{{"slug", slug}},
		}},
	}
	isDuplicated := commons.ValidatorDuplicateCommon(constants.CustomerModel, filter)
	logAuth.Debug("Customer Is Duplicated = ", isDuplicated)
	if isDuplicated {
		err = fmt.Errorf("Customer has duplicate slug is (%s) or email is (%s)", slug, customer.Email)
		result = commons.TemplateErrorCommon(c, err, resources.MsgCodeRequestDuplicatedError)
		logAuth.Error("Validate Duplicate Request Failed With Err = ", "["+err.Error()+"]")
		return result
	}
	// hash password
	hashPass := utils.HashPassword(customer.Password)
	// audit attribute
	createdAt, updatedAt := utils.AuditAttribute("create")

	newCustomer := models.CustomerModel{
		FirstName: customer.FirstName,
		LastName:  customer.LastName,
		Email:     customer.Email,
		Password:  hashPass,
		Slug:      slug,
		CreatedAt: createdAt,
		UpdatedAt: updatedAt,
	}

	err = commons.ValidatorStructCommon(&newCustomer)
	if err != nil {
		result = commons.TemplateErrorCommon(c, err, resources.MsgCodeRequestDataInvalid)
		logAuth.Error("Validate Register Struct Failed With Err = ", "["+err.Error()+"]")
		return result
	}

	data, err := database.CreateOne(constants.CustomerModel, newCustomer)
	if err != nil {
		result = commons.TemplateErrorCommon(c, err, resources.MsgCodeRequestDatabaseError)
		logAuth.Error("Insert One Customer Failed With Err=", "["+err.Error()+"]")
		return result
	}

	dataRes, _ := json.Marshal(data)
	logAuth.Data("Data Customer Response = ", string(dataRes))

	customerDTO := cusMapper.ToCreateDto(data)

	res := commons.TemplateSuccessCommon(c, customerDTO, resources.MsgCodeRequestDataIsValid)
	logAuth.Info("[END] AuthRegister Service...")
	return res
}
