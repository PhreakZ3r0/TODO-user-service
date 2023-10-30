package service

import (
	"github.com/PhreakZ3r0/TODO-user-service/internal/model"
	"github.com/PhreakZ3r0/TODO-user-service/internal/repository"
)

type UserService struct {
	repo *repository.UserRepository
}

func (s *UserService) CreateUser(user *model.User) error {
	return s.repo.CreateUser(user)
}

func (s *UserService) GetUserByID(id string) (*model.User, error) {
	return s.repo.GetUserByID(id)
}
