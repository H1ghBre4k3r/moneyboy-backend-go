package auth

import (
	"git.pesca.dev/pesca-dev/moneyboy-backend/internal/database"
	"github.com/gofiber/fiber/v2"
)

type AuthModule struct {
	authController *AuthController
	authService    *AuthService
}

func New(router fiber.Router, db *database.Connection, jwt JWT) *AuthModule {

	authService := createService(db, jwt)

	authController := createController(authService)
	authController.RegisterRoutes(router.Group("/auth"))

	return &AuthModule{
		authController,
		authService,
	}
}

func (m *AuthModule) Service() *AuthService {
	return m.authService
}
