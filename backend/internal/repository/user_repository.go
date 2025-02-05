package repository

import "github.com/DroneBreaker/Admin-Dashboard-App/internal/models"

type UserRepository interface {
	Create(user *models.User)
	GetUserByID(id int) (*models.User, error)
}
