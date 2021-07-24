package auth

import "github.com/gofiber/fiber/v2"

type AuthModule struct {
	authController *AuthController
	authService    *AuthService
}

func New(router fiber.Router) *AuthModule {

	authService := createService()

	authController := createController(authService)
	authController.RegisterRoutes(router.Group("/auth"))

	return &AuthModule{
		authController,
		authService,
	}
}
