package repository

import (
	"MyFirstGoWebServer/internal/core"
	"context"
	"go.mongodb.org/mongo-driver/mongo"
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
	panic("implement me")
}

func (j jobOfferRepositoryImpl) Fetch(ctx context.Context, roleFilter string, companyFilter string) ([]*core.JobOffer, error) {
	panic("implement me")
}
