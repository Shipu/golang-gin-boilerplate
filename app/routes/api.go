package routes

import (
	"github.com/gin-gonic/gin"
	. "golang-gin-boilerplate/artifact"
)

func Setup() {
	Router.GET("/", func(c *gin.Context) {
		Res.Message("Done").Data("ok").Json(c)
	})
}
