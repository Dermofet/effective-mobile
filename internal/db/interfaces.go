package db

import (
	"context"
	"effective-mobile-test/internal/entity"

	"github.com/google/uuid"
)

//go:generate mockgen -source=interfaces.go -destination=./source_mock.go -package=db
type CarSource interface {
	CreateCar(ctx context.Context, car *entity.Car) error
	GetAllCars(ctx context.Context, limit, offset int) ([]*entity.Car, error)
	UpdateCar(ctx context.Context, id uuid.UUID, car *entity.CarUpdate) error
	DeleteCar(ctx context.Context, id uuid.UUID) error
}
