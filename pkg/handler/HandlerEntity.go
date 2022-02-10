package handler

import (
	"todo/pkg/service"
)

type Handler struct {
	services *service.Service
}

type errorResponse struct {
	Message string `json:"message"`
}
