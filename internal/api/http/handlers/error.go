package handlers

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

func NotImplementedHandler(c *gin.Context) {
	c.AbortWithStatus(http.StatusMethodNotAllowed)
}

var (
	ErrReadBody = errors.New("can't read body")
)