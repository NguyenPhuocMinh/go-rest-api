package commons

import (
	"fmt"

	"fast-food-api-client/core/database"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func ValidatorRequestBodyCommon(c *gin.Context, s interface{}) error {
	var err error

	// validate the request body
	if err = c.BindJSON(&s); err != nil {
		fmt.Println("Validate Request Body Throw ErrMsg = ", err)
		return err
	}

	return nil
}

func ValidatorStructCommon(s interface{}) error {
	validate := validator.New()

	// validate the required field
	err := validate.Struct(s)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			errMsg := fmt.Errorf("StructField (field is %s) validation %s (value is %s)", err.StructField(), err.Tag(), err.Type())
			fmt.Println("Validate Struct Throw ErrMsg = ", errMsg)
		}
		return err
	}

	return nil
}

func ValidatorEmailCommon(s string) error {
	validate := validator.New()

	errs := validate.Var(s, "required,email")

	if errs != nil {
		fmt.Println("Validate Email Throw Err = ", errs) // output: Key: "" Error:Field validation for "" failed on the "email" tag
		return errs
	}

	return nil
}

func ValidatorDuplicateCommon(schemaType string, filters interface{}) bool {
	count := database.Count(schemaType, filters)

	if count >= 1 {
		return true
	}

	return false
}
