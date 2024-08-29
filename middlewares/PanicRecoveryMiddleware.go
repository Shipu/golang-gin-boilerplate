package middlewares

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/shipu/artifact"
	"net/http"
)

func PanicRecoveryMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer handlePanic(c)
		c.Next()
	}
}

func handlePanic(c *gin.Context) {
	if err := recover(); err != nil {
		var errStr string
		switch v := err.(type) {
		case string:
			errStr = v
		case error:
			errStr = v.Error()
		default:
			errStr = fmt.Sprintf("recovered from: %v", v)
		}
		artifact.Res.Code(http.StatusBadRequest).Message(errStr).Json(c)
	}
}
