package repository

import (
	"MyFirstGoWebServer/internal/core"
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
)

type applicationRepositoryImpl struct {
	db           *mongo.Database
	collection   *mongo.Collection
	jobOfferRepo core.JobOfferRepository
	client       *mongo.Client
}

func NewApplicationRepository(jobOfferRepo core.JobOfferRepository, db *mongo.Database, client *mongo.Client) core.ApplicationRepository {
	return &applicationRepositoryImpl{
		db:           db,
		collection:   db.Collection("applications"),
		jobOfferRepo: jobOfferRepo,
		client:       client,
	}
}

func (a applicationRepositoryImpl) Store(ctx context.Context, application *core.Application) (string, error) {
	session, err := a.client.StartSession()
	if err != nil {
		return "", err
	}
	defer session.EndSession(ctx)

	if err = session.StartTransaction(); err != nil {
		return "", err
	}
	if err = mongo.WithSession(ctx, session, func(sc mongo.SessionContext) error {
		_, err := a.collection.InsertOne(sc, application)
		if err != nil {
			log.Println("Error on inserting new application", err)
			return err
		}

		if err = a.jobOfferRepo.PushApplication(sc, application.JobOfferId, application.Id); err != nil {
			return err
		}

		if err = session.CommitTransaction(sc); err != nil {
			log.Println("cannot commit transaction", sc, err)
			return err
		}
		return nil
	}); err != nil {
		return "", err
	}
	return application.Id, nil
}

func (a applicationRepositoryImpl) Delete(ctx context.Context, id string) error {
	session, err := a.client.StartSession()
	if err != nil {
		return err
	}
	defer session.EndSession(ctx)

	if err = session.StartTransaction(); err != nil {
		return err
	}
	if err = mongo.WithSession(ctx, session, func(sc mongo.SessionContext) error {
		_, err := a.collection.DeleteOne(sc, bson.M{"_id": id})
		if err != nil {
			log.Println("Error on deleting application", err)
			return err
		}
		if err = a.jobOfferRepo.PopApplication(sc, id); err != nil {
			return err
		}

		if err = session.CommitTransaction(sc); err != nil {
			log.Println("cannot commit transaction", sc, err)
			return err
		}
		return nil
	}); err != nil {
		return err
	}
	return nil
}

func (a applicationRepositoryImpl) GetByID(ctx context.Context, id string) (core.Application, error) {
	var application core.Application
	filter := bson.M{"_id": id}
	documentReturned := a.collection.FindOne(ctx, filter)
	err := documentReturned.Decode(&application)
	log.Println("Get application document", application, err, filter)
	return application, err
}
