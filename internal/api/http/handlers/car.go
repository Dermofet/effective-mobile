package handlers

import (
	"context"
	"effective-mobile-test/internal/entity"
	"effective-mobile-test/internal/usecase"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	_ "github.com/swaggo/swag"
	"go.uber.org/zap"
)

type carHandlers struct {
	i      usecase.CarInteractor
	logger *zap.Logger
}

func NewCarHandlers(i usecase.CarInteractor, logger *zap.Logger) *carHandlers {
	return &carHandlers{
		i:      i,
		logger: logger,
	}
}

// Create godoc
// @Summary Add new car
// @Description Add new car from a list of registration numbers
// @Tags Car
// @Accept json
// @Param request body entity.RegNums true "List of registration numbers"
// @Success 201 "Car added"
// @Failure 400 "Invalid request body or parameter"
// @Failure 500 "Server error"
// @Router /car/new [post]
func (h *carHandlers) Create(c *gin.Context) {
	regNums := entity.RegNums{}
	h.logger.Debug("/handlers/car.Create: decoding request body")
	if err := json.NewDecoder(c.Request.Body).Decode(&regNums); err != nil {
		h.logger.Error("/handlers/car.Create: can't unmarshal body", zap.Error(err))
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	ctx := context.Background()
	h.logger.Debug("/handlers/car.Create: creating car")
	if err := h.i.Create(ctx, &regNums); err != nil {
		h.logger.Error("/handlers/car.Create: can't create car", zap.Error(err))
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	c.Status(http.StatusCreated)
}

// GetAll godoc
// @Summary Get all cars
// @Description Get all cars
// @Tags Car
// @Produce json
// @Param limit query int true "Limit of cars in response" default(100)
// @Param offset query int true "Offset of cars in response" default(0)
// @Success 200 "OK"
// @Failure 400 "Invalid request body or parameter"
// @Failure 500 "Server error"
// @Router /car/all [get]
func (h *carHandlers) GetAll(c *gin.Context) {
	ctx := context.Background()

	h.logger.Debug("/handlers/car.GetAll: get limit")
	limit, err := strconv.Atoi(c.Query("limit"))
	if err != nil {
		h.logger.Error("/handlers/car.GetAll: can't parse limit", zap.Error(err))
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	h.logger.Debug("/handlers/car.GetAll: get offset")
	offset, err := strconv.Atoi(c.Query("offset"))
	if err != nil {
		h.logger.Error("/handlers/car.GetAll: can't parse offset", zap.Error(err))
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	h.logger.Debug("/handlers/car.GetAll: get cars")
	cars, err := h.i.GetAll(ctx, limit, offset)
	if err != nil {
		h.logger.Error("/handlers/car.GetAll: can't get cars", zap.Error(err))
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, cars)
}

// Update godoc
// @Summary Update car
// @Description Update car by id
// @Tags Car
// @Accept json
// @Param id path string true "Car id"
// @Param request body entity.Car true "Car data"
// @Success 200 "OK"
// @Failure 400 "Invalid request body or parameter"
// @Failure 500 "Server error"
// @Router /car/update/{id} [put]
func (h *carHandlers) Update(c *gin.Context) {
	h.logger.Debug("/handlers/car.Update: get id")
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		h.logger.Error("/handlers/car.Update: can't parse id", zap.Error(err))
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	h.logger.Debug("/handlers/car.Update: decoding request body")
	car := entity.CarUpdate{}
	if err := json.NewDecoder(c.Request.Body).Decode(&car); err != nil {
		h.logger.Error("/handlers/car.Update: can't unmarshal body", zap.Error(err))
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	ctx := context.Background()

	h.logger.Debug("/handlers/car.Update: update car")
	if err := h.i.Update(ctx, id, &car); err != nil {
		h.logger.Error("/handlers/car.Update: can't update car", zap.Error(err))
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	h.logger.Debug("/handlers/car.Update: finished successfully")

	c.Status(http.StatusOK)
}

// Delete godoc
// @Summary Delete car
// @Description Delete car by id
// @Tags Car
// @Param id path string true "Car id"
// @Success 204 "Car deleted"
// @Failure 400 "Invalid request body or parameter"
// @Failure 500 "Server error"
// @Router /car/delete/{id} [delete]
func (h *carHandlers) Delete(c *gin.Context) {
	h.logger.Debug("/handlers/car.Delete: get id")
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		h.logger.Error("/handlers/car.Delete: can't parse id", zap.Error(err))
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	ctx := context.Background()

	h.logger.Debug("/handlers/car.Delete: delete car")
	if err := h.i.Delete(ctx, id); err != nil {
		h.logger.Error("/handlers/car.Delete: can't delete car", zap.Error(err))
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	c.Status(http.StatusNoContent)
}
