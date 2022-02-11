package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
	"strings"
	"todo/todo"
)

type TodoItemPostgres struct {
	db *sqlx.DB
}

func NewTodoItemPostgres(db *sqlx.DB) *TodoItemPostgres {
	return &TodoItemPostgres{db: db}
}

func (t *TodoItemPostgres) Create(listId int, item todo.TodoItem) (int, error) {
	tx, err := t.db.Begin()
	if err != nil {
		return 0, err
	}
	var id int
	createItemQuery := fmt.Sprintf("INSERT INTO %s (title,description) VALUES ($1, $2) RETURNING id", todoItemsTable)
	row := tx.QueryRow(createItemQuery, item.Title, item.Description)
	if err = row.Scan(&id); err != nil {
		_ = tx.Rollback()
		return 0, err
	}
	createListsItemsQuery := fmt.Sprintf("INSERT INTO %s (item_id,list_id) VALUES ($1, $2)", listsItemsTable)
	_, err = tx.Exec(createListsItemsQuery, id, listId)
	if err != nil {
		_ = tx.Rollback()
		return 0, err
	}
	return id, tx.Commit()
}

func (t *TodoItemPostgres) GetAllItems(userId, listId int) ([]todo.TodoItem, error) {
	var items []todo.TodoItem
	query := fmt.Sprintf(`SELECT ti.id,ti.title, ti.description,ti.done FROM %s ti JOIN %s li 
								ON ti.id=li.item_id JOIN %s ul ON ul.list_id = li.list_id WHERE ul.user_id=$1 AND li.list_id=$2`,
		todoItemsTable, listsItemsTable, userListsTable)

	err := t.db.Select(&items, query, userId, listId)
	return items, err
}

func (t *TodoItemPostgres) GetItemById(userId, id int) (todo.TodoItem, error) {
	var item todo.TodoItem
	query := fmt.Sprintf(`SELECT ti.id,ti.title, ti.description,ti.done FROM %s ti JOIN %s li 
								ON ti.id=li.item_id JOIN %s ul ON ul.list_id = li.list_id WHERE ul.user_id=$1 AND ti.id=$2`,
		todoItemsTable, listsItemsTable, userListsTable)

	err := t.db.Get(&item, query, userId, id)
	return item, err
}

func (t *TodoItemPostgres) Delete(userId, id int) error {
	query := fmt.Sprintf(`DELETE FROM  %s ti USING %s li,%s ul WHERE ti.id=li.item_id AND ul.list_id = li.list_id 
								AND ul.user_id=$1 AND ti.id=$2`,
		todoItemsTable, listsItemsTable, userListsTable)
	_, err := t.db.Exec(query, userId, id)
	return err
}

func (t *TodoItemPostgres) Update(userId, id int, input todo.UpdateItemInput) error {
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
	if input.Status != nil {
		setValues = append(setValues, fmt.Sprintf("done=$%d", argId))
		args = append(args, *input.Status)
		argId++
	}
	setQuery := strings.Join(setValues, ", ")
	query := fmt.Sprintf(`UPDATE %s ti SET %s FROM %s li, %s ul WHERE ti.id=li.item_id AND ul.list_id = li.list_id
		 AND ul.user_id=$%d AND li.item_id=$%d`,
		todoItemsTable, setQuery, listsItemsTable, userListsTable, argId, argId+1)
	args = append(args, userId, id)
	logrus.Debugf("updateQuery: %s", query)
	logrus.Debugf("args: %s", args)
	_, err := t.db.Exec(query, args...)
	return err
}
