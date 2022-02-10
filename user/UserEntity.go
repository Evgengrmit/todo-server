package user

type User struct {
	Id       uint   `json:"-"`
	Name     string `json:"name" binding:"required"`
	Login    string `json:"login" binding:"required"`
	Password string `json:"password" binding:"required"`
}
