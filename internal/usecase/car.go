package usecase

import (
	"context"
	"effective-mobile-test/internal/entity"
	"effective-mobile-test/internal/repository"
	"fmt"
	"sync"

	"github.com/google/uuid"
	"go.uber.org/zap"
)

type carInteractor struct {
	logger      *zap.Logger
	carRepo     repository.CarRepository
	serviceRepo repository.ServiceRepository
}

func NewCarInteractor(carRepo repository.CarRepository, serviceRepo repository.ServiceRepository, logger *zap.Logger) *carInteractor {
	return &carInteractor{
		carRepo:     carRepo,
		serviceRepo: serviceRepo,
		logger:      logger,
	}
}

func (c *carInteractor) Create(ctx context.Context, regNums *entity.RegNums) error {
	c.logger.Debug("/usecase/car.Create: started")

	var wg sync.WaitGroup
	wg.Add(len(regNums.Nums))

	var errors []error
	errorsLock := sync.Mutex{}

	for _, num := range regNums.Nums {
		go func(num string) {
			defer wg.Done()

			c.logger.Debug("/usecase/car.Create: getting car info for registration number", zap.String("reg_num", num))

			car, err := c.serviceRepo.GetCarInfo(ctx, num)
			if err != nil {
				errorsLock.Lock()
				defer errorsLock.Unlock()
				errors = append(errors, fmt.Errorf("/usecase/car.Create: can't get car info: %w", err))
				return
			}

			c.logger.Debug("/usecase/car.Create: creating car", zap.String("reg_num", num))

			if err := c.carRepo.Create(ctx, car); err != nil {
				errorsLock.Lock()
				defer errorsLock.Unlock()
				errors = append(errors, fmt.Errorf("/usecase/car.Create: can't create car: %w", err))
			}
		}(num)
	}

	wg.Wait()

	if len(errors) > 0 {
		return fmt.Errorf("errors occurred during car creation: %v", errors)
	}

	c.logger.Debug("/usecase/car.Create: finished successfully")

	return nil
}

func (c *carInteractor) GetAll(ctx context.Context, limit, offset int) ([]*entity.Car, error) {
	c.logger.Debug("/usecase/car.GetAll: started")

	cars, err := c.carRepo.GetAll(ctx, limit, offset)
	if err != nil {
		return nil, err
	}

	c.logger.Debug("/usecase/car.GetAll: finished successfully")

	return cars, nil
}

func (c *carInteractor) Update(ctx context.Context, id uuid.UUID, car *entity.CarUpdate) error {
	c.logger.Debug("/usecase/car.Update: started")

	if err := c.carRepo.Update(ctx, id, car); err != nil {
		return err
	}

	c.logger.Debug("/usecase/car.Update: finished successfully")

	return nil
}

func (c *carInteractor) Delete(ctx context.Context, id uuid.UUID) error {
	c.logger.Debug("/usecase/car.Delete: started")

	if err := c.carRepo.Delete(ctx, id); err != nil {
		return err
	}

	c.logger.Debug("/usecase/car.Delete: finished successfully")

	return nil
}
