package person

import (
	"time"
)

func ManagePersonToSave(model PersonModel) (interface{}, error) {
	p := Person{
		FirstName:  model.FirstName,
		LastName:   model.LastName,
		Email:      model.Email,
		Age:        model.Age,
		Gender:     model.Gender,
		Date:       time.Now(),
		Experience: "NO",
	}

	switch {
	case p.Age > 40:
		p.Experience = "HIGH"
	case p.Age > 30:
		p.Experience = "MEDIUM"
	case p.Age > 20:
		p.Experience = "LOW"
	}

	id, err := PersistPerson(p)

	return id, err
}

func ManagePersonToGet(id string) PersonModel {
	person := RetrievePersonFromDB(id)
	p := PersonModel{
		FirstName:  person.FirstName,
		LastName:   person.LastName,
		Email:      person.Email,
		Age:        person.Age,
		Gender:     person.Gender,
		Date:       person.Date,
		Experience: person.Experience,
	}
	return p
}
