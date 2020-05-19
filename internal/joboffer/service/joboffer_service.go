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
	jobOffer, err := j.JobOfferRepo.GetByID(ctx, id)
	if err != nil {
		return core.JobOfferRepresentation{}, err
	}
	jobOfferRepresentation := core.JobOfferRepresentation{
		Id:           jobOffer.Id,
		Company:      jobOffer.Company,
		Description:  jobOffer.Description,
		Role:         jobOffer.Role,
		Applications: jobOffer.Applications,
	}
	return jobOfferRepresentation, nil
}

func (j jobOfferServiceImpl) SaveKafkaMessage(ctx context.Context, id string, p *core.JobOfferRepresentation) error {
	jobOffer := core.JobOffer{
		Id:           id,
		Company:      p.Company,
		Description:  p.Description,
		Role:         p.Role,
		Applications: []string{},
	}
	err := j.JobOfferRepo.Store(ctx, &jobOffer)
	return err
}

func (j jobOfferServiceImpl) Fetch(ctx context.Context, roleFilter string, companyFilter string) ([]core.JobOfferRepresentation, error) {
	var jobOffersRepresentation []core.JobOfferRepresentation
	jobOffers, err := j.JobOfferRepo.Fetch(ctx, roleFilter, companyFilter)
	if err != nil {
		return jobOffersRepresentation, err
	}
	for _, jobOffer := range jobOffers {
		jobOfferRepresentation := core.JobOfferRepresentation{
			Id:           jobOffer.Id,
			Company:      jobOffer.Company,
			Description:  jobOffer.Description,
			Role:         jobOffer.Role,
			Applications: jobOffer.Applications,
		}
		jobOffersRepresentation = append(jobOffersRepresentation, jobOfferRepresentation)
	}
	return jobOffersRepresentation, err
}
