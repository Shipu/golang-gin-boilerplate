package routes

import (
	"golang-gin-boilerplate/artifact"
	. "golang-gin-boilerplate/cmd/todo/controllers"
)

func Setup() {
	artifact.Router.GET("todos", TodoIndex())
	artifact.Router.POST("todos", TodoCreate())
	artifact.Router.GET("todos/:todoId", TodoShow())
	artifact.Router.PUT("todos/:todoId", TodoUpdate())
	artifact.Router.DELETE("todos/:todoId", TodoDelete())
}
