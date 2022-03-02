package controllers

import "C"
import (
	"github.com/gin-gonic/gin"
	"github.com/shipu/artifact"
	"github.com/shipu/golang-gin-boilerplate/src/todo/models"
	"github.com/shipu/golang-gin-boilerplate/src/todo/services"
	"net/http"
)

func TodoIndex() gin.HandlerFunc {
	return func(c *gin.Context) {
		todos := services.AllTodo()

		artifact.Res.Data(todos).Json(c)
	}
}

func TodoCreate() gin.HandlerFunc {
	return func(c *gin.Context) {
		var todo models.Todo

		defer func() {
			if err := recover(); err != nil {
				artifact.Res.Code(http.StatusUnprocessableEntity).Message("error").Data(err).Json(c)
			}
		}()

		if err := c.ShouldBind(&todo); err != nil {
			artifact.Res.Code(http.StatusBadRequest).Message("Bad Request").Data(err.Error()).AbortWithStatusJSON(c)
			return
		}

		todo = services.CreateATodo(todo)

		artifact.Res.Code(http.StatusCreated).Message("success").Data(todo).Json(c)
	}
}

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

func TodoUpdate() gin.HandlerFunc {
	return func(c *gin.Context) {
		var updateTodo models.Todo

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

func TodoDelete() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				artifact.Res.Code(http.StatusUnprocessableEntity).Message("error").Data(err).Json(c)
			}
		}()

		todoId := c.Param("todoId")
		todo, err := services.DeleteATodo(todoId)

		if !err {
			artifact.Res.Code(http.StatusInternalServerError).Message("something wrong").Json(c)
			return
		}

		artifact.Res.Code(http.StatusOK).Message("Successfully Delete !!!").Data(todo).Json(c)
	}
}
