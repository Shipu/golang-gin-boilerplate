package routes

import (
	. "github.com/shipu/artifact"
	c "github.com/shipu/golang-gin-boilerplate/pkg/todo/controllers"
)

func Setup() {
	Router.GET("todos", c.TodoIndex())
	Router.POST("todos", c.TodoCreate())
	Router.GET("todos/:todoId", c.TodoShow())
	Router.PUT("todos/:todoId", c.TodoUpdate())
	Router.DELETE("todos/:todoId", c.TodoDelete())
}
