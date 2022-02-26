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
}

func NewMongoDB() *MongoDB {

	mongoUri := "mongodb+srv://" + Config.GetString("DB.Username") + ":" + Config.GetString("DB.Password") + "@" + Config.GetString("DB.Host") + "/" + Config.GetString("DB.Database") + "?retryWrites=true&w=majority"

	client, err := mongo.NewClient(options.Client().ApplyURI(mongoUri))
	if err != nil {
		log.Fatal(err)
	}

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
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

	return &MongoDB{Client: client, Database: database}
}
