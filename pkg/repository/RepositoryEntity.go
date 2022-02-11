package repository

import (
	"github.com/jmoiron/sqlx"
	"todo/todo"
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
	Create(userId int, list todo.TodoList) (int, error)
	GetAllLists(userID int) ([]todo.TodoList, error)
	GetListById(userId, id int) (todo.TodoList, error)
	Delete(userId, id int) error
	Update(userId int, id int, input todo.UpdateListInput) error
}

type TodoItem interface {
	Create(listId int, list todo.TodoItem) (int, error)
	GetAllItems(userId, listId int) ([]todo.TodoItem, error)
	GetItemById(userId, id int) (todo.TodoItem, error)
	Delete(userId, id int) error
	Update(userId, id int, input todo.UpdateItemInput) error
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
		TodoList:      NewTodoListPostgres(db),
		TodoItem:      NewTodoItemPostgres(db),
	}
}
