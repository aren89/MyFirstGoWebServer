package repository

import (
	"MyFirstGoWebServer/internal/core"
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
)

type personRepositoryImpl struct {
	db *mongo.Database
}

func NewPersonRepository(db *mongo.Database) core.PersonRepository {
	return &personRepositoryImpl{
		db: db,
	}
}

func (p personRepositoryImpl) GetByID(ctx context.Context, id string) (core.Person, error) {
	var person core.Person
	collection := p.db.Collection("persons")
	objID, _ := primitive.ObjectIDFromHex(id)
	filter := bson.M{"_id": objID}
	documentReturned := collection.FindOne(ctx, filter)
	err := documentReturned.Decode(&person)
	log.Println("Get person document", person, err, filter)
	return person, err
}

func (p personRepositoryImpl) Store(ctx context.Context, person *core.Person) (string, error) {
	collection := p.db.Collection("persons")
	insertResult, err := collection.InsertOne(ctx, person)
	if err != nil {
		log.Fatalln("Error on inserting new person", err)
		return "", err
	}
	return insertResult.InsertedID.(primitive.ObjectID).Hex(), nil
}
