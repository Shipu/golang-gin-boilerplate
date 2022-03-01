package routes

import (
	"github.com/gin-gonic/gin"
	. "github.com/shipu/artifact"
	todoRoute "github.com/shipu/golang-gin-boilerplate/src/todo/routes"
)

func Register() {
	BaseRoute()

	todoRoute.Setup()
}

func BaseRoute() {
	// Example Route
	Router.GET("/", func(c *gin.Context) {
		data := map[string]interface{}{
			"app": Config.GetString("App.Name"),
		}

		//or
		//data := gin.H{
		//	"message": "Hello World",
		//}

		Res.Status(200).
			Message("success").
			Data(data).Json(c)
	})
}
