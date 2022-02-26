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
	client   *mongo.Client
	database *mongo.Database
}

func NewMongoDB() *MongoDB {

	mongoUri := "mongodb+srv://" + Config.GetString("Database.Username") + ":" + Config.GetString("Database.Password") + "@" + Config.GetString("Database.Host") + "/" + Config.GetString("Database.Database") + "?retryWrites=true&w=majority"

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

	database := client.Database(Config.GetString("Database.Database"))

	return &MongoDB{client: client, database: database}
}
