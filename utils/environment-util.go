package utils

import (
	"os"

	constants "fast-food-api-client/constants"
	coreLogger "fast-food-api-client/core/logger"

	dotenv "github.com/joho/godotenv"
)

var logEnv = coreLogger.Logger(constants.AppName, constants.StructEnvUtil)

func GetEnv(envVariableName string, defaultValue string) string {
	err := dotenv.Load(".env")
	if err != nil {
		logEnv.Error("Failed to GetEnv. err = ", err)
	}

	var value string = os.Getenv(envVariableName)
	if len(value) > 0 {
		logEnv.Info(envVariableName, "=", value)
		return value
	}

	return defaultValue
}
