package services

import (
	"fmt"
	pagination "github.com/gobeam/mongo-go-pagination"
	"github.com/shipu/golang-gin-boilerplate/src/todo/dto"
	"github.com/shipu/golang-gin-boilerplate/src/todo/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"strconv"
	"time"
)

func AllTodo(requestFilter map[string]interface{}) ([]models.Todo, pagination.PaginationData) {
	var todos []models.Todo

	filter := bson.M{}

	if requestFilter["status"] != "" {
		filter["status"] = requestFilter["status"]
	}

	page, _ := strconv.ParseInt(requestFilter["page"].(string), 10, 64)
	limit, _ := strconv.ParseInt(requestFilter["limit"].(string), 10, 64)

	paginatedData, err := pagination.New(models.TodoCollection.Collection).
		Page(page).
		Limit(limit).
		Sort("created_at", -1).
		Decode(&todos).
		Filter(filter).
		Find()

	if err != nil {
		panic(err)
	}
	return todos, paginatedData.Pagination
}

func CreateATodo(createTodoDto dto.CreateTodoRequest) models.Todo {
	todo := models.Todo{
		Id:        primitive.NewObjectID(),
		Todo:      createTodoDto.Todo,
		Status:    createTodoDto.Status,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	result, err := models.TodoCollection.InsertOne(todo)
	if err != nil || result == nil {
		panic(err)
	}

	return todo
}

func UpdateATodo(todoId string, updateTodoDto dto.UpdateTodoRequest) (models.Todo, error) {

	objId, _ := primitive.ObjectIDFromHex(todoId)

	upsert := true
	after := options.After
	opt := options.FindOneAndUpdateOptions{
		ReturnDocument: &after,
		Upsert:         &upsert,
	}

	result := models.TodoCollection.FindOneAndUpdate(
		bson.M{"_id": objId},
		bson.D{
			{"$set", bson.M{
				"todo":       updateTodoDto.Todo,
				"status":     updateTodoDto.Status,
				"updated_at": time.Now(),
			}},
		},
		&opt,
	)

	if result.Err() != nil {
		log.Println("Err ", result.Err())
		return models.Todo{}, result.Err()
	}

	var todo models.Todo
	if err := result.Decode(&todo); err != nil {
		return models.Todo{}, err
	}

	return todo, nil
}

func ATodo(todoId string) models.Todo {
	var todo models.Todo

	objId, _ := primitive.ObjectIDFromHex(todoId)

	err := models.TodoCollection.FindOne(bson.M{"_id": objId}).Decode(&todo)

	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	return todo
}

func DeleteATodo(todoId string) bool {
	objId, _ := primitive.ObjectIDFromHex(todoId)

	result := models.TodoCollection.FindOneAndDelete(bson.D{{"_id", objId}})

	if result.Err() != nil {
		return false
	}

	return true
}
