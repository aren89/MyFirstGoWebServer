package repository

import (
	"MyFirstGoWebServer/internal/core"
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
)

type jobOfferRepositoryImpl struct {
	db *mongo.Database
}

func NewJobOfferRepository(db *mongo.Database) core.JobOfferRepository {
	return &jobOfferRepositoryImpl{
		db: db,
	}
}

func (j jobOfferRepositoryImpl) GetByID(ctx context.Context, id string) (core.JobOffer, error) {
	panic("implement me")
}

func (j jobOfferRepositoryImpl) Store(ctx context.Context, p *core.JobOffer) error {
	collection := j.db.Collection("jobOffers")
	filter := bson.M{"_id": p.Id}
	if result := collection.FindOneAndReplace(ctx, filter, p); result != nil {
		_, err := collection.InsertOne(ctx, p)
		if err != nil {
			log.Println("Error on inserting new job offer", err)
			return err
		}
	}
	return nil
}

func (j jobOfferRepositoryImpl) Fetch(ctx context.Context, roleFilter string, companyFilter string) ([]*core.JobOffer, error) {
	panic("implement me")
}
