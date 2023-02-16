package resources

import "fast-food-api-client/constants"

// successes code
const (
	MsgCodeRequestDataSuccess = "S-00100"
)

// errors code
const (
	MsgCodeRequestNotDefined             = "E-00100"
	MsgCodeRequestRouteNotFound          = "E-00200"
	MsgCodeRequestMethodNotFound         = "E-00300"
	MsgCodeRequestServiceNotFound        = "E-00400"
	MsgCodeRequestDataInvalidError       = "E-00500"
	MsgCodeRequestServerError            = "E-00600"
	MsgCodeRequestDatabaseError          = "E-00700"
	MsgCodeRequestDuplicatedError        = "E-00800"
	MsgCodeRequestEmailInValidError      = "E-00900"
	MsgCodeRequestEmailNotFoundError     = "E-01000"
	MsgCodeRequestPasswordInCorrectError = "E-01100"
	MsgCodeRequestSingedTokenError       = "E-01200"
	MsgCodeRequestTokenNotFoundError     = "E-01300"
	MsgCodeRequestTokenIsInValidError    = "E-01400"
	MsgCodeRequestTokenIsExpiredError    = "E-01500"
)

var MapMsgStatus = map[string]int{
	MsgCodeRequestDataSuccess:            200,
	MsgCodeRequestNotDefined:             400,
	MsgCodeRequestRouteNotFound:          404,
	MsgCodeRequestMethodNotFound:         405,
	MsgCodeRequestServiceNotFound:        404,
	MsgCodeRequestDataInvalidError:       400,
	MsgCodeRequestServerError:            500,
	MsgCodeRequestDatabaseError:          500,
	MsgCodeRequestDuplicatedError:        409,
	MsgCodeRequestEmailInValidError:      400,
	MsgCodeRequestEmailNotFoundError:     400,
	MsgCodeRequestPasswordInCorrectError: 400,
	MsgCodeRequestSingedTokenError:       500,
	MsgCodeRequestTokenNotFoundError:     401,
	MsgCodeRequestTokenIsInValidError:    401,
	MsgCodeRequestTokenIsExpiredError:    401,
}

var MapMsgName = map[string]string{
	MsgCodeRequestDataSuccess:            "RequestDataSuccess",
	MsgCodeRequestNotDefined:             "ExceptionMsgCodeNotDefined",
	MsgCodeRequestRouteNotFound:          "ExceptionRequestRouteNotFound",
	MsgCodeRequestMethodNotFound:         "ExceptionRequestMethodNotFound",
	MsgCodeRequestServiceNotFound:        "ExceptionRequestServiceNotFound",
	MsgCodeRequestDataInvalidError:       "ExceptionRequestDataInvalid",
	MsgCodeRequestServerError:            "ExceptionRequestServerError",
	MsgCodeRequestDatabaseError:          "ExceptionRequestDatabaseError",
	MsgCodeRequestDuplicatedError:        "ExceptionRequestDuplicatedError",
	MsgCodeRequestEmailInValidError:      "ExceptionRequestEmailInValid",
	MsgCodeRequestEmailNotFoundError:     "ExceptionRequestEmailNotFound",
	MsgCodeRequestPasswordInCorrectError: "ExceptionRequestPasswordInCorrect",
	MsgCodeRequestSingedTokenError:       "ExceptionRequestSingedTokenError",
	MsgCodeRequestTokenNotFoundError:     "ExceptionRequestTokenNotFound",
	MsgCodeRequestTokenIsInValidError:    "ExceptionRequestTokenIsInValidError",
	MsgCodeRequestTokenIsExpiredError:    "ExceptionRequestTokenIsExpiredError",
}

var MapMsgTranslator = map[string]map[string]string{
	constants.LangEN: {
		MsgCodeRequestDataSuccess:            "Request data successfully",
		MsgCodeRequestNotDefined:             "Exception message code not defined",
		MsgCodeRequestRouteNotFound:          "Exception request route not found",
		MsgCodeRequestMethodNotFound:         "Exception request method not found",
		MsgCodeRequestServiceNotFound:        "Exception request service not found",
		MsgCodeRequestServerError:            "Exception request internal server error",
		MsgCodeRequestDatabaseError:          "Exception request database error",
		MsgCodeRequestDataInvalidError:       "Exception request data invalid",
		MsgCodeRequestDuplicatedError:        "Exception request duplicate data",
		MsgCodeRequestEmailInValidError:      "Exception request email invalid",
		MsgCodeRequestEmailNotFoundError:     "Exception request email not found",
		MsgCodeRequestPasswordInCorrectError: "Exception request password incorrect",
		MsgCodeRequestSingedTokenError:       "Exception request signed token error",
		MsgCodeRequestTokenNotFoundError:     "Exception request token not found in header",
		MsgCodeRequestTokenIsInValidError:    "",
		MsgCodeRequestTokenIsExpiredError:    "",
	},
	constants.LangVN: {
		MsgCodeRequestDataSuccess:            "Yêu cầu dữ liệu thành công",
		MsgCodeRequestNotDefined:             "Mã thông báo ngoại lệ không được xác định",
		MsgCodeRequestRouteNotFound:          "Không tìm thấy tuyến đường yêu cầu ngoại lệ",
		MsgCodeRequestMethodNotFound:         "Không tìm thấy phương thức yêu cầu ngoại lệ",
		MsgCodeRequestServiceNotFound:        "Không tìm thấy dịch vụ yêu cầu ngoại lệ",
		MsgCodeRequestServerError:            "Yêu cầu ngoại lệ lỗi máy chủ nội bộ",
		MsgCodeRequestDatabaseError:          "Lỗi cơ sở dữ liệu yêu cầu ngoại lệ",
		MsgCodeRequestDataInvalidError:       "Dữ liệu yêu cầu ngoại lệ không hợp lệ",
		MsgCodeRequestDuplicatedError:        "Ngoại lệ yêu cầu dữ liệu trùng lặp",
		MsgCodeRequestEmailInValidError:      "Email yêu cầu ngoại lệ không hợp lệ",
		MsgCodeRequestEmailNotFoundError:     "Không tìm thấy email yêu cầu ngoại lệ",
		MsgCodeRequestPasswordInCorrectError: "Email yêu cầu ngoại lệ không chính xác",
		MsgCodeRequestSingedTokenError:       "Yêu cầu ngoại lệ lỗi mã thông báo đã ký",
		MsgCodeRequestTokenNotFoundError:     "Không tìm thấy mã thông báo trong tiêu đề",
		MsgCodeRequestTokenIsInValidError:    "",
		MsgCodeRequestTokenIsExpiredError:    "",
	},
}
