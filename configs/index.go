package configs

import "fast-food-api-client/utils"

var (
	AppPort     = utils.GetEnv("APP_PORT", "8080")
	AppMongoURI = utils.GetEnv("APP_MONGO_URI", "")
)
