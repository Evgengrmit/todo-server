package service

import (
	"crypto/sha1"
	"fmt"
	"todo/pkg/repository"
	"todo/user"
)

const salt = "zhdieur484hfr584u4eiodhiuahfwieduso"

func NewAuthService(repo repository.Authorization) *AuthService {
	return &AuthService{repo: repo}
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
	}
}

func (s *AuthService) CreateUser(u user.User) (int, error) {
	u.Password = generatePasswordHash(u.Password)
	return s.repo.CreateUser(u)
}

func generatePasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))
	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}
