package configs

import "fast-food-api-client/utils"

var (
	AppPort      = utils.GetEnv("APP_PORT", "8080")
	AppMongoURI  = utils.GetEnv("APP_MONGO_URI", "")
	AppSecretKey = utils.GetEnv("APP_SECRET_KEY", "")
	AppIssuer    = utils.GetEnv("APP_ISSUER", "")
	AppAudience  = utils.GetEnv("APP_AUDIENCE", "")
)
