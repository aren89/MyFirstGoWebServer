package person

import (
	"MyFirstGoWebServer/src/common"
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"log"
	"time"
)

type Person struct {
	FirstName  string    `json:"first_name" binding:"required"`
	LastName   string    `json:"last_name" binding:"required"`
	Email      string    `json:"email" binding:"required,email"`
	Age        int       `json:"age" binding:"required"`
	Gender     string    `json:"gender"`
	Date       time.Time `json:"date" binding:"required"`
	Experience string    `json:"date" binding:"required"`
}

func PersistPerson(p Person) (interface{}, error) {
	db := common.GetDB()
	collection := db.Collection("persons")
	insertResult, err := collection.InsertOne(context.TODO(), p)
	if err != nil {
		log.Fatalln("Error on inserting new person", err)
		return 0, err
	}
	return insertResult.InsertedID, nil
}

func RetrievePersonFromDB(id string) Person {
	var person Person
	db := common.GetDB()
	collection := db.Collection("persons")
	objID, _ := primitive.ObjectIDFromHex(id)
	filter := bson.M{"_id": objID}
	documentReturned := collection.FindOne(context.TODO(), filter)
	err := documentReturned.Decode(&person)
	log.Println("Get person document", person, err, filter)
	return person
}
