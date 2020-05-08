package service

import (
	"MyFirstGoWebServer/internal/core"
	"context"
	"github.com/google/uuid"
	"log"
	"time"
)

type personServiceImpl struct {
	PersonRepo core.PersonRepository
}

func NewPersonService(pr core.PersonRepository) core.PersonService {
	return &personServiceImpl{
		PersonRepo: pr,
	}
}

func (p personServiceImpl) StoreWithEstimatedLevel(ctx context.Context, personRepresentation *core.PersonRepresentation) (string, error) {
	_uuid := uuid.New().String()
	log.Println("_uuid", _uuid)
	person := core.Person{
		Id:                       _uuid,
		FirstName:                personRepresentation.FirstName,
		LastName:                 personRepresentation.LastName,
		Email:                    personRepresentation.Email,
		Age:                      personRepresentation.Age,
		Gender:                   personRepresentation.Gender,
		Date:                     time.Now(),
		YearsOfExperienceWorking: personRepresentation.YearsOfExperienceWorking,
		EstimatedLevel:           "JUNIOR",
	}

	switch {
	case person.YearsOfExperienceWorking > 5:
		person.EstimatedLevel = "SENIOR"
	case person.YearsOfExperienceWorking > 3:
		person.EstimatedLevel = "MID"
	}

	id, err := p.PersonRepo.Store(ctx, &person)

	return id, err
}

func (p personServiceImpl) Fetch(ctx context.Context, emailFilter string, yearsOfExperienceWorkingFilter string, estimatedLevelFilter string) ([]core.PersonRepresentation, error) {
	var personsRepresentation []core.PersonRepresentation
	persons, err := p.PersonRepo.Fetch(ctx, emailFilter, yearsOfExperienceWorkingFilter, estimatedLevelFilter)
	if err != nil {
		return personsRepresentation, err
	}
	for _, person := range persons {
		personRepresentation := core.PersonRepresentation{
			Id:                       person.Id,
			FirstName:                person.FirstName,
			LastName:                 person.LastName,
			Email:                    person.Email,
			Age:                      person.Age,
			Gender:                   person.Gender,
			YearsOfExperienceWorking: person.YearsOfExperienceWorking,
		}
		personsRepresentation = append(personsRepresentation, personRepresentation)
	}
	return personsRepresentation, err
}

func (p personServiceImpl) GetById(ctx context.Context, id string) (core.PersonRepresentation, error) {
	person, err := p.PersonRepo.GetByID(ctx, id)
	if err != nil {
		return core.PersonRepresentation{}, err
	}
	personRepresentation := core.PersonRepresentation{
		Id:                       person.Id,
		FirstName:                person.FirstName,
		LastName:                 person.LastName,
		Email:                    person.Email,
		Age:                      person.Age,
		Gender:                   person.Gender,
		YearsOfExperienceWorking: person.YearsOfExperienceWorking,
	}
	return personRepresentation, nil
}
