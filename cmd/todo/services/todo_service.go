package services

import "golang-gin-boilerplate/artifact"

func AllTodo() interface{} {
	return map[string]interface{}{
		"todo": map[string]interface{}{
			"app": artifact.Config.GetString("App.Name"),
		},
	}
}

func CreateATodo() {

}

func UpdateATodo() {

}

func ATodo() {

}

func DeleteATodo() {

}
