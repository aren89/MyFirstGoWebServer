package person

import (
	"time"
)

func ManagePersonToSave(model PersonModel) (interface{}, error) {
	p := Person{
		FirstName:  model.FirstName,
		LastName:   model.LastName,
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


