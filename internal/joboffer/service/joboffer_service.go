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

func (j jobOfferServiceImpl) SaveKafkaMessage(ctx context.Context, id string, p *core.JobOfferRepresentation) error {
	jobOffer := core.JobOffer{
		Id:           id,
		Company:      p.Company,
		Description:  p.Description,
		Role:         p.Role,
	}
	err := j.JobOfferRepo.Store(ctx, &jobOffer)
	return err
}

func (j jobOfferServiceImpl) Fetch(ctx context.Context, roleFilter string, companyFilter string) ([]core.JobOfferRepresentation, error) {
	panic("implement me")
}
