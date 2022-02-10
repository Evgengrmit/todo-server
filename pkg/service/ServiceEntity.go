package service

import (
	"todo/pkg/repository"
	"todo/todo"
	"todo/user"
)

type Service struct {
	Authorization
	TodoList
	TodoItem
}

type Authorization interface {
	CreateUser(u user.User) (int, error)
	GenerateToken(login, password string) (string, error)
	ParseToken(token string) (int, error)
}

type TodoList interface {
	Create(userId int, list todo.TodoList) (int, error)
	GetAllLists(userId int) ([]todo.TodoList, error)
	GetListById(userId, id int) (todo.TodoList, error)
	Delete(userId, id int) error
	Update(userId int, id int, input todo.UpdateListInput) error
}

type TodoItem interface {
}

type TodoListService struct {
	repo repository.TodoList
}
