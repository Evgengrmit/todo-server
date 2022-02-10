package handler

import (
	"todo/pkg/service"
)

type Handler struct {
	services *service.Service
}

type error struct {
	Message string `json:"message"`
}
