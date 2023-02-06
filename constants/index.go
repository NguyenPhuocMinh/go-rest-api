package constants

const AppName = "fast-food-api"

// main
const StructMain = "main"
const StructController = "controller"
const StructService = "service"
const StructTranslator = "translator"

// utils
const (
	StructEnvUtil    = "environment-util"
	StructMapperUtil = "mapper-util"
)

// middleware
const (
	StructAuthMiddleware     = "auth-middleware"
	StructLoggerMiddleware   = "logger-middleware"
	StructNoMethodMiddleware = "no-method-middleware"
	StructNoRouteMiddleware  = "no-route-middleware"
)

// languages
const (
	LangEN = "en"
	LangVN = "vi"
)

// log types
const (
	TypeError = "ERROR"
	TypeWarn  = "WARN"
	TypeInfo  = "INFO"
	TypeDebug = "DEBUG"
)

// status
const (
	StatusOK                  = 200
	StatusCreated             = 201
	StatusAccepted            = 202
	StatusNoContent           = 204
	StatusBadRequest          = 400
	StatusUnauthorized        = 401
	StatusForbidden           = 403
	StatusNotFound            = 404
	StatusMethodNotAllowed    = 405
	StatusConflict            = 409
	StatusInternalServerError = 500
)

// msg types
const (
	MsgTypeAuth = "Auth"
)

// msg actions
const (
	MsgActionAuthLogin = "AuthLogin"
)

// controllers
const StructAuthController = "auth-controller"

// services
const (
	StructLookupService = "lookup-service"
	StructAuthService   = "auth-service"
)
