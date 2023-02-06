package commons

import "github.com/gin-gonic/gin"

func (r *ResponseModel) ResponseCommon(c *gin.Context) {
	res := r.BuildJSON(c)
	c.JSON(r.StatusCode, res)
}

func (r *ResponseModel) BuildJSON(c *gin.Context) *gin.H {
	r.Endpoint = c.Request.URL.String()
	r.Method = c.Request.Method

	res := gin.H{
		"data":              r.Data,
		"error":             r.Error,
		"endpoint":          r.Endpoint,
		"method":            r.Method,
		"statusCode":        r.StatusCode,
		"messageCode":       r.MessageCode,
		"messageName":       r.MessageName,
		"messageTranslator": r.MessageTranslator,
		"pagination":        r.Pagination,
	}

	return &res
}
