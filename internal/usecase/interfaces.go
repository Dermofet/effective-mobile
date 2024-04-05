package usecase

import (
	"context"
	"effective-mobile-test/internal/entity"

	"github.com/google/uuid"
)

//go:generate mockgen -source=./interfaces.go -destination=./usecases_mock.go -package=usecase

type CarInteractor interface {
	Create(ctx context.Context, regNums *entity.RegNums) error
	GetAll(ctx context.Context, limit, offset int) ([]*entity.Car, error)
	Update(ctx context.Context, id uuid.UUID, car *entity.CarUpdate) error
	Delete(ctx context.Context, id uuid.UUID) error
}
