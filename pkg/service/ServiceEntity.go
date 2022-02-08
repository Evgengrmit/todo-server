package service

type Service struct {
	Authorization
	TodoList
	TodoItem
}

type Authorization interface {
}

type TodoList interface {
}

type TodoItem interface {
}
