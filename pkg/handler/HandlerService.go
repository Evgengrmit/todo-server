package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
	"strconv"
	"todo/pkg/service"
	"todo/todo"
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

// LISTS

func (h *Handler) CreateList(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}
	var input todo.TodoList
	if err = c.BindJSON(&input); err != nil {
		NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	id, err := h.services.TodoList.Create(userId, input)
	if err != nil {
		NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	c.JSON(http.StatusCreated, gin.H{"id": id})
}

func (h *Handler) GetAllLists(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}
	lists, err := h.services.TodoList.GetAllLists(userId)
	if err != nil {
		NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	c.JSON(http.StatusOK, getAllListsResponse{Data: lists})
}

func (h *Handler) GetListById(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}
	listId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		NewErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}
	list, err := h.services.TodoList.GetListById(userId, listId)
	if err != nil {
		NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	c.JSON(http.StatusOK, list)
}

func (h *Handler) UpdateList(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}
	listId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		NewErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}
	var input todo.UpdateListInput
	if err = c.BindJSON(&input); err != nil {
		NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	if err = h.services.Update(userId, listId, input); err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, statusResponse{Status: "ok"})
}

func (h *Handler) DeleteList(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}
	listId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		NewErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}
	err = h.services.TodoList.Delete(userId, listId)
	if err != nil {
		NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	c.JSON(http.StatusOK, statusResponse{Status: "ok"})
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
	c.AbortWithStatusJSON(statusCode, errorResponse{message})
}
