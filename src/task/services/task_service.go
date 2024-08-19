package services

import (
	"github.com/shipu/artifact"
	"github.com/shipu/golang-gin-boilerplate/src/task/dto"
	"github.com/shipu/golang-gin-boilerplate/src/task/models"
)

func AllTask(requestFilter map[string]interface{}) ([]models.Task, artifact.PaginationMeta, error) {
	var todos []models.Task

	filter := make(map[string]interface{})

	if requestFilter["status"] != "" {
		filter["status"] = requestFilter["status"]
	}

	paginationInstance := artifact.NewPaginator(todos, requestFilter)

	models.TaskModel.Where(filter).Scopes(paginationInstance.PaginateScope()).Find(&todos)

	return todos, paginationInstance.Meta, nil

	// cursor pagination demo
	//result, cursor, err := paginationInstance.Paginate(models.TaskTable.Where(filter), &todos)
	//if err != nil {
	//	return nil, cursor, err
	//}
	//
	//if result.Error != nil {
	//	return nil, paginator.Cursor{}, result.Error
	//}
	//
	//return tasks, cursor, nil
}

func CreateATask(createTodoDto dto.CreateTaskRequest) models.Task {
	todo := models.Task{
		Task:   createTodoDto.Task,
		Status: createTodoDto.Status,
	}

	result := models.TaskModel.Create(&todo)
	if result == nil {
		panic("Error creating todo")
	}

	return todo
}

func UpdateATask(todoId string, updateTodoDto dto.UpdateTaskRequest) (models.Task, error) {
	todo := models.Task{
		Task:   updateTodoDto.Task,
		Status: updateTodoDto.Status,
	}

	models.TaskModel.Where("id = ?", todoId).Updates(&todo)

	return todo, nil
}

func ATask(todoId string) models.Task {
	var todo models.Task

	models.TaskModel.First(&todo, todoId)

	return todo
}

func DeleteATask(todoId string) bool {
	var todo models.Task
	models.TaskModel.Delete(&todo, todoId)

	return true
}
