package models

import (
	"github.com/shipu/artifact"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var TodoCollection artifact.MongoCollection

type Todo struct {
	Id     primitive.ObjectID `json:"id,omitempty" bson:"_id"`
	Task   string             `json:"task" bson:"task"`
	Status string             `json:"status" bson:"status"`
}

func Setup() {
	TodoCollection = artifact.Mongo.Collection("todos")
}
