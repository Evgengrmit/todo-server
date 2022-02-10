package todo

import "errors"

type TodoList struct {
	Id          uint   `json:"id" db:"id"`
	Title       string `json:"title" db:"title" binding:"required"`
	Description string `json:"description" db:"description"`
}

type UserList struct {
	Id     uint
	UserId uint
	ListId uint
}
type TodoItem struct {
	Id          uint   `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Status      bool   `json:"status"`
}

type ListItem struct {
	Id     uint
	UserId uint
	ListId uint
}

type UpdateListInput struct {
	Title       *string `json:"title,omitempty"`
	Description *string `json:"description,omitempty"`
}

func (i UpdateListInput) Validate() error {
	if i.Title == nil && i.Description == nil {
		return errors.New("update structure has no values")
	}
	return nil
}
