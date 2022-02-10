package handler

import "todo/todo"

type signInInput struct {
	Login    string `json:"login" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type getAllListsResponse struct {
	Data []todo.TodoList `json:"data"`
}
type statusResponse struct {
	Status string `json:"status"`
}
