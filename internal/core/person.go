package core

import (
	"context"
	"time"
)

type PersonRepresentation struct {
	FirstName  string    `json:"firstName" binding:"required"`
	LastName   string    `json:"lastName" binding:"required"`
	Email      string    `json:"email" binding:"required,email"`
	Age        int       `json:"age" binding:"required,max=100"`
	Gender     string    `json:"gender"`
	Date       time.Time `json:"createdAt"`
	Experience string    `json:"experience"`
}

type Person struct {
	FirstName  string    `json:"firstName" binding:"required"`
	LastName   string    `json:"lastName" binding:"required"`
	Email      string    `json:"email" binding:"required,email"`
	Age        int       `json:"age" binding:"required,max=100"`
	Gender     string    `json:"gender"`
	Date       time.Time `json:"createdAt"`
	Experience string    `json:"experience"`
}


type PersonService interface {
	GetById(ctx context.Context, id string) (PersonRepresentation, error)
	Store(ctx context.Context, p *PersonRepresentation) (string, error)
}

type PersonRepository interface {
	GetByID(ctx context.Context, id string) (Person, error)
	Store(ctx context.Context, p *Person) (string, error)
}
