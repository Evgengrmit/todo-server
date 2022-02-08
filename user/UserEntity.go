package user

type User struct {
	Id       uint   `json:"-"`
	Name     string `json:"name"`
	Login    string `json:"login"`
	Password []byte `json:"password"`
}
