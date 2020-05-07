package service

import (
	"MyFirstGoWebServer/internal/core"
	"context"
)

type jobOfferServiceImpl struct {
	JobOfferRepo core.JobOfferRepository
}

func NewJobOfferService(jr core.JobOfferRepository) core.JobOfferService {
	return &jobOfferServiceImpl{
		JobOfferRepo: jr,
	}
}

func (j jobOfferServiceImpl) GetById(ctx context.Context, id string) (core.JobOfferRepresentation, error) {
	panic("implement me")
}

func (j jobOfferServiceImpl) SaveKafkaMessage(ctx context.Context, p *core.JobOfferRepresentation) error {
	panic("implement me")
}

func (j jobOfferServiceImpl) Fetch(ctx context.Context, emailFilter string, yearsOfExperienceWorkingFilter string, estimatedLevelFilter string) ([]core.JobOfferRepresentation, error) {
	panic("implement me")
}
