package app

import (
	"context"
	"effective-mobile-test/cmd/effective-mobile-test/config"
	"effective-mobile-test/internal/api/http"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
)

type app struct {
	config     *config.Config
	dbConn     *sqlx.DB
	logger     *zap.Logger
	httpServer http.Server
}

func NewApp(cfg *config.Config, logger *zap.Logger) *app {
	return &app{
		config: cfg,
		logger: logger,
	}
}

func (a *app) Start(ctx context.Context) {
	appCtx, cancelApp := context.WithCancel(ctx)
	logger := a.logger
	defer func() {
		if e := recover(); e != nil {
			logger.Panic("application shutdown", zap.Error(fmt.Errorf("%s", e)))
			cancelApp()
		}
	}()
	// Инициализируем БД
	dbConn, err := a.initDb(appCtx,
		a.config.DB.Host,
		a.config.DB.Port,
		a.config.DB.Name,
		a.config.DB.Username,
		a.config.DB.Password,
		a.config.DB.SSLMode,
	)
	if err != nil {
		logger.Fatal("init db error", zap.Error(err))
	}
	logger.Debug("init db success")
	a.dbConn = dbConn

	// Запуск миграций
	err = a.startMigrate(appCtx, migrationsPath, a.config.DB.Name, a.dbConn)
	if err != nil {
		logger.Error("db migration error", zap.Error(err))
	}
	logger.Debug("db migration success")

	defer func() {
		if e := recover(); e != nil {
			logger.Panic("http start panic", zap.Error(fmt.Errorf("%s", e)))
		}
	}()

	if a.config.Debug {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}

	addr := fmt.Sprintf("%s:%d", a.config.HttpServer.Host, a.config.HttpServer.Port)
	a.httpServer = http.NewServer(addr, a.dbConn, logger, a.config.ServiceURL)
	if a.httpServer == nil {
		cancelApp()
		logger.Fatal("can't create http server")
		return
	}
	err = a.httpServer.Run(appCtx)

	cancelApp()
	if err != nil {
		logger.Error("can't start http server", zap.Error(err))
		return
	}
}

// GracefulShutdown graceful shutdown приложения
func (a *app) GracefulShutdown(ctx context.Context) error {
	err := a.httpServer.Shutdown(ctx)
	if err != nil {
		return fmt.Errorf("can't shutdown http-server: %w", err)
	}
	err = a.dbConn.Close()
	if err != nil {
		return fmt.Errorf("can't shutdown db: %w", err)
	}
	return nil
}

// initDb инициализация базы данных
func (a *app) initDb(
	ctx context.Context,
	host string,
	port int,
	dbName string,
	user string,
	password string,
	sslmode string,
) (*sqlx.DB, error) {
	db, err := sqlx.ConnectContext(
		ctx,
		"postgres",
		fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=%s", host, port, user, password, dbName, sslmode),
	)
	if err != nil {
		return nil, err
	}

	return db, nil
}
