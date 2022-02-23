package artifact

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// Router -> Gin Router
type Router struct {
	*gin.Engine
}

//NewRouter : all the routes are defined here
func NewRouter() *Router {

	gin.ForceConsoleColor()

	//gin.SetMode("debug")

	httpRouter := gin.New()

	//httpRouter.SetTrustedProxies([]string{"0.0.0.0"})

	httpRouter.GET("/health-check", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"data": "Up and Running"})
	})

	return &Router{
		httpRouter,
	}
}
