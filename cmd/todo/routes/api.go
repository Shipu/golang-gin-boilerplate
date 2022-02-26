package routes

import (
	"golang-gin-boilerplate/artifact"
	. "golang-gin-boilerplate/cmd/todo/controllers"
)

func Setup() {
	artifact.Router.GET("todos", TodoIndex())
}
