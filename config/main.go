package config

import (
	. "github.com/shipu/artifact"
	todo "github.com/shipu/golang-gin-boilerplate/src/todo/models"
)

func Register() {
	Config.AddConfig("App", new(AppConfig))
	Config.AddConfig("DB", new(DatabaseConfig))
	Config.Load()
}

func Boot() {
	todo.Setup()
}
