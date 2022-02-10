package repository

import (
	"github.com/jmoiron/sqlx"
	"todo/user"
)

type Repository struct {
	Authorization
	TodoList
	TodoItem
}

type Authorization interface {
	CreateUser(u user.User) (int, error)
	GetUser(login, password string) (user.User, error)
}

type TodoList interface {
}

type TodoItem interface {
}

type AuthPostgres struct {
	db *sqlx.DB
}
