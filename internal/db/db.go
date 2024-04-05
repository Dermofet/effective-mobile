package db

import (
	"time"

	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
)

const (
	QueryTimeout = 10 * time.Second
)

type source struct {
	logger *zap.Logger
	db     *sqlx.DB
}

func NewSource(db *sqlx.DB, logger *zap.Logger) *source {
	return &source{
		db:     db,
		logger: logger,
	}
}
