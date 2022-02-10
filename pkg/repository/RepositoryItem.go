package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
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

func (t *TodoItemPostgres) Update(userId, listId, id int, input todo.UpdateListInput) error {
	//TODO implement me
	panic("implement me")
}
