package core

import (
	"context"
)

type JobOfferRepresentation struct {
	Company      string   `json:"company" binding:"required"`
	Description  string   `json:"description" binding:"required"`
	Role         string   `json:"role" binding:"required"`
	Applications []string `json:"applications"`
}

type JobOffer struct {
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
	SaveKafkaMessage(ctx context.Context, p *JobOfferRepresentation) error
	Fetch(ctx context.Context, emailFilter string, yearsOfExperienceWorkingFilter string, estimatedLevelFilter string) ([]JobOfferRepresentation, error)
}

type JobOfferRepository interface {
	GetByID(ctx context.Context, id string) (JobOffer, error)
	Store(ctx context.Context, p *JobOffer) error
	Fetch(ctx context.Context, roleFilter string, companyFilter string) ([]*JobOffer, error)
}
