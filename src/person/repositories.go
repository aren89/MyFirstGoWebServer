package person

import (
	"MyFirstGoWebServer/src/common"
	"context"
	"log"
	"time"
)

type Person struct {
	FirstName  string
	LastName   string
	Age        int
	Gender     string
	Date       time.Time
	Experience string
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
