package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"todo/user"
)

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
	}
}

func NewAuthPostgres(db *sqlx.DB) *AuthPostgres {
	return &AuthPostgres{db: db}
}

func (p *AuthPostgres) CreateUser(u user.User) (int, error) {
	var id int
	query := fmt.Sprintf("INSERT INTO %s (name,username,password_hash) values($1, $2, $3)  RETURNING id", usersTable)
	row := p.db.QueryRow(query, u.Name, u.Login, u.Password)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}
	return id, nil
}
