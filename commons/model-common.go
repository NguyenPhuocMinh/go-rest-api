package commons

type ResponseModel struct {
	Data              interface{} `json:"data"`
	Error             interface{} `json:"error"`
	Endpoint          string      `json:"endpoint"`
	Method            string      `json:"method"`
	StatusCode        int         `json:"statusCode"`
	MessageCode       string      `json:"messageCode"`
	MessageName       string      `json:"messageName"`
	MessageTranslator string      `json:"messageTranslator"`
	Pagination        interface{} `json:"pagination"`
}
