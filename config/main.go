package config

import (
	. "github.com/shipu/artifact"
	todo "github.com/shipu/golang-gin-boilerplate/src/todo/models"
)

func Register() {
	Config.AddConfig("App", new(AppConfig))
	Config.AddConfig("DB", new(DatabaseConfig))
	Config.AddConfig("NoSql", new(MongoConfig))
	Config.Load()
}

func Boot() {
	todo.TodoSetup()
}
