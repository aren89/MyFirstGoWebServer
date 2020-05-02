package core

import (
	"context"
	"time"
)

type PersonRepresentation struct {
	FirstName                string  `json:"firstName" binding:"required"`
	LastName                 string  `json:"lastName" binding:"required"`
	Email                    string  `json:"email" binding:"required,email"`
	Age                      int     `json:"age" binding:"required,max=100"`
	Gender                   string  `json:"gender"`
	YearsOfExperienceWorking float32 `json:"yearsOfExperienceWorking" binding:"required"`
}

type Person struct {
	FirstName                string    `bson:"firstName" binding:"required"`
	LastName                 string    `bson:"lastName" binding:"required"`
	Email                    string    `bson:"email" binding:"required,email"`
	Age                      int       `bson:"age" binding:"required,max=100"`
	Date                     time.Time `bson:"createdAt" binding:"required"`
	EstimatedLevel           string    `bson:"estimatedLevel" binding:"required"`
	Gender                   string    `bson:"gender"`
	YearsOfExperienceWorking float32   `bson:"yearsOfExperienceWorking" binding:"required"`
}

type PersonService interface {
	GetById(ctx context.Context, id string) (PersonRepresentation, error)
	StoreWithEstimatedLevel(ctx context.Context, p *PersonRepresentation) (string, error)
	Fetch(ctx context.Context, emailFilter string, yearsOfExperienceWorkingFilter string, estimatedLevelFilter string) ([]PersonRepresentation, error)
}

type PersonRepository interface {
	GetByID(ctx context.Context, id string) (Person, error)
	Store(ctx context.Context, p *Person) (string, error)
	Fetch(ctx context.Context, emailFilter string, yearsOfExperienceWorkingFilter string, estimatedLevelFilter string) ([]*Person, error)
}
