package common

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"time"
)


var DB *mongo.Database

func Init() *mongo.Client {
	clientOptions := options.Client().ApplyURI("mongodb://mongodb:27017").SetConnectTimeout(30*time.Second)
	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}
	DB = client.Database("myDB")
	return client
}

func GetDB() *mongo.Database {
	return DB
}