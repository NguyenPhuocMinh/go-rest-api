package constants

const AppName = "fast-food-api"

// main
const StructMain = "main"

const (
	StructDatabase   = "database"
	StructController = "controller"
	StructService    = "service"
	StructTranslator = "translator"
)

// utils
const (
	StructEnvUtil        = "environment-util"
	StructMapperUtil     = "mapper-util"
	StructGeneratorUtil  = "generator-util"
	StructPaginationUtil = "pagination-util"
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

// mode
const (
	CreateMode = "create"
)

// log types
const (
	TypeError = "ERROR"
	TypeWarn  = "WARN"
	TypeInfo  = "INFO"
	TypeDebug = "DEBUG"
	TypeData  = "DATA"
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
	MsgTypeCommon = "Common"
	MsgTypeAuth   = "Auth"
)

// msg actions
const (
	MsgActionTest = "Test"
)

const (
	MsgActionAuthLogin    = "AuthLogin"
	MsgActionAuthRegister = "AuthRegister"
)

// controllers
const (
	StructCommonController = "common-controller"
	StructAuthController   = "auth-controller"
)

// services
const (
	StructLookupService = "lookup-service"
	StructCommonervice  = "common-service"
	StructAuthService   = "auth-service"
)
