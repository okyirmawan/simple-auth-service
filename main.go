package main

import (
	"flag"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/okyirmawan/auth_service/config"
	"github.com/okyirmawan/auth_service/controllers"
	"github.com/okyirmawan/auth_service/database/seeds"
	"github.com/okyirmawan/auth_service/middlewares"
	"github.com/okyirmawan/auth_service/models"
	"github.com/okyirmawan/auth_service/repositories"
	"github.com/okyirmawan/auth_service/services"
	"log"
)

func main() {
	// Parse command-line arguments.
	seedCmd := flag.Bool("seed", false, "Run database seeder")
	flag.Parse()

	// Load .env file
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading .env file")
	}

	// Initialize the database
	config.InitDB()

	if *seedCmd {
		// Run the seeder
		seeds.SeedUsers()
		fmt.Println("Seeding completed successfully!")
		return
	}

	models.InitJwtKey()

	// Initialization repository, service, dan controller
	userRepo := repositories.NewUserRepository(config.DB)

	authService := services.NewAuthService(userRepo)
	authController := controllers.NewAuthController(authService)

	userService := services.NewUserService(userRepo)
	userController := controllers.NewUserController(userService)

	r := gin.Default()

	r.POST("/login", authController.Login)
	r.POST("/refresh-token", authController.RefreshToken)
	r.GET("/profile", middlewares.Auth(), userController.GetProfile)

	r.Run(":8080")
}
