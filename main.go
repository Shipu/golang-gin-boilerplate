package main

import (
	"github.com/shipu/artifact"
	"github.com/shipu/golang-gin-boilerplate/config"
	"github.com/shipu/golang-gin-boilerplate/middlewares"
	"github.com/shipu/golang-gin-boilerplate/routes"
)

// @title        Artifact Boilerplate
// @version      1.0
// @description  example of artifact web-framework

// @BasePath  /api/v1
func main() {
	// Initialize the application
	artifact.New()

	//global middleware to capture panic
	artifact.Router.Use(middlewares.PanicRecoveryMiddleware())

	config.Register() // will load the config file
	routes.Register() // will register all the routes

	// After Initialize Set up the application for serve
	artifact.NoSqlConnection()    // Mongo connection will be established here
	artifact.DatabaseConnection() // Relation Database connection will be established here
	config.Boot()                 // if you need any initialization

	artifact.Run()

}
