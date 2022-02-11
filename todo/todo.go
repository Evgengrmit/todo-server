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
	Id          uint   `json:"id" db:"id"`
	Title       string `json:"title" db:"title" binding:"required"`
	Description string `json:"description" db:"description"`
	Status      bool   `json:"status"  db:"done"`
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

type UpdateItemInput struct {
	Title       *string `json:"title,omitempty"`
	Description *string `json:"description,omitempty"`
	Status      *bool   `json:"status,omitempty"`
}

func (i UpdateItemInput) Validate() error {
	if i.Title == nil && i.Description == nil && i.Status == nil {
		return errors.New("update structure has no values")
	}
	return nil
}
