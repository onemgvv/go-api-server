package http

import (
	"github.com/gin-gonic/gin"
	"github.com/onemgvv/go-api-server/internal/logger"
)

type Error struct {
	StatusCode int    `json:"int"`
	Message    string `json:"message"`
}

func newErrorResponse(c *gin.Context, statusCode int, message string) {
	logger.ErrorLogger.Errorf("[ERROR FROM RESPONSE]: %s", message)
	c.AbortWithStatusJSON(statusCode, Error{statusCode, message})
}
