package artifact

import (
	"github.com/gin-gonic/gin"
)

type Response struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type ResponseBuilder struct {
	Response
}

func (response *ResponseBuilder) Status(status int) *ResponseBuilder {
	response.Response.Status = status
	return response
}

func (response *ResponseBuilder) Message(message string) *ResponseBuilder {
	response.Response.Message = message
	return response
}

func (response *ResponseBuilder) Data(data interface{}) *ResponseBuilder {
	response.Response.Data = data
	return response
}

func (response *ResponseBuilder) Build() interface{} {
	if response.Response.Status == 0 {
		response.Status(200)
	}

	return map[string]interface{}{"status": response.Response.Status, "message": response.Response.Message, "data": response.Response.Data}
}

func (response *ResponseBuilder) Json(c *gin.Context) {
	c.JSON(response.Response.Status, response.Build())
}
