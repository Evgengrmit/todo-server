package todo

type TodoList struct {
	Id          uint   `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
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
