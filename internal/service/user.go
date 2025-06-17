package service

import (
	"avito-test-pvz/internal/models"
	"github.com/google/uuid"
	"log"
)

type UserRepository interface {
	CreateUser(user *models.User) error
}

func (s *Usersrv) CreateUser(User *models.RegisterRequest) (*models.User, error) {
	log.SetPrefix("Service Create User")
	user := &models.User{
		ID:       uuid.New(),
		Email:    User.Email,
		Password: User.Password,
		Role:     User.Role,
	}
	if err := user.HashPassword(); err != nil {
		log.Printf("Error hashing password: %v", err)
		return nil, err
	}
	if err := s.userRepo.CreateUser(user); err != nil {
		log.Printf("Error creating user: %v", err)
		return nil, err
	}
	return user, nil
}
