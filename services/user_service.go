// services/user_service.go
package services

import (
	"github.com/okyirmawan/auth_service/models"
	"github.com/okyirmawan/auth_service/repositories"
)

type UserService interface {
	GetProfile(email string) (*models.User, error)
}

type userService struct {
	userRepository repositories.UserRepository
}

func NewUserService(userRepo repositories.UserRepository) UserService {
	return &userService{userRepo}
}

func (s *userService) GetProfile(email string) (*models.User, error) {
	return s.userRepository.FindByEmail(email)
}
