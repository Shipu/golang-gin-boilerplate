package models

import (
	"github.com/shipu/artifact"
	"gorm.io/gorm"
	"time"
)

var TaskModel *gorm.DB

type Task struct {
	//gorm.Model
	Id        uint      `json:"id" gorm:"primaryKey"`
	Task      string    `json:"task"`
	Status    string    `json:"status"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func TaskSetup() {
	err := artifact.DB.AutoMigrate(&Task{})
	if err != nil {
		return
	}
	TaskModel = artifact.DB.Model(&Task{}).Session(&gorm.Session{NewDB: true})
}
