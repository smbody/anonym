package mongo

import (
	"context"
	"github.com/smbody/anonym/config"
	"github.com/smbody/anonym/errors"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"time"
)


func Init() *Users {
	return initUsers(initMonga())
}

func initMonga() *mongo.Database {
	mongoURI := config.DataSourceName()
	client, err := mongo.NewClient(options.Client().ApplyURI(mongoURI))
	if err != nil {
		log.Fatalf("Error occurred while establishing connection to %s", mongoURI)
		errors.Throw(errors.CantConnectToToDatabase)
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
	name := config.DatabaseName()
	log.Printf("Connect to Mongo (URI = %s, Database = %s)", mongoURI, name)
	return client.Database(name)
}
