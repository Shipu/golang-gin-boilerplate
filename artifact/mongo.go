package artifact

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"time"
)

var Mongo *MongoDB

type MongoDB struct {
	Client   *mongo.Client
	Database *mongo.Database
	Ctx      context.Context
}

type MongoCollection struct {
	*mongo.Collection
	Ctx        context.Context
	CancelFunc context.CancelFunc
}

func NewMongoDB() *MongoDB {

	mongoUri := "mongodb+srv://" + Config.GetString("DB.Username") + ":" + Config.GetString("DB.Password") + "@" + Config.GetString("DB.Host") + "/" + Config.GetString("DB.Database") + "?retryWrites=true&w=majority"

	client, err := mongo.NewClient(options.Client().ApplyURI(mongoUri))
	if err != nil {
		log.Fatal(err)
	}

	ctx := context.Background()
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}

	//ping the database
	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("[C-Log] Connected to MongoDB")

	database := client.Database(Config.GetString("DB.Database"))

	return &MongoDB{Client: client, Database: database, Ctx: ctx}
}

func (mongodb *MongoDB) Collection(name string) *MongoCollection {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	return &MongoCollection{mongodb.Database.Collection(name), ctx, cancel}
}

func (collection MongoCollection) Find(filter interface{},
	opts ...*options.FindOptions) (*mongo.Cursor, error, context.Context) {

	cursor, err := collection.Collection.Find(nil, filter, opts...)
	return cursor, err, collection.Ctx
}

func (collection MongoCollection) InsertOne(document interface{},
	opts ...*options.InsertOneOptions) (*mongo.InsertOneResult, error) {

	return collection.Collection.InsertOne(nil, document, opts...)
}

func (collection MongoCollection) FindOneAndUpdate(filter interface{},
	update interface{}, opts ...*options.FindOneAndUpdateOptions) *mongo.SingleResult {

	return collection.Collection.FindOneAndUpdate(nil, filter, update, opts...)
}

func (collection MongoCollection) FindOne(filter interface{},
	opts ...*options.FindOneOptions) *mongo.SingleResult {

	result := collection.Collection.FindOne(nil, filter, opts...)

	return result
}

func (collection MongoCollection) FindOneAndDelete(filter interface{},
	opts ...*options.FindOneAndDeleteOptions) *mongo.SingleResult {

	return collection.Collection.FindOneAndDelete(nil, filter, opts...)
}
