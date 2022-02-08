package service

import "todo/pkg/repository"

func NewServer(repos *repository.Repository) *Service {
	return &Service{}
}
