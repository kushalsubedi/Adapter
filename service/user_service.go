package service

import (
	"fmt"
	"project/models"
	"project/repository"
)

// UserService handles business logic for user operations
type UserService struct {
	repo repository.UserRepository
}

// NewUserService creates a new user service
func NewUserService(repo repository.UserRepository) *UserService {
	return &UserService{repo: repo}
}

// RegisterUser creates a new user
func (s *UserService) RegisterUser(name string) error {
	if name == "" {
		return fmt.Errorf("user name cannot be empty")
	}

	user := models.User{Name: name}
	if err := s.repo.Create(user); err != nil {
		return fmt.Errorf("failed to register user: %w", err)
	}

	return nil
}

// ListUsers retrieves all registered users
func (s *UserService) ListUsers() ([]models.User, error) {
	users, err := s.repo.GetAll()
	if err != nil {
		return nil, fmt.Errorf("failed to list users: %w", err)
	}
	return users, nil
}
