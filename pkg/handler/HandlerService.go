package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
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
	c.JSON(http.StatusCreated, &gin.H{"token": token})
}

// LISTS

func (h *Handler) CreateList(c *gin.Context) {

}

func (h *Handler) GetAllLists(c *gin.Context) {

}

func (h *Handler) GetListById(c *gin.Context) {

}

func (h *Handler) UpdateList(c *gin.Context) {

}
func (h *Handler) DeleteList(c *gin.Context) {

}

//ITEMS

func (h *Handler) CreateItem(c *gin.Context) {

}

func (h *Handler) GetAllItems(c *gin.Context) {

}

func (h *Handler) GetItemById(c *gin.Context) {

}

func (h *Handler) UpdateItem(c *gin.Context) {

}
func (h *Handler) DeleteItem(c *gin.Context) {

}

func NewErrorResponse(c *gin.Context, statusCode int, message string) {
	logrus.Error(message)
	c.AbortWithStatusJSON(statusCode, error{message})
}
