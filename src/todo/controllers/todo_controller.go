package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/shipu/artifact"
	"github.com/shipu/golang-gin-boilerplate/src/todo/dto"
	"github.com/shipu/golang-gin-boilerplate/src/todo/services"
	"net/http"
)

// TodoIndex
// @Summary  all todos
// @Schemes
// @Description  All todos
// @Tags         Todo
// @Accept       json
// @Produce      json
// @Param        page    query  string  false  "Page"
// @Param        limit   query  string  false  "Limit"
// @Param        status  query  string  false  "Status"
// @Success      200
// @Failure      401
// @Router       /todos [get]
func TodoIndex() gin.HandlerFunc {
	return func(c *gin.Context) {
        page := c.DefaultQuery("page", "1")
        limit := c.DefaultQuery("limit", "10")
        status := c.DefaultQuery("status", "")

        var filter map[string]interface{} = make(map[string]interface{})
        filter["page"] = page
        filter["limit"] = limit
        filter["status"] = status

        todos, paginate := services.AllTodo(filter)

        artifact.Res.Code(200).Data(todos).Raw(map[string]interface{}{
            "meta": paginate,
        }).Json(c)
	}
}

// TodoCreate
// @Summary  create a todo
// @Schemes
// @Description  create a todo
// @Tags         Todo
// @Accept       json
// @Produce      json
// @Param        request  body  dto.CreateTodoRequest  true  "Create Todo Request"
// @Success      200
// @Failure      401
// @Router       /todos [post]
func TodoCreate() gin.HandlerFunc {
	return func(c *gin.Context) {
		var createTodo dto.CreateTodoRequest

		defer func() {
			if err := recover(); err != nil {
				artifact.Res.Code(http.StatusUnprocessableEntity).Message("error").Data(err).Json(c)
			}
		}()

		if err := c.ShouldBind(&createTodo); err != nil {
			artifact.Res.Code(http.StatusBadRequest).Message("Bad Request").Data(err.Error()).AbortWithStatusJSON(c)
			return
		}

		todo := services.CreateATodo(createTodo)

		artifact.Res.Code(http.StatusCreated).Message("success").Data(todo).Json(c)
	}
}

// TodoShow
// @Summary  todo details
// @Schemes
// @Description  Todo Details
// @Tags         Todo
// @Accept       json
// @Produce      json
// @Param        todoId  path  string  true  "Todo ID"
// @Success      200
// @Failure      401
// @Router       /todos/{todoId} [get]
func TodoShow() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				artifact.Res.Code(http.StatusNotFound).Message(http.StatusText(http.StatusNotFound)).Json(c)
			}
		}()

		todoId := c.Param("todoId")

		todo := services.ATodo(todoId)

		artifact.Res.Code(http.StatusOK).Message("success").Data(todo).Json(c)
	}
}

// TodoUpdate
// @Summary  update a todo
// @Schemes
// @Description  update a todo
// @Tags         Todo
// @Accept       json
// @Produce      json
// @Param        todoId   path  string                 true  "Todo ID"
// @Param        request  body  dto.UpdateTodoRequest  true  "Update Todo Request"
// @Success      200
// @Failure      401
// @Router       /todos/{todoId} [put]
func TodoUpdate() gin.HandlerFunc {
	return func(c *gin.Context) {
		var updateTodo dto.UpdateTodoRequest

		defer func() {
			if err := recover(); err != nil {
				artifact.Res.Code(http.StatusUnprocessableEntity).Message(http.StatusText(http.StatusUnprocessableEntity)).Data(err).Json(c)
			}
		}()

		todoId := c.Param("todoId")

		if err := c.ShouldBind(&updateTodo); err != nil {
			artifact.Res.Code(http.StatusBadRequest).Message(http.StatusText(http.StatusBadRequest)).Data(err.Error()).AbortWithStatusJSON(c)
			return
		}

		todo, err := services.UpdateATodo(todoId, updateTodo)

		if err != nil {
			artifact.Res.Code(http.StatusInternalServerError).Message(http.StatusText(http.StatusInternalServerError)).Json(c)
			return
		}

		artifact.Res.Code(http.StatusOK).Message("Successfully Updated !!!").Data(todo).Json(c)
	}
}

// TodoDelete
// @Summary  delete a todo
// @Schemes
// @Description  delete a todo
// @Tags         Todo
// @Accept       json
// @Produce      json
// @Param        todoId  path  string  true  "Todo ID"
// @Success      200
// @Failure      422
// @Router       /todos/{todoId} [delete]
func TodoDelete() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				artifact.Res.Code(http.StatusUnprocessableEntity).Message("error").Data(err).Json(c)
			}
		}()

		todoId := c.Param("todoId")
		err := services.DeleteATodo(todoId)

		if !err {
			artifact.Res.Code(http.StatusInternalServerError).Message("something wrong").Json(c)
			return
		}

		artifact.Res.Code(http.StatusOK).Message("Successfully Delete !!!").Json(c)
	}
}