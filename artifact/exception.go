package artifact

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func ExceptionHandler(router *Router) {
	router.Use(gin.CustomRecovery(func(c *gin.Context, recovered interface{}) {
		if err, ok := recovered.(string); ok {
			c.String(http.StatusInternalServerError, fmt.Sprintf("error: %s", err))
		}
		c.AbortWithStatus(http.StatusInternalServerError)
	}))
}
