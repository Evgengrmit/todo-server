package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"todo/pkg/service"
)

type Handler struct {
	services *service.Service
}

type errorResponse struct {
	Message string `json:"message"`
}

func NewErrorResponse(c *gin.Context, statusCode int, message string) {
	logrus.Error(message)
	c.AbortWithStatusJSON(statusCode, errorResponse{message})
}
