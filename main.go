package main

import (
	"golang-gin-boilerplate/bootstrap"
)

func main() {
	bootstrap.NewApp()

	RegisterRoute()

	bootstrap.Run()
}
