package handlers

import "github.com/gin-gonic/gin"

//go:generate mockgen -source=./interfaces.go -destination=./handlers_mock.go -package=handlers

type CarHandlers interface {
	Create(c *gin.Context)
	GetAll(c *gin.Context)
	Update(c *gin.Context)
	Delete(c *gin.Context)
}
