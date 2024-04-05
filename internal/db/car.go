package db

import (
	"context"
	"effective-mobile-test/internal/entity"
	"fmt"
	"strings"

	"github.com/google/uuid"
)

func (s *source) CreateCar(ctx context.Context, car *entity.Car) error {
	s.logger.Debug("/db/car.CreateCar: started")

	dbCtx, dbCancel := context.WithTimeout(ctx, QueryTimeout)
	defer dbCancel()

	row := s.db.QueryRowContext(
		dbCtx,
		`INSERT INTO cars (id, reg_num, mark, model, year, owner_name, owner_surname, owner_patronymic) 
        VALUES ($1, $2, $3, $4, $5, $6, $7, $8)`,
		uuid.New(), car.RegNum, car.Mark, car.Model, car.Year, car.Owner.Name, car.Owner.Surname, car.Owner.Patronymic,
	)
	if err := row.Err(); err != nil {
		return fmt.Errorf("/db/car.CreateCar: %w", err)
	}

	s.logger.Debug("/db/car.CreateCar: finished successfully")

	return nil
}

func (s *source) GetAllCars(ctx context.Context, limit, offset int) ([]*entity.Car, error) {
	s.logger.Debug("/db/car.GetAllCars: started")

	dbCtx, dbCancel := context.WithTimeout(ctx, QueryTimeout)
	defer dbCancel()

	rows, err := s.db.QueryxContext(
		dbCtx,
		`SELECT id, reg_num, mark, model, year, owner_name, owner_surname, owner_patronymic 
        FROM cars LIMIT $1 OFFSET $2`, limit, offset,
	)
	if err != nil {
		return nil, fmt.Errorf("/db/car.GetAllCars: %w", err)
	}
	defer rows.Close()

	cars := make([]*entity.Car, 0)
	for rows.Next() {
		car := &entity.Car{
			Owner: &entity.Owner{},
		}

		if err := rows.Scan(
			&car.ID,
			&car.RegNum,
			&car.Mark,
			&car.Model,
			&car.Year,
			&car.Owner.Name,
			&car.Owner.Surname,
			&car.Owner.Patronymic,
		); err != nil {
			return nil, fmt.Errorf("/db/car.GetAllCars: %w", err)
		}
		cars = append(cars, car)
	}

	s.logger.Debug("/db/car.GetAllCars: finished successfully")

	return cars, nil
}

func (s *source) UpdateCar(ctx context.Context, id uuid.UUID, car *entity.CarUpdate) error {
	s.logger.Debug("/db/car.UpdateCar: started")

	dbCtx, dbCancel := context.WithTimeout(ctx, QueryTimeout)
	defer dbCancel()

	query := `UPDATE cars SET `
	args := []interface{}{}

	i := 1
	if car.RegNum != nil {
		query += fmt.Sprintf("reg_num = $%d, ", i)
		args = append(args, *car.RegNum)
		i++
	}

	if car.Mark != nil {
		query += fmt.Sprintf("mark = $%d, ", i)
		args = append(args, *car.Mark)
		i++
	}

	if car.Model != nil {
		query += fmt.Sprintf("model = $%d, ", i)
		args = append(args, *car.Model)
		i++
	}

	if car.Year != nil {
		query += fmt.Sprintf("year = $%d, ", i)
		args = append(args, *car.Year)
		i++
	}

	if car.Owner != nil {
		if car.Owner.Name != nil {
			query += fmt.Sprintf("owner_name = $%d, ", i)
			args = append(args, *car.Owner.Name)
			i++
		}
		if car.Owner.Surname != nil {
			query += fmt.Sprintf("owner_surname = $%d, ", i)
			args = append(args, *car.Owner.Surname)
			i++
		}
		if car.Owner.Patronymic != nil {
			query += fmt.Sprintf("owner_patronymic = $%d, ", i)
			args = append(args, *car.Owner.Patronymic)
			i++
		}
	}

	query = fmt.Sprintf("%s WHERE id = $%d", strings.TrimSuffix(query, ", "), i)

	args = append(args, id)

	row := s.db.QueryRowContext(dbCtx, query, args...)

	if err := row.Err(); err != nil {
		return fmt.Errorf("/db/car.UpdateCar: %w", err)
	}

	s.logger.Debug("/db/car.UpdateCar: finished successfully")

	return nil
}

func (s *source) DeleteCar(ctx context.Context, id uuid.UUID) error {
	s.logger.Debug("/db/car.DeleteCar: started")

	dbCtx, dbCancel := context.WithTimeout(ctx, QueryTimeout)
	defer dbCancel()

	row := s.db.QueryRowContext(
		dbCtx,
		"DELETE FROM cars WHERE id = $1",
		id,
	)
	if err := row.Err(); err != nil {
		return fmt.Errorf("/db/car.DeleteCar: %w", err)
	}

	s.logger.Debug("/db/car.DeleteCar: finished successfully")

	return nil
}
