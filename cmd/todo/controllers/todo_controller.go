package controllers

import "C"
import (
	"github.com/gin-gonic/gin"
	"golang-gin-boilerplate/artifact"
	"golang-gin-boilerplate/cmd/todo/models"
	"golang-gin-boilerplate/cmd/todo/services"
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
				artifact.Res.Status(http.StatusUnprocessableEntity).Message("error").Data(err).Json(c)
			}
		}()

		if err := c.ShouldBind(&todo); err != nil {
			artifact.Res.Status(http.StatusBadRequest).Message("Bad Request").Data(err.Error()).AbortWithStatusJSON(c)
			return
		}

		todo = services.CreateATodo(todo)

		artifact.Res.Status(http.StatusCreated).Message("success").Data(todo).Json(c)
	}
}

func TodoShow() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				artifact.Res.Status(http.StatusUnprocessableEntity).Message("error").Data(err).Json(c)
			}
		}()

		todoId := c.Param("todoId")

		todo := services.ATodo(todoId)

		artifact.Res.Status(http.StatusOK).Message("success").Data(todo).Json(c)
	}
}

func TodoUpdate() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				artifact.Res.Status(http.StatusUnprocessableEntity).Message("error").Data(err).Json(c)
			}
		}()

		todoId := c.Param("todoId")

		todo, err := services.UpdateATodo(todoId)

		if err != nil {
			artifact.Res.Status(http.StatusInternalServerError).Message("something wrong").Json(c)
			return
		}

		artifact.Res.Status(http.StatusOK).Message("Successfully Delete !!!").Data(todo).Json(c)
	}
}

func TodoDelete() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				artifact.Res.Status(http.StatusUnprocessableEntity).Message("error").Data(err).Json(c)
				artifact.Res.Status(http.StatusUnprocessableEntity).Message("error").Data(err).Json(c)
			}
		}()

		todoId := c.Param("todoId")
		todo, err := services.DeleteATodo(todoId)

		if !err {
			artifact.Res.Status(http.StatusInternalServerError).Message("something wrong").Json(c)
			return
		}

		artifact.Res.Status(http.StatusOK).Message("Successfully Delete !!!").Data(todo).Json(c)
	}
}
