package repository

import (
	"context"
	"effective-mobile-test/internal/db"
	"effective-mobile-test/internal/entity"
	"fmt"

	"github.com/google/uuid"
	"go.uber.org/zap"
)

type carRepository struct {
	logger *zap.Logger
	source db.CarSource
}

func NewCarRepository(source db.CarSource, logger *zap.Logger) *carRepository {
	return &carRepository{
		source: source,
		logger: logger,
	}
}

func (r *carRepository) Create(ctx context.Context, carCreate *entity.Car) error {
	// Логирование начала выполнения функции
	r.logger.Debug("/repository/car.Create: started")

	if err := r.source.CreateCar(ctx, carCreate); err != nil {
		// Логирование ошибки
		r.logger.Error("/repository/car.Create: can't create car", zap.Error(err))
		return fmt.Errorf("/repository/car.Create: %w", err)
	}

	// Логирование успешного выполнения
	r.logger.Debug("/repository/car.Create: finished successfully")

	return nil
}

func (r *carRepository) GetAll(ctx context.Context, limit, offset int) ([]*entity.Car, error) {
	// Логирование начала выполнения функции
	r.logger.Debug("/repository/car.GetAll: started")

	// Вызов метода получения всех машин из источника данных
	cars, err := r.source.GetAllCars(ctx, limit, offset)
	if err != nil {
		// Логирование ошибки получения машин
		r.logger.Error("/repository/car.GetAll: can't get cars", zap.Error(err))
		return nil, fmt.Errorf("/repository/car.GetAll: %w", err)
	}

	// Логирование успешного завершения функции
	r.logger.Debug("/repository/car.GetAll: finished successfully")

	return cars, nil
}

func (r *carRepository) Update(ctx context.Context, id uuid.UUID, carUpdate *entity.CarUpdate) error {
	// Логирование начала выполнения функции
	r.logger.Debug("/repository/car.Update: started")

	// Вызов метода обновления машины в источнике данных
	if err := r.source.UpdateCar(ctx, id, carUpdate); err != nil {
		// Логирование ошибки обновления машины
		r.logger.Error("/repository/car.Update: can't update car", zap.Error(err))
		return fmt.Errorf("/repository/car.Update: %w", err)
	}

	// Логирование успешного завершения функции
	r.logger.Debug("/repository/car.Update: finished successfully")

	return nil
}

func (r *carRepository) Delete(ctx context.Context, id uuid.UUID) error {
	// Логирование начала выполнения функции
	r.logger.Debug("/repository/car.Delete: started")

	// Вызов метода удаления машины из источника данных
	if err := r.source.DeleteCar(ctx, id); err != nil {
		// Логирование ошибки удаления машины
		r.logger.Error("/repository/car.Delete: can't delete car", zap.Error(err))
		return fmt.Errorf("/repository/car.Delete: %w", err)
	}

	// Логирование успешного завершения функции
	r.logger.Debug("/repository/car.Delete: finished successfully")

	return nil
}
