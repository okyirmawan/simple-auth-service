package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/okyirmawan/auth_service/models"
	"github.com/okyirmawan/auth_service/services"
	"net/http"
)

type AuthController struct {
	authService services.AuthService
}

func NewAuthController(authService services.AuthService) *AuthController {
	return &AuthController{authService}
}

func (ctrl *AuthController) Login(c *gin.Context) {
	var credentials = models.Credentials{}

	if err := c.BindJSON(&credentials); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	user, accessToken, err := ctrl.authService.Login(credentials.Email, credentials.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"access_token": accessToken,
		"user":         user,
	})
}

func (ctrl *AuthController) RefreshToken(c *gin.Context) {
	var requestBody struct {
		RefreshToken string `json:"refresh_token"`
	}

	if err := c.BindJSON(&requestBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	accessToken, newRefreshToken, err := ctrl.authService.RefreshToken(requestBody.RefreshToken)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"access_token":  accessToken,
		"refresh_token": newRefreshToken,
	})
}
