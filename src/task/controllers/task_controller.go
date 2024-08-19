package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/shipu/artifact"
	"github.com/shipu/golang-gin-boilerplate/src/task/dto"
	"github.com/shipu/golang-gin-boilerplate/src/task/services"
	"net/http"
)

// TaskIndex
// @Summary  all tasks
// @Schemes
// @Description  All tasks
// @Tags         Task
// @Accept       json
// @Produce      json
// @Param        page    query  string  false  "Page"
// @Param        limit   query  string  false  "Limit"
// @Param        status  query  string  false  "Status"
// @Success      200
// @Failure      401
// @Router       /tasks [get]
func TaskIndex() gin.HandlerFunc {
	return func(c *gin.Context) {
		page := c.DefaultQuery("page", "1")
		limit := c.DefaultQuery("limit", "10")
		status := c.DefaultQuery("status", "")

		var filter map[string]interface{} = make(map[string]interface{})
		filter["page"] = page
		filter["limit"] = limit
		filter["status"] = status

		todos, paginate, _ := services.AllTask(filter)

		artifact.Res.Code(200).Data(todos).Raw(map[string]interface{}{
			"meta": paginate,
		}).Json(c)
	}
}

// TaskCreate
// @Summary  create a task
// @Schemes
// @Description  create a task
// @Tags         Task
// @Accept       json
// @Produce      json
// @Param        request  body  dto.CreateTaskRequest  true  "Create Task Request"
// @Success      200
// @Failure      401
// @Router       /tasks [post]
func TaskCreate() gin.HandlerFunc {
	return func(c *gin.Context) {
		var createTodo dto.CreateTaskRequest

		defer func() {
			if err := recover(); err != nil {
				artifact.Res.Code(http.StatusUnprocessableEntity).Message("error").Data(err).Json(c)
			}
		}()

		if err := c.ShouldBind(&createTodo); err != nil {
			artifact.Res.Code(http.StatusBadRequest).Message("Bad Request").Data(err.Error()).AbortWithStatusJSON(c)
			return
		}

		todo := services.CreateATask(createTodo)

		artifact.Res.Code(http.StatusCreated).Message("success").Data(todo).Json(c)
	}
}

// TaskShow
// @Summary  task details
// @Schemes
// @Description  Task Details
// @Tags         Task
// @Accept       json
// @Produce      json
// @Param        taskId  path  string  true  "Task ID"
// @Success      200
// @Failure      401
// @Router       /tasks/{taskId} [get]
func TaskShow() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				artifact.Res.Code(http.StatusNotFound).Message(http.StatusText(http.StatusNotFound)).Json(c)
			}
		}()

		todoId := c.Param("todoId")

		todo := services.ATask(todoId)

		artifact.Res.Code(http.StatusOK).Message("success").Data(todo).Json(c)
	}
}

// TaskUpdate
// @Summary  update a task
// @Schemes
// @Description  update a task
// @Tags         Task
// @Accept       json
// @Produce      json
// @Param        taskId   path  string                 true  "Task ID"
// @Param        request  body  dto.UpdateTodoRequest  true  "Update Task Request"
// @Success      200
// @Failure      401
// @Router       /tasks/{taskId} [put]
func TaskUpdate() gin.HandlerFunc {
	return func(c *gin.Context) {
		var updateTodo dto.UpdateTaskRequest

		defer func() {
			if err := recover(); err != nil {
				artifact.Res.Code(http.StatusUnprocessableEntity).Message(http.StatusText(http.StatusUnprocessableEntity)).Data(err).Json(c)
			}
		}()

		taskId := c.Param("taskId")

		if err := c.ShouldBind(&updateTodo); err != nil {
			artifact.Res.Code(http.StatusBadRequest).Message(http.StatusText(http.StatusBadRequest)).Data(err.Error()).AbortWithStatusJSON(c)
			return
		}

		todo, err := services.UpdateATask(taskId, updateTodo)

		if err != nil {
			artifact.Res.Code(http.StatusInternalServerError).Message(http.StatusText(http.StatusInternalServerError)).Json(c)
			return
		}

		artifact.Res.Code(http.StatusOK).Message("Successfully Updated !!!").Data(todo).Json(c)
	}
}

// TaskDelete
// @Summary  delete a task
// @Schemes
// @Description  delete a task
// @Tags         Task
// @Accept       json
// @Produce      json
// @Param        todoId  path  string  true  "Task ID"
// @Success      200
// @Failure      422
// @Router       /tasks/{taskId} [delete]
func TaskDelete() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				artifact.Res.Code(http.StatusUnprocessableEntity).Message("error").Data(err).Json(c)
			}
		}()

		todoId := c.Param("todoId")
		err := services.DeleteATask(todoId)

		if !err {
			artifact.Res.Code(http.StatusInternalServerError).Message("something wrong").Json(c)
			return
		}

		artifact.Res.Code(http.StatusOK).Message("Successfully Delete !!!").Json(c)
	}
}
