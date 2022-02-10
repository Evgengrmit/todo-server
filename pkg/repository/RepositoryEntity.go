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
}

type TodoList interface {
}

type TodoItem interface {
}

type AuthPostgres struct {
	db *sqlx.DB
}
