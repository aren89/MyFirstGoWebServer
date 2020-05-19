package service

import (
	"MyFirstGoWebServer/internal/core"
	"context"
	"errors"
	"github.com/google/uuid"
)

type applicationServiceImpl struct {
	ApplicationRepo core.ApplicationRepository
	PersonService   core.PersonService
}

func NewApplicationService(ar core.ApplicationRepository, pr core.PersonService) core.ApplicationService {
	return &applicationServiceImpl{
		ApplicationRepo: ar,
		PersonService:   pr,
	}
}

func (a applicationServiceImpl) SaveApplication(ctx context.Context, p *core.ApplicationRepresentation) (string, error) {
	personRepresentation, err := a.PersonService.GetById(ctx, p.PersonId)
	if (core.PersonRepresentation{} == personRepresentation) {
		return "", errors.New("no person found")
	}
	if err != nil {
		return "", err
	}
	_uuid := uuid.New().String()
	app := core.Application{
		Id:              _uuid,
		PersonId:        p.PersonId,
		JobOfferId:      p.JobOfferId,
		RequestedSalary: p.RequestedSalary,
	}

	if p.RequestedSalary > 100000 {
		return "", errors.New("max requested salary too high")
	}

	return a.ApplicationRepo.Store(ctx, &app)
}

func (a applicationServiceImpl) Delete(ctx context.Context, id string) error {
	return a.ApplicationRepo.Delete(ctx, id)
}

func (a applicationServiceImpl) GetById(ctx context.Context, id string) (core.ApplicationRepresentation, error) {
	application, err := a.ApplicationRepo.GetByID(ctx, id)
	if err != nil {
		return core.ApplicationRepresentation{}, err
	}
	applicationRepresentation := core.ApplicationRepresentation{
		Id:              application.Id,
		PersonId:        application.PersonId,
		JobOfferId:      application.JobOfferId,
		RequestedSalary: application.RequestedSalary,
	}
	return applicationRepresentation, nil
}
