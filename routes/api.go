package routes

import (
	"github.com/gin-gonic/gin"
	. "golang-gin-boilerplate/artifact"
	"golang-gin-boilerplate/cmd/todo/routes"
)

func RegisterRoute() {
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

	routes.Setup()
}
