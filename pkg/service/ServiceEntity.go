package service

import (
	"todo/pkg/repository"
	"todo/user"
)

type Service struct {
	Authorization
	TodoList
	TodoItem
}

type Authorization interface {
	CreateUser(u user.User) (int, error)
}

type TodoList interface {
}

type TodoItem interface {
}

type AuthService struct {
	repo repository.Authorization
}
