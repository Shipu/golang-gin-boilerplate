package services

import (
	"fmt"
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
	//result, cursor, err := paginationInstance.Paginate(models.TaskModel.Where(filter), &todos)
	//if err != nil {
	//	return nil, cursor, err
	//}
	//
	//if result.Error != nil {
	//	return nil, paginator.Cursor{}, result.Error
	//}
	//
	//return todos, cursor, nil
}

func CreateATask(createTaskDto dto.CreateTaskRequest) models.Task {
	todo := models.Task{
		Task:   createTaskDto.Task,
		Status: createTaskDto.Status,
	}

	result := models.TaskModel.Create(&todo)
	if result == nil {
		panic("Error creating todo")
	}

	return todo
}

func UpdateATask(todoId string, updateTaskDto dto.UpdateTaskRequest) (models.Task, error) {
	todo := ATask(todoId)

	todo.Task = updateTaskDto.Task
	todo.Status = updateTaskDto.Status

	models.TaskModel.Save(&todo)

	return todo, nil
}

func ATask(todoId string) models.Task {
	var todo models.Task

	result := models.TaskModel.First(&todo, todoId)

	if result.Error != nil {
		panic("Task not found")
	}

	return todo
}

func DeleteATask(todoId string) bool {
	todo := ATask(todoId)
	result := models.TaskModel.Delete(&todo)
	fmt.Print("result", result.Error)

	if result.RowsAffected == 0 {
		return false
	}

	return true
}
