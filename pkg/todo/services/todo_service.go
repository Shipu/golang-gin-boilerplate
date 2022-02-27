package services

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"golang-gin-boilerplate/artifact"
	"golang-gin-boilerplate/pkg/todo/models"
	"log"
	"time"
)

func AllTodo() []models.Todo {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	cursor, err := artifact.Mongo.Database.Collection("todos").Find(ctx, bson.M{})
	defer cancel()

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
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	newEnrollment := models.Todo{
		Id:     primitive.NewObjectID(),
		Task:   todo.Task,
		Status: todo.Status,
	}

	result, err := artifact.Mongo.Database.Collection("todos").InsertOne(ctx, newEnrollment)
	if err != nil || result == nil {
		panic(err)
	}

	return newEnrollment
}

func UpdateATodo(todoId string) (models.Todo, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	var todo models.Todo
	defer cancel()

	objId, _ := primitive.ObjectIDFromHex(todoId)

	result := artifact.Mongo.Database.Collection("todos").FindOneAndUpdate(ctx, bson.M{"_id": objId}, bson.D{
		{"$set", bson.M{"task": todo.Task, "status": todo.Status}},
	})
	if result.Err() != nil {
		log.Println("Err ", result.Err())
		return models.Todo{}, result.Err()
	}
	if err := result.Decode(&todo); err != nil {
		return models.Todo{}, err
	}
	return todo, nil
}

func ATodo(todoId string) models.Todo {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	var todo models.Todo
	defer cancel()

	objId, _ := primitive.ObjectIDFromHex(todoId)

	err := artifact.Mongo.Database.Collection("todos").FindOne(ctx, bson.M{"_id": objId}).Decode(&todo)

	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	return todo
}

func DeleteATodo(todoId string) (error, bool) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	var todo models.Todo
	defer cancel()

	objId, _ := primitive.ObjectIDFromHex(todoId)

	result := artifact.Mongo.Database.Collection("todos").FindOneAndDelete(ctx, bson.D{{"_id", objId}})
	fmt.Println()

	if result.Err() != nil {
		return result.Err(), false
	}

	return result.Decode(&todo), true

}
