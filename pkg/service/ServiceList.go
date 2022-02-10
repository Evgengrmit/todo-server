package service

import (
	"todo/pkg/repository"
	"todo/todo"
)

type TodoListService struct {
	repo repository.TodoList
}

func NewTodoListService(repo repository.TodoList) *TodoListService {
	return &TodoListService{repo: repo}
}

func (l *TodoListService) Create(userId int, list todo.TodoList) (int, error) {
	return l.repo.Create(userId, list)
}
func (l *TodoListService) GetAllLists(userId int) ([]todo.TodoList, error) {
	return l.repo.GetAllLists(userId)

}
func (l *TodoListService) GetListById(userId, id int) (todo.TodoList, error) {
	return l.repo.GetListById(userId, id)
}

func (l *TodoListService) Delete(userId, id int) error {
	return l.repo.Delete(userId, id)
}

func (l *TodoListService) Update(userId int, listId int, input todo.UpdateListInput) error {
	if err := input.Validate(); err != nil {
		return err
	}
	return l.repo.Update(userId, listId, input)
}
