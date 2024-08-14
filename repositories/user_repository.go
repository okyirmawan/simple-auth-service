package repositories

import (
	"github.com/okyirmawan/auth_service/models"
	"gorm.io/gorm"
)

type UserRepository interface {
	FindByEmail(email string) (*models.User, error)
	Save(user *models.User) error
	UpdateRefreshToken(email, refreshToken string) error
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db}
}

func (r *userRepository) FindByEmail(email string) (*models.User, error) {
	var user models.User
	if err := r.db.Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *userRepository) Save(user *models.User) error {
	return r.db.Save(user).Error
}

func (r *userRepository) UpdateRefreshToken(email, refreshToken string) error {
	return r.db.Model(&models.User{}).Where("email = ?", email).Update("refresh_token", refreshToken).Error
}
