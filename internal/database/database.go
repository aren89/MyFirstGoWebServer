package database

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"os"
	"time"
)

func Init() (*mongo.Database, *mongo.Client) {
	uri := os.Getenv("DB_URI")
	name := os.Getenv("DB_NAME")
	clientOptions := options.Client().ApplyURI(uri).SetConnectTimeout(30 * time.Second).SetReplicaSet("rs0")
	client, err := mongo.Connect(context.Background(), clientOptions)

	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(context.Background(), nil)
	if err != nil {
		log.Fatal(err)
	}
	return client.Database(name), client
}
