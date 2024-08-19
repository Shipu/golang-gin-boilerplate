package config

import (
	. "github.com/shipu/artifact"
	task "github.com/shipu/golang-gin-boilerplate/src/task/models"
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
	task.TaskSetup()
}
