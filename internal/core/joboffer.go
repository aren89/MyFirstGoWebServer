package core

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
)

type JobOfferRepresentation struct {
	Id           string   `json:"id"`
	Company      string   `json:"company" binding:"required"`
	Description  string   `json:"description" binding:"required"`
	Role         string   `json:"role" binding:"required"`
	Applications []string `json:"applications"`
}

type JobOffer struct {
	Id           string   `bson:"_id,omitempty"`
	Company      string   `bson:"company" binding:"required"`
	Description  string   `bson:"description" binding:"required"`
	Role         string   `bson:"role" binding:"required"`
	Applications []string `bson:"applications"`
}

type JobOfferConsumer interface {
	HandleMessage(key []byte, value []byte)
}

type JobOfferService interface {
	GetById(ctx context.Context, id string) (JobOfferRepresentation, error)
	SaveKafkaMessage(ctx context.Context, id string, p *JobOfferRepresentation) error
	Fetch(ctx context.Context, roleFilter string, companyFilter string) ([]JobOfferRepresentation, error)
}

type JobOfferRepository interface {
	GetByID(ctx context.Context, id string) (JobOffer, error)
	Store(ctx context.Context, p *JobOffer) error
	Fetch(ctx context.Context, roleFilter string, companyFilter string) ([]*JobOffer, error)
	PushApplication(ctx mongo.SessionContext, id string, applicationId string) error
	PopApplication(ctx mongo.SessionContext, applicationId string) error
}
