package routes

import (
	. "github.com/shipu/artifact"
	c "github.com/shipu/golang-gin-boilerplate/src/task/controllers"
)

func TaskSetup() {
	v1 := Router.Group("api/v1")
	v1.GET("tasks", c.TaskIndex())
	v1.POST("tasks", c.TaskCreate())
	v1.GET("tasks/:taskId", c.TaskShow())
	v1.PUT("tasks/:taskId", c.TaskUpdate())
	v1.DELETE("tasks/:taskId", c.TaskDelete())
}
