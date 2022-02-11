package service

import (
	"todo/pkg/repository"
	"todo/todo"
)

type TodoItemService struct {
	repo     repository.TodoItem
	listRepo repository.TodoList
}

func NewTodoItemService(repo repository.TodoItem, listRepo repository.TodoList) *TodoItemService {
	return &TodoItemService{repo: repo, listRepo: listRepo}
}

func (s *TodoItemService) Create(userId, listId int, list todo.TodoItem) (int, error) {
	_, err := s.listRepo.GetListById(userId, listId)
	if err != nil {
		return 0, err
	}
	return s.repo.Create(listId, list)
}

func (s *TodoItemService) GetAllItems(userId, listId int) ([]todo.TodoItem, error) {
	return s.repo.GetAllItems(userId, listId)
}

func (s *TodoItemService) GetItemById(userId, id int) (todo.TodoItem, error) {
	return s.repo.GetItemById(userId, id)
}

func (s *TodoItemService) Delete(userId, id int) error {
	return s.repo.Delete(userId, id)
}

func (s *TodoItemService) Update(userId, id int, input todo.UpdateItemInput) error {
	return s.repo.Update(userId, id, input)
}
