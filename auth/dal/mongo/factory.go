package mongo

import (
	"context"
	"github.com/smbody/anonym/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"time"
)
var
	monga = initMonga()


func initMonga() *mongo.Database {
	mongoURI := config.GetString("mongo.uri")
	client, err := mongo.NewClient(options.Client().ApplyURI(mongoURI))
	if err != nil {
		log.Fatalf("Error occurred while establishing connection to %s", mongoURI)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	err = client.Ping(context.Background(), nil)
	if err != nil {
		log.Fatal(err)
	}
	return client.Database(config.GetString("mongo.name"))
}
