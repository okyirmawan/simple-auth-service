package seeds

import (
	"github.com/okyirmawan/auth_service/config"
	"github.com/okyirmawan/auth_service/models"
	"golang.org/x/crypto/bcrypt"
	"log"
)

// HashPassword hashes a password using bcrypt.
func hashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}

// SeedUsers seeds the database with example users.
func SeedUsers() {
	db := config.DB

	users := []models.User{
		{Email: "example1@example.com", Name: "John Doe", Password: "password123", RefreshToken: "token1"},
		{Email: "example2@example.com", Name: "Jane Smith", Password: "password456", RefreshToken: "token2"},
	}

	for i, user := range users {
		hashedPassword, err := hashPassword(user.Password)
		if err != nil {
			log.Fatalf("Failed to hash password for user %s: %v", user.Email, err)
		}
		users[i].Password = hashedPassword
	}

	if err := db.Create(&users).Error; err != nil {
		log.Fatalf("Failed to seed users: %v", err)
	}
}
