package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Todo struct {
	Id     primitive.ObjectID `json:"id,omitempty" bson:"_id"`
	Task   string             `json:"task" bson:"task"`
	Status string             `json:"status" bson:"status"`
}
