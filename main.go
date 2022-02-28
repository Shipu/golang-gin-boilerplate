package main

import (
	"github.com/shipu/artifact"
	"github.com/shipu/golang-gin-boilerplate/config"
	"github.com/shipu/golang-gin-boilerplate/routes"
)

func main() {
	artifact.New()

	config.RegisterConfig()
	routes.RegisterRoute()

	artifact.Run()
}
