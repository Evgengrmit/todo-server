package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"todo/pkg/service"
	"todo/user"
)

func NewHandler(serv *service.Service) *Handler {
	return &Handler{services: serv}
}

// AUTH

func (h *Handler) SignUp(c *gin.Context) {
	var input user.User
	if err := c.BindJSON(&input); err != nil {
		NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	id, err := h.services.Authorization.CreateUser(input)
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusCreated, &gin.H{"id": id})

}
func (h *Handler) SignIn(c *gin.Context) {
	var input signInInput
	if err := c.BindJSON(&input); err != nil {
		NewErrorResponse(c, http.StatusBadRequest, err.Error()+"jhfshfdkhdshdkdahd")
		return
	}
	token, err := h.services.Authorization.GenerateToken(input.Login, input.Password)
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusAccepted, &gin.H{"token": token})
}
