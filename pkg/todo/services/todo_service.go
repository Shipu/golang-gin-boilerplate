package services

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"golang-gin-boilerplate/pkg/todo/models"
	"log"
)

func AllTodo() []models.Todo {
	cursor, err, ctx := models.TodoCollection.Find(bson.M{})

	var todos []models.Todo

	if err != nil {
		log.Fatal(err)
	}
	if err = cursor.All(ctx, &todos); err != nil {
		log.Fatal(err)
	}

	defer func(cursor *mongo.Cursor, ctx context.Context) {
		var err = cursor.Close(ctx)
		if err != nil {

		}
	}(cursor, ctx)

	return todos
}

func CreateATodo(todo models.Todo) models.Todo {
	newEnrollment := models.Todo{
		Id:     primitive.NewObjectID(),
		Task:   todo.Task,
		Status: todo.Status,
	}

	result, err := models.TodoCollection.InsertOne(newEnrollment)
	if err != nil || result == nil {
		panic(err)
	}

	return newEnrollment
}

func UpdateATodo(todoId string, updateTodo models.Todo) (models.Todo, error) {

	objId, _ := primitive.ObjectIDFromHex(todoId)

	upsert := true
	after := options.After
	opt := options.FindOneAndUpdateOptions{
		ReturnDocument: &after,
		Upsert:         &upsert,
	}

	result := models.TodoCollection.FindOneAndUpdate(bson.M{"_id": objId}, bson.D{
		{"$set", bson.M{"task": updateTodo.Task, "status": updateTodo.Status}},
	}, &opt)

	if result.Err() != nil {
		log.Println("Err ", result.Err())
		return models.Todo{}, result.Err()
	}

	if err := result.Decode(&updateTodo); err != nil {
		return models.Todo{}, err
	}

	return updateTodo, nil
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

func DeleteATodo(todoId string) (error, bool) {
	var todo models.Todo

	objId, _ := primitive.ObjectIDFromHex(todoId)

	result := models.TodoCollection.FindOneAndDelete(bson.D{{"_id", objId}})

	if result.Err() != nil {
		return result.Err(), false
	}

	return result.Decode(&todo), true

}
