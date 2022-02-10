package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
	"strings"
	"todo/todo"
)

func NewTodoListPostgres(db *sqlx.DB) *TodoListPostgres {
	return &TodoListPostgres{db: db}
}

func (p *TodoListPostgres) Create(userId int, list todo.TodoList) (int, error) {
	tx, err := p.db.Begin()
	if err != nil {
		return 0, err
	}
	var id int
	createListQuery := fmt.Sprintf("INSERT INTO %s (title,description) VALUES ($1, $2) RETURNING id", todoListsTable)
	row := tx.QueryRow(createListQuery, list.Title, list.Description)
	if err = row.Scan(&id); err != nil {
		_ = tx.Rollback()
		return 0, err
	}
	createUsersListQuery := fmt.Sprintf("INSERT INTO %s (user_id,list_id) VALUES ($1, $2)", userListsTable)
	_, err = tx.Exec(createUsersListQuery, userId, id)
	if err != nil {
		_ = tx.Rollback()
		return 0, err
	}
	return id, tx.Commit()
}

func (p *TodoListPostgres) GetAllLists(userId int) ([]todo.TodoList, error) {
	var lists []todo.TodoList
	query := fmt.Sprintf(`SELECT tl.id, tl.title, tl.description FROM %s tl JOIN %s ul 
								ON tl.id=ul.list_id WHERE ul.user_id=$1`, todoListsTable, userListsTable)
	err := p.db.Select(&lists, query, userId)
	return lists, err
}

func (p *TodoListPostgres) GetListById(userId, id int) (todo.TodoList, error) {
	var list todo.TodoList
	query := fmt.Sprintf(`SELECT tl.id, tl.title, tl.description FROM %s tl JOIN %s ul
								ON tl.id=ul.list_id WHERE ul.user_id=$1 AND ul.list_id=$2`, todoListsTable, userListsTable)
	err := p.db.Get(&list, query, userId, id)
	return list, err
}

func (p *TodoListPostgres) Delete(userId, id int) error {
	query := fmt.Sprintf(`DELETE FROM  %s tl USING %s ul WHERE tl.id=ul.list_id AND ul.user_id=$1 AND ul.list_id=$2`,
		todoListsTable, userListsTable)
	_, err := p.db.Exec(query, userId, id)
	return err
}

func (p *TodoListPostgres) Update(userId int, id int, input todo.UpdateListInput) error {
	setValues := make([]string, 0)
	args := make([]interface{}, 0)
	argId := 1
	if input.Title != nil {
		setValues = append(setValues, fmt.Sprintf("title=$%d", argId))
		args = append(args, *input.Title)
		argId++
	}
	if input.Description != nil {
		setValues = append(setValues, fmt.Sprintf("description=$%d", argId))
		args = append(args, *input.Description)
		argId++
	}
	setQuery := strings.Join(setValues, ", ")
	query := fmt.Sprintf("UPDATE %s tl SET %s FROM %s ul WHERE tl.id=ul.list_id AND ul.user_id=$%d AND ul.list_id=$%d",
		todoListsTable, setQuery, userListsTable, argId, argId+1)
	args = append(args, userId, id)
	logrus.Debugf("updateQuery: %s", query)
	logrus.Debugf("args: %s", args)
	_, err := p.db.Exec(query, args...)
	return err
}
