package repository

import (
	"github.com/PhreakZ3r0/TODO-user-service/internal/model"
)

type UserRepository struct {
	// DB connection

}

func (repo *UserRepository) CreateUser(user *model.User) error {
	// Implement DB logic
	
	return nil
}

func (repo *UserRepository) GetUserByID(id string) (*model.User, error) {
	// Implement DB logic
	return nil, nil
}

