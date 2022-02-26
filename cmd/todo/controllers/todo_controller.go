package controllers

import (
	"github.com/gin-gonic/gin"
	"golang-gin-boilerplate/artifact"
	"golang-gin-boilerplate/cmd/todo/services"
)

func TodoIndex() gin.HandlerFunc {
	return func(c *gin.Context) {
		todos := services.AllTodo()

		artifact.Res.Data(todos).Json(c)
	}
}

func TodoCreate() gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}

func TodoShow() gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}

func TodoUpdate() gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}

func TodoDelete() gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}
