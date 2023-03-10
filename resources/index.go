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
		MsgCodeRequestDataSuccess:            "Y??u c???u d??? li???u th??nh c??ng",
		MsgCodeRequestNotDefined:             "M?? th??ng b??o ngo???i l??? kh??ng ???????c x??c ?????nh",
		MsgCodeRequestRouteNotFound:          "Kh??ng t??m th???y tuy???n ???????ng y??u c???u ngo???i l???",
		MsgCodeRequestMethodNotFound:         "Kh??ng t??m th???y ph????ng th???c y??u c???u ngo???i l???",
		MsgCodeRequestServiceNotFound:        "Kh??ng t??m th???y d???ch v??? y??u c???u ngo???i l???",
		MsgCodeRequestServerError:            "Y??u c???u ngo???i l??? l???i m??y ch??? n???i b???",
		MsgCodeRequestDatabaseError:          "L???i c?? s??? d??? li???u y??u c???u ngo???i l???",
		MsgCodeRequestDataInvalidError:       "D??? li???u y??u c???u ngo???i l??? kh??ng h???p l???",
		MsgCodeRequestDuplicatedError:        "Ngo???i l??? y??u c???u d??? li???u tr??ng l???p",
		MsgCodeRequestEmailInValidError:      "Email y??u c???u ngo???i l??? kh??ng h???p l???",
		MsgCodeRequestEmailNotFoundError:     "Kh??ng t??m th???y email y??u c???u ngo???i l???",
		MsgCodeRequestPasswordInCorrectError: "Email y??u c???u ngo???i l??? kh??ng ch??nh x??c",
		MsgCodeRequestSingedTokenError:       "Y??u c???u ngo???i l??? l???i m?? th??ng b??o ???? k??",
		MsgCodeRequestTokenNotFoundError:     "Kh??ng t??m th???y m?? th??ng b??o trong ti??u ?????",
		MsgCodeRequestTokenIsInValidError:    "",
		MsgCodeRequestTokenIsExpiredError:    "",
	},
}
