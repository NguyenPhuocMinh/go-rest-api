package utils

import (
	constants "fast-food-api-client/constants"
	coreLogger "fast-food-api-client/core/logger"

	"golang.org/x/crypto/bcrypt"
)

var logGenerator = coreLogger.Logger(constants.AppName, constants.StructGeneratorUtil)

var HashPassword = func(password string) string {
	logGenerator.Info("[BEGIN] HashPassword...")
	passHash, _ := bcrypt.GenerateFromPassword([]byte(password), 14)
	logGenerator.Info("[END] HashPassword...")
	return string(passHash)
}

var ComparePassword = func(hashedPassword, providedPassword string) error {
	logGenerator.Info("[BEGIN] ComparePassword...")

	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(providedPassword))
	if err != nil {
		return err
	}
	logGenerator.Info("[END] ComparePassword...")
	return nil
}
