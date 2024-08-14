package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/okyirmawan/auth_service/services"
	"net/http"
)

type UserController struct {
	userService services.UserService
}

func NewUserController(userService services.UserService) *UserController {
	return &UserController{userService}
}

func (ctrl *UserController) GetProfile(c *gin.Context) {
	email, exists := c.Get("email")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid tokens"})
		return
	}

	user, err := ctrl.userService.GetProfile(email.(string))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"id":    user.ID,
		"Email": user.Email,
		"Name":  user.Name,
	})
}
