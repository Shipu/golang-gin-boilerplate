package models

import (
	"context"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang-gin-boilerplate/artifact"
	"time"
)

var TodoCollection *artifact.MongoCollection = artifact.Mongo.Collection("todos")

var Ctx, Cancel = context.WithTimeout(context.Background(), 10*time.Second)

type Todo struct {
	Id     primitive.ObjectID `json:"id,omitempty" bson:"_id"`
	Task   string             `json:"task" bson:"task"`
	Status string             `json:"status" bson:"status"`
}
