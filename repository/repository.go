package repository

import "project/models"

// UserRepository defines the contract for user data access
type UserRepository interface {
	Create(user models.User) error
	GetAll() ([]models.User, error)
}
