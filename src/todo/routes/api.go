package routes

import (
	. "github.com/shipu/artifact"
	c "github.com/shipu/golang-gin-boilerplate/src/todo/controllers"
)

func TodoSetup() {
	v1 := Router.Group("api/v1")
	v1.GET("todos", c.TodoIndex())
	v1.POST("todos", c.TodoCreate())
	v1.GET("todos/:todoId", c.TodoShow())
	v1.PUT("todos/:todoId", c.TodoUpdate())
	v1.DELETE("todos/:todoId", c.TodoDelete())
}
