package repository

import (
	"MyFirstGoWebServer/internal/core"
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"strconv"
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
	filter := bson.M{"_id": id}
	documentReturned := collection.FindOne(ctx, filter)
	err := documentReturned.Decode(&person)
	log.Println("Get person document", person, err, filter)
	return person, err
}

func (p personRepositoryImpl) Fetch(ctx context.Context, emailFilter string, yearsOfExperienceWorkingFilter string, estimatedLevelFilter string) ([]*core.Person, error) {
	var results []*core.Person
	collection := p.db.Collection("persons")
	filter := bson.M{}
	if emailFilter != "" {
		filter["email"] = emailFilter
	}
	if yearsOfExperienceWorkingFilter != "" {
		yoew, _ := strconv.ParseFloat(yearsOfExperienceWorkingFilter, 32)
		filter["yearsOfExperienceWorking"] = bson.M{"$gte": yoew}
	}
	if estimatedLevelFilter != "" {
		filter["estimatedLevel"] = estimatedLevelFilter
	}

	opts := options.Find()
	opts.SetSort(bson.D{{"createdAt", -1}})
	cur, err := collection.Find(ctx, filter, opts)
	if err != nil {
		return nil, err
	}
	defer cur.Close(ctx)

	for cur.Next(ctx) {
		var elem core.Person
		err := cur.Decode(&elem)
		if err != nil {
			log.Println("Error while decoding person", err)
		}
		results = append(results, &elem)
	}
	return results, err
}

func (p personRepositoryImpl) Store(ctx context.Context, person *core.Person) (string, error) {
	collection := p.db.Collection("persons")
	insertResult, err := collection.InsertOne(ctx, person)
	if err != nil {
		log.Println("Error on inserting new person", err)
		return "", err
	}
	return insertResult.InsertedID.(string), nil
}
