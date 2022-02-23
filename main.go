package main

import (
	"golang-gin-boilerplate/bootstrap"
	"golang-gin-boilerplate/routes"
)

func main() {
	bootstrap.NewApp()

	routes.RegisterRoute()

	bootstrap.Run()
}
