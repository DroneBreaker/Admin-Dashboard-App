package service

import (
	"github.com/DroneBreaker/Admin-Dashboard-App/backend/internal/models"
	"github.com/DroneBreaker/Admin-Dashboard-App/backend/internal/repository"
)

type UserService interface {
	GetAll() ([]models.User, error)
	Create(user models.User) error
	GetByID(id int) (*models.User, error)
	GetByUsername(username string) (*models.User, error)
}

type userService struct {
	repo repository.UserRepository
}

func NewUserHandler(repo repository.UserRepository) UserService {
	return &userService{repo}
}

func (s *userService) GetAll() ([]models.User, error) {
	return s.repo.GetAll()
}

func (s *userService) Create(user models.User) error {
	return s.repo.Create(&user)
}

func (s *userService) GetByID(id int) (*models.User, error) {
	return s.repo.GetByID(id)
}

func (s *userService) GetByUsername(username string) (*models.User, error) {
	return s.repo.GetByUsername(username)
}
