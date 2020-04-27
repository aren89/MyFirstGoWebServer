package service

import (
	"MyFirstGoWebServer/internal/core"
	"context"
	"time"
)

type personServiceImpl struct {
	personRepo core.PersonRepository
}

func NewPersonService(pr core.PersonRepository) core.PersonService {
	return &personServiceImpl{
		personRepo: pr,
	}
}

func (p personServiceImpl) Store(ctx context.Context, personRepresentation *core.PersonRepresentation) (string, error) {
	person := core.Person{
		FirstName:  personRepresentation.FirstName,
		LastName:   personRepresentation.LastName,
		Email:      personRepresentation.Email,
		Age:        personRepresentation.Age,
		Gender:     personRepresentation.Gender,
		Date:       time.Now(),
		Experience: "NO",
	}

	switch {
	case person.Age > 40:
		person.Experience = "HIGH"
	case person.Age > 30:
		person.Experience = "MEDIUM"
	case person.Age > 20:
		person.Experience = "LOW"
	}

	id, err := p.personRepo.Store(ctx, &person)

	return id, err
}

func (p personServiceImpl) GetById(ctx context.Context, id string) (core.PersonRepresentation, error) {
	person, err := p.personRepo.GetByID(ctx, id)
	if err != nil {
		return core.PersonRepresentation{}, err
	}
	personRepresentation := core.PersonRepresentation{
		FirstName:  person.FirstName,
		LastName:   person.LastName,
		Email:      person.Email,
		Age:        person.Age,
		Gender:     person.Gender,
		Date:       person.Date,
		Experience: person.Experience,
	}
	return personRepresentation, nil
}
