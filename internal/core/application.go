package core

import (
	"context"
)

type ApplicationRepresentation struct {
	Id              string  `json:"id"`
	PersonId        string  `json:"personId" binding:"required"`
	JobOfferId      string  `json:"jobOfferId" binding:"required"`
	RequestedSalary float32 `json:"requestedSalary" binding:"required"`
}

type Application struct {
	Id              string  `bson:"_id,omitempty"`
	PersonId        string  `bson:"personId" binding:"required"`
	JobOfferId      string  `bson:"jobOfferId" binding:"required"`
	RequestedSalary float32 `bson:"requestedSalary" binding:"required"`
}

type ApplicationService interface {
	SaveApplication(ctx context.Context, p *ApplicationRepresentation) (string, error)
	Delete(ctx context.Context, id string) error
}

type ApplicationRepository interface {
	Store(ctx context.Context, p *Application) (string, error)
	Delete(ctx context.Context, id string) error
}
