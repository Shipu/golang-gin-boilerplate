package models

import (
	"github.com/shipu/artifact"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"gorm.io/gorm"
	"time"
)

var TaskModel *gorm.DB

type Task struct {
	Id        primitive.ObjectID `json:"id,omitempty" bson:"_id"`
	Task      string             `json:"task" bson:"task"`
	Status    string             `json:"status" bson:"status"`
	CreatedAt time.Time          `json:"created_at,omitempty" bson:"created_at"`
	UpdatedAt time.Time          `json:"updated_at,omitempty" bson:"updated_at"`
}

func TaskSetup() {
	err := artifact.DB.AutoMigrate(&Task{})
	if err != nil {
		return
	}
	TaskModel = artifact.DB.Model(&Task{}).Session(&gorm.Session{NewDB: true})
}
