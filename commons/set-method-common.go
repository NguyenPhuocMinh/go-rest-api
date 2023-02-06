package commons

func (r *ResponseModel) SetData(data interface{}) *ResponseModel {
	r.Data = data
	return r
}

func (r *ResponseModel) SetError(err interface{}) *ResponseModel {
	r.Error = err
	return r
}

func (r *ResponseModel) SetStatusCode(statusCode int) *ResponseModel {
	r.StatusCode = statusCode
	return r
}

func (r *ResponseModel) SetMessageCode(messageCode string) *ResponseModel {
	r.MessageCode = messageCode
	return r
}

func (r *ResponseModel) SetMessageName(messageName string) *ResponseModel {
	r.MessageName = messageName
	return r
}

func (r *ResponseModel) SetMessageTranslator(messageTranslator string) *ResponseModel {
	r.MessageTranslator = messageTranslator
	return r
}

func (r *ResponseModel) SetPagination(pagination interface{}) *ResponseModel {
	r.Pagination = pagination
	return r
}
