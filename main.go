package main

import (
	"golang-gin-boilerplate/app/routes"
	"golang-gin-boilerplate/artifact"
)

func main() {
	artifact.Start()

	routes.Setup()

	artifact.Run()
}
