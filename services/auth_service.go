package services

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"github.com/okyirmawan/auth_service/models"
	"github.com/okyirmawan/auth_service/repositories"
	"golang.org/x/crypto/bcrypt"
	"time"
)

type AuthService interface {
	Login(email, password string) (*models.User, string, error)
	RefreshToken(refreshToken string) (string, string, error)
}

type authService struct {
	userRepository repositories.UserRepository
}

func NewAuthService(userRepo repositories.UserRepository) AuthService {
	return &authService{userRepo}
}

func (s *authService) Login(email, password string) (*models.User, string, error) {
	user, err := s.userRepository.FindByEmail(email)
	if err != nil {
		return nil, "", errors.New("invalid email or password")
	}

	// Verify password
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return nil, "", errors.New("invalid email or password")
	}

	// Generate access token
	expirationTime := time.Now().Add(1 * time.Minute)
	claims := &models.Claims{
		Email: email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	accessTokenString, err := accessToken.SignedString(models.JwtKey)
	if err != nil {
		return nil, "", errors.New("could not generate access token")
	}

	// Generate refresh token
	refreshExpirationTime := time.Now().Add(30 * 24 * time.Hour)
	refreshClaims := &models.Claims{
		Email: email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: refreshExpirationTime.Unix(),
		},
	}

	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshClaims)
	refreshTokenString, err := refreshToken.SignedString(models.JwtKey)
	if err != nil {
		return nil, "", errors.New("could not generate refresh token")
	}

	// Save refresh token ke database
	user.RefreshToken = refreshTokenString
	if err := s.userRepository.Save(user); err != nil {
		return nil, "", errors.New("could not save refresh token")
	}

	return user, accessTokenString, nil
}

func (s *authService) RefreshToken(refreshToken string) (string, string, error) {
	claims := &models.Claims{}
	token, err := jwt.ParseWithClaims(refreshToken, claims, func(token *jwt.Token) (interface{}, error) {
		return models.JwtKey, nil
	})

	if err != nil || !token.Valid {
		return "", "", errors.New("invalid refresh token")
	}

	user, err := s.userRepository.FindByEmail(claims.Email)
	if err != nil || user.RefreshToken != refreshToken {
		return "", "", errors.New("invalid refresh token")
	}

	expirationTime := time.Now().Add(15 * time.Minute)
	newClaims := &models.Claims{
		Email: claims.Email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	newToken := jwt.NewWithClaims(jwt.SigningMethodHS256, newClaims)
	accessTokenString, err := newToken.SignedString(models.JwtKey)
	if err != nil {
		return "", "", errors.New("could not generate access token")
	}

	newRefreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, &models.Claims{
		Email: claims.Email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(30 * 24 * time.Hour).Unix(),
		},
	})
	newRefreshTokenString, err := newRefreshToken.SignedString(models.JwtKey)
	if err != nil {
		return "", "", errors.New("could not generate refresh token")
	}

	if err := s.userRepository.UpdateRefreshToken(claims.Email, newRefreshTokenString); err != nil {
		return "", "", errors.New("could not update refresh token")
	}

	return accessTokenString, newRefreshTokenString, nil
}
