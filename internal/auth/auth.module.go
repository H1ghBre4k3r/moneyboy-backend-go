package auth

import (
	"git.pesca.dev/pesca-dev/moneyboy-backend/internal/database"
	"github.com/gofiber/fiber/v2"
)

type AuthModule struct {
	authController *AuthController
	authService    *AuthService
}

func New(router fiber.Router, db *database.Connection) *AuthModule {

	authService := createService(db)

	authController := createController(authService)
	authController.RegisterRoutes(router.Group("/auth"))

	return &AuthModule{
		authController,
		authService,
	}
}
