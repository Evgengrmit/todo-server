package service

import (
	"crypto/sha1"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"time"
	"todo/pkg/repository"
	"todo/user"
)

const (
	salt      = "zhdieur484hfr584u4eiodhiuahfwieduso"
	signedKey = "nx74nfiuyfgd4ju7j8ref74e"
	tokenTTL  = 12 * time.Hour
)

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

func (s *AuthService) GenerateToken(login, password string) (string, error) {
	u, err := s.repo.GetUser(login, generatePasswordHash(password))
	if err != nil {
		return "", err
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256,
		&TokenClaims{
			jwt.StandardClaims{
				ExpiresAt: time.Now().Add(tokenTTL).Unix(),
				IssuedAt:  time.Now().Unix()},
			u.Id,
		})
	return token.SignedString([]byte(signedKey))
}

func generatePasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))
	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}
