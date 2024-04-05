package repository

import (
	"context"
	"effective-mobile-test/internal/entity"

	"github.com/google/uuid"
)

//go:generate mockgen -source=interfaces.go -destination=./repositories_mock.go -package=repository
type CarRepository interface {
	Create(ctx context.Context, carCreate *entity.Car) error
	GetAll(ctx context.Context, limit, offset int) ([]*entity.Car, error)
	Update(ctx context.Context, id uuid.UUID, carUpdate *entity.CarUpdate) error
	Delete(ctx context.Context, id uuid.UUID) error
}

type ServiceRepository interface {
	GetCarInfo(ctx context.Context, regNum string) (*entity.Car, error)
}
