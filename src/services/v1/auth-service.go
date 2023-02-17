package v1

import (
	"encoding/json"
	"fmt"
	"time"

	commons "fast-food-api-client/commons"
	configs "fast-food-api-client/configs"
	constants "fast-food-api-client/constants"
	database "fast-food-api-client/core/database"
	coreLogger "fast-food-api-client/core/logger"
	resources "fast-food-api-client/resources"
	mappers "fast-food-api-client/src/mappers"
	models "fast-food-api-client/src/models"
	utils "fast-food-api-client/utils"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var logAuth = coreLogger.Logger(constants.AppName, constants.StructAuthService)

var (
	authMapper = mappers.AuthMapper{}
	secretKey  = []byte(configs.AppSecretKey)
	issuer     = configs.AppIssuer
	audience   = configs.AppAudience
)

func AuthLogin(c *gin.Context) *commons.ResponseModel {
	logAuth.Info("[BEGIN] AuthLogin Service...")

	var result *commons.ResponseModel
	var customer models.AuthLoginModel
	var err error

	// validate request
	err = commons.ValidatorRequestBodyCommon(c, &customer)
	if err != nil {
		result = commons.TemplateErrorCommon(c, err, resources.MsgCodeRequestDataInvalidError)
		logAuth.Error("Validate Register Request Failed With Err = ", "["+err.Error()+"]")
		return result
	}

	dataReq, _ := json.Marshal(customer)
	logAuth.Data("Data Auth Login Request = ", string(dataReq))

	// validate email
	err = commons.ValidatorEmailCommon(customer.Email)
	if err != nil {
		result = commons.TemplateErrorCommon(c, err, resources.MsgCodeRequestDataInvalidError)
		logAuth.Error("Validate Email Request Failed With Err = ", "["+err.Error()+"]")
		return result
	}

	// check customer is exists in db
	var opts options.FindOneOptions
	filters := bson.D{{"email", customer.Email}}
	projection := bson.D{{"password", 1}, {"email", 1}}
	opts.SetProjection(projection)

	var authCompare *models.AuthComparePasswordModel
	err = database.GetOne(constants.CustomerModel, filters, &opts, &authCompare)
	if err != nil {
		result = commons.TemplateErrorCommon(c, err, resources.MsgCodeRequestEmailNotFoundError)
		logAuth.Error("Request Email Not Found Failed With Err = ", "["+err.Error()+"]")
		return result
	}

	// check compare password
	err = utils.ComparePassword(authCompare.Password, customer.Password)
	if err != nil {
		result = commons.TemplateErrorCommon(c, err, resources.MsgCodeRequestPasswordInCorrectError)
		logAuth.Error("Request Compare Password Failed With Err = ", "["+err.Error()+"]")
		return result
	}

	// sign token
	jti, _ := uuid.NewRandom() // => generator uuid v4()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"exp":   time.Now().Add(time.Hour * time.Duration(1)).Unix(), // 1 hours
		"id":    authCompare.ID,
		"email": authCompare.Email,
		"iss":   issuer,
		"aud":   audience,
		"jti":   jti.String(),
	})

	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		result = commons.TemplateErrorCommon(c, err, resources.MsgCodeRequestSingedTokenError)
		logAuth.Error("Request Signed Token Failed With Err = ", "["+err.Error()+"]")
		return result
	}

	authLoginDTO := authMapper.ToAuthLoginDTO(tokenString)

	dataRes, _ := json.Marshal(authLoginDTO)
	logAuth.Data("Data Auth Login Response = ", string(dataRes))

	res := commons.TemplateSuccessCommon(c, authLoginDTO, resources.MsgCodeRequestDataSuccess)
	logAuth.Info("[END] AuthLogin Service...")
	return res
}

func AuthRegister(c *gin.Context) *commons.ResponseModel {
	logAuth.Info("[BEGIN] AuthRegister Service...")

	var result *commons.ResponseModel
	var customer models.AuthRegisterModel
	var err error

	// validate request
	err = commons.ValidatorRequestBodyCommon(c, &customer)
	if err != nil {
		result = commons.TemplateErrorCommon(c, err, resources.MsgCodeRequestDataInvalidError)
		logAuth.Error("Validate Register Request Failed With Err = ", "["+err.Error()+"]")
		return result
	}

	dataReq, _ := json.Marshal(customer)
	logAuth.Data("Data Auth Register Request = ", string(dataReq))

	// validate email
	err = commons.ValidatorEmailCommon(customer.Email)
	if err != nil {
		result = commons.TemplateErrorCommon(c, err, resources.MsgCodeRequestDataInvalidError)
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

	newCustomer := models.CustomerRegisterModel{
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
		result = commons.TemplateErrorCommon(c, err, resources.MsgCodeRequestDataInvalidError)
		logAuth.Error("Validate Register Struct Failed With Err = ", "["+err.Error()+"]")
		return result
	}

	data, err := database.CreateOne(constants.CustomerModel, newCustomer)
	if err != nil {
		result = commons.TemplateErrorCommon(c, err, resources.MsgCodeRequestDatabaseError)
		logAuth.Error("Insert One Customer Failed With Err=", "["+err.Error()+"]")
		return result
	}

	authRegisterDTO := authMapper.ToAuthRegisterDTO(data)

	dataRes, _ := json.Marshal(authRegisterDTO)
	logAuth.Data("Data Auth Register Response = ", string(dataRes))

	res := commons.TemplateSuccessCommon(c, authRegisterDTO, resources.MsgCodeRequestDataSuccess)
	logAuth.Info("[END] AuthRegister Service...")
	return res
}
