package http

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
)

//go:generate mockgen -source=server.go -destination=./server_mock.go -package=http

const RequestTimeOut = 30 * time.Second

type Server interface {
	Run(ctx context.Context) error
	Shutdown(ctx context.Context) error
}

type server struct {
	server *http.Server
	db     *sqlx.DB
	logger *zap.Logger
}

func NewServer(
	addr string,
	db *sqlx.DB,
	logger *zap.Logger,
	serviceUrl string, // url to service
) *server {
	s := &server{
		db:     db,
		logger: logger,
	}

	r := NewRouter(db, logger, addr, serviceUrl)
	err := r.Init()
	if err != nil {
		s.logger.Error("can't init router:", zap.Error(err))
		return nil
	}

	httpServer := &http.Server{
		Addr:              addr,
		Handler:           r.router,
		ReadHeaderTimeout: RequestTimeOut,
	}
	s.server = httpServer

	return s
}

func (s *server) Run(ctx context.Context) error {
	go func() {
		<-ctx.Done()
		err := s.server.Shutdown(ctx)
		if err != nil {
			s.logger.Error("can't shutdown http-server", zap.Error(err))
			return
		}
	}()

	return s.server.ListenAndServe()
}

func (s *server) Shutdown(ctx context.Context) error {
	err := s.server.Shutdown(ctx)
	if err != nil {
		return fmt.Errorf("http server shutdown error: %w", err)
	}
	return nil
}
