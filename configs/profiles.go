package profiles

import "fast-food-api-client/utils"

var AppPort = utils.GetEnv("APP_PORT", "8080")
