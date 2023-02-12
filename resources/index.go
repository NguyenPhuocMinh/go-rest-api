package resources

import "fast-food-api-client/constants"

// successes code
const (
	MsgCodeRequestDataIsValid = "S-00100"
)

// errors code
const (
	MsgCodeRequestNotDefined      = "E-00100"
	MsgCodeRequestRouteNotFound   = "E-00200"
	MsgCodeRequestMethodNotFound  = "E-00300"
	MsgCodeRequestServiceNotFound = "E-00400"
	MsgCodeRequestDataInvalid     = "E-00500"
	MsgCodeRequestServerError     = "E-00600"
	MsgCodeRequestDatabaseError   = "E-00700"
	MsgCodeRequestDuplicatedError = "E-00800"
	MsgCodeRequestEmailInValid    = "E-00900"
)

var MapMsgStatus = map[string]int{
	MsgCodeRequestDataIsValid:     200,
	MsgCodeRequestNotDefined:      400,
	MsgCodeRequestRouteNotFound:   404,
	MsgCodeRequestMethodNotFound:  405,
	MsgCodeRequestServiceNotFound: 404,
	MsgCodeRequestDataInvalid:     400,
	MsgCodeRequestServerError:     500,
	MsgCodeRequestDatabaseError:   500,
	MsgCodeRequestDuplicatedError: 409,
	MsgCodeRequestEmailInValid:    400,
}

var MapMsgName = map[string]string{
	MsgCodeRequestDataIsValid:     "RequestDataIsValid",
	MsgCodeRequestNotDefined:      "ExceptionMsgCodeNotDefined",
	MsgCodeRequestRouteNotFound:   "ExceptionRequestRouteNotFound",
	MsgCodeRequestMethodNotFound:  "ExceptionRequestMethodNotFound",
	MsgCodeRequestServiceNotFound: "ExceptionRequestServiceNotFound",
	MsgCodeRequestDataInvalid:     "ExceptionRequestDataInvalid",
	MsgCodeRequestServerError:     "ExceptionRequestServerError",
	MsgCodeRequestDatabaseError:   "ExceptionRequestDatabaseError",
	MsgCodeRequestDuplicatedError: "ExceptionRequestDuplicatedError",
	MsgCodeRequestEmailInValid:    "ExceptionRequestEmailInValid",
}

var MapMsgTranslator = map[string]map[string]string{
	constants.LangEN: {
		MsgCodeRequestNotDefined:      "Exception message code not defined",
		MsgCodeRequestRouteNotFound:   "Exception request route not found",
		MsgCodeRequestMethodNotFound:  "Exception request method not found",
		MsgCodeRequestServiceNotFound: "Exception request service not found",
		MsgCodeRequestServerError:     "Exception request internal server error",
		MsgCodeRequestDatabaseError:   "Exception request database error",
		MsgCodeRequestDataInvalid:     "Exception request data invalid",
		MsgCodeRequestDuplicatedError: "Exception request duplicate data",
		MsgCodeRequestEmailInValid:    "Exception request email invalid",
		MsgCodeRequestDataIsValid:     "Request data is valid",
	},
	constants.LangVN: {
		MsgCodeRequestNotDefined:      "Exception message code not defined",
		MsgCodeRequestRouteNotFound:   "Exception request route not found",
		MsgCodeRequestMethodNotFound:  "Exception request method not found",
		MsgCodeRequestServiceNotFound: "Exception request service not found",
		MsgCodeRequestServerError:     "Exception request internal server error",
		MsgCodeRequestDatabaseError:   "Exception request database error",
		MsgCodeRequestDataInvalid:     "Exception request data invalid",
		MsgCodeRequestDuplicatedError: "Exception request duplicate data",
		MsgCodeRequestEmailInValid:    "Exception request email invalid",
		MsgCodeRequestDataIsValid:     "Request data is valid",
	},
}
