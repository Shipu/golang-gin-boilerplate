package main

import (
	"golang-gin-boilerplate/artifact"
	"golang-gin-boilerplate/routes"
)

func main() {
	artifact.Start()

	routes.RegisterRoute()

	artifact.Run()
}
