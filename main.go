package main

import (
	"golang-gin-boilerplate/app/routes"
	"golang-gin-boilerplate/bootstrap"
)

func main() {
	bootstrap.Start()

	routes.Setup()

	bootstrap.Run()
}
