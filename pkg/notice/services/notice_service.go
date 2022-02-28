package services

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"golang-gin-boilerplate/pkg/notice/models"
	"log"
)

func AllNotice() []models.Notice {
	cursor, err, ctx := models.NoticeCollection.Find(bson.M{})

	var notices []models.Notice

	if err != nil {
		log.Fatal(err)
	}
	if err = cursor.All(ctx, &notices); err != nil {
		log.Fatal(err)
	}

	defer func(cursor *mongo.Cursor, ctx context.Context) {
		var err = cursor.Close(ctx)
		if err != nil {

		}
	}(cursor, ctx)

	return notices
}

func CreateANotice(notice models.Notice) models.Notice {
	newEnrollment := models.Notice{
		Id:     primitive.NewObjectID(),
		Task:   notice.Task,
		Status: notice.Status,
	}

	result, err := models.NoticeCollection.InsertOne(newEnrollment)
	if err != nil || result == nil {
		panic(err)
	}

	return newEnrollment
}

func UpdateANotice(noticeId string, updateNotice models.Notice) (models.Notice, error) {

	objId, _ := primitive.ObjectIDFromHex(noticeId)

	upsert := true
	after := options.After
	opt := options.FindOneAndUpdateOptions{
		ReturnDocument: &after,
		Upsert:         &upsert,
	}

	result := models.NoticeCollection.FindOneAndUpdate(bson.M{"_id": objId}, bson.D{
		{"$set", bson.M{"task": updateNotice.Task, "status": updateNotice.Status}},
	}, &opt)

	if result.Err() != nil {
		log.Println("Err ", result.Err())
		return models.Notice{}, result.Err()
	}

	if err := result.Decode(&updateNotice); err != nil {
		return models.Notice{}, err
	}

	return updateNotice, nil
}

func ANotice(noticeId string) models.Notice {
	var notice models.Notice

	objId, _ := primitive.ObjectIDFromHex(noticeId)

	err := models.NoticeCollection.FindOne(bson.M{"_id": objId}).Decode(&notice)

	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	return notice
}

func DeleteANotice(noticeId string) (error, bool) {
	var notice models.Notice

	objId, _ := primitive.ObjectIDFromHex(noticeId)

	result := models.NoticeCollection.FindOneAndDelete(bson.D{{"_id", objId}})

	if result.Err() != nil {
		return result.Err(), false
	}

	return result.Decode(&notice), true

}
