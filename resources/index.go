package resources

// successes code
const (
	MsgCodeRequestDataSuccess = "S-00100"
)

// errors code
const (
	MsgCodeNotDefined         = "E-00100"
	MsgCodeRouteNotFound      = "E-00200"
	MsgCodeMethodNotFound     = "E-00300"
	MsgCodeRequestDataInvalid = "E-00400"
	MsgCodeRequestServerError = "E-00500"
)

var MapMsgStatus = map[string]int{
	MsgCodeRequestDataSuccess: 200,
	MsgCodeRequestDataInvalid: 400,
	MsgCodeRouteNotFound:      404,
	MsgCodeMethodNotFound:     405,
	MsgCodeRequestServerError: 500,
}

var MapMsgName = map[string]string{
	MsgCodeRequestDataSuccess: "RequestDataSuccess",
	MsgCodeNotDefined:         "ExceptionMsgCodeNotDefined",
	MsgCodeRouteNotFound:      "ExceptionRequestRouteNotFound",
	MsgCodeMethodNotFound:     "ExceptionRequestMethodNotFound",
	MsgCodeRequestDataInvalid: "ExceptionRequestDataInvalid",
	MsgCodeRequestServerError: "ExceptionRequestServerError",
}
