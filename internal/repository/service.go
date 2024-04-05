package repository

import (
	"context"
	"effective-mobile-test/internal/entity"
	"encoding/json"
	"fmt"
	"net/http"

	"go.uber.org/zap"
)

type serviceRepository struct {
	url    string
	logger *zap.Logger
}

func NewServiceRepository(url string, logger *zap.Logger) *serviceRepository {
	return &serviceRepository{
		url:    url,
		logger: logger,
	}
}

func (s *serviceRepository) GetCarInfo(ctx context.Context, regNum string) (*entity.Car, error) {
	url := fmt.Sprintf("%s/%s", s.url, regNum)

	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("/repository/service.GetCarInfo: failed to send GET request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("/repository/service.GetCarInfo: unexpected HTTP status code: %d", resp.StatusCode)
	}

	var car entity.Car
	if err := json.NewDecoder(resp.Body).Decode(&car); err != nil {
		return nil, fmt.Errorf("/repository/service.GetCarInfo: failed to decode response: %w", err)
	}

	s.logger.Debug("/repository/service.GetCarInfo: decoded car", zap.Any("car", car))

	return &car, nil
}
