package http

import (
	"effective-mobile-test/docs"
	"effective-mobile-test/internal/api/http/handlers"
	"effective-mobile-test/internal/db"
	"effective-mobile-test/internal/repository"
	"effective-mobile-test/internal/usecase"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"go.uber.org/zap"

	"github.com/gin-contrib/cors"
)

type routerHandlers struct {
	carHandlers handlers.CarHandlers
}

type router struct {
	router     *gin.Engine
	db         *sqlx.DB
	handlers   routerHandlers
	logger     *zap.Logger
	addr       string
	serviceUrl string // for external service
}

func NewRouter(db *sqlx.DB, logger *zap.Logger, addr string, serviceUrl string) *router {
	return &router{
		router:     gin.New(),
		db:         db,
		logger:     logger,
		addr:       addr,
		serviceUrl: serviceUrl,
	}
}

func (r *router) Init() error {
	r.router.Use(
		gin.Logger(),
		gin.CustomRecovery(r.recovery),
	)
	err := r.registerRoutes()
	if err != nil {
		return fmt.Errorf("can't init router: %w", err)
	}

	return nil
}

func (r *router) recovery(c *gin.Context, recovered any) {
	defer func() {
		if e := recover(); e != nil {
			r.logger.Fatal("http server panic", zap.Error(fmt.Errorf("%s", recovered)))
		}
	}()
	c.AbortWithStatus(http.StatusInternalServerError)
}

func (r *router) registerRoutes() error {
	r.router.NoMethod(handlers.NotImplementedHandler)
	r.router.NoRoute(handlers.NotImplementedHandler)

	corsMiddleware := cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"Authorization", "Content-Type"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	})

	r.router.Use(corsMiddleware)

	docs.SwaggerInfo.BasePath = "/"
	r.router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	pgSource := db.NewSource(r.db, r.logger)
	carRepository := repository.NewCarRepository(pgSource, r.logger)
	serviceRepository := repository.NewServiceRepository(r.serviceUrl, r.logger)
	carInteractor := usecase.NewCarInteractor(carRepository, serviceRepository, r.logger)

	r.handlers.carHandlers = handlers.NewCarHandlers(carInteractor, r.logger)

	r.router.POST("/car/new", r.handlers.carHandlers.Create)
	r.router.GET("/car/all", r.handlers.carHandlers.GetAll)
	r.router.PUT("/car/update/:id", r.handlers.carHandlers.Update)
	r.router.DELETE("/car/delete/:id", r.handlers.carHandlers.Delete)

	return nil
}
