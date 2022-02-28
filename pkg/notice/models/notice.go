package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang-gin-boilerplate/artifact"
)

var NoticeCollection artifact.MongoCollection = artifact.Mongo.Collection("notices")

type Notice struct {
	Id     primitive.ObjectID `json:"id,omitempty" bson:"_id"`
	Task   string             `json:"task" bson:"task"`
    Status string             `json:"status" bson:"status"`
}
