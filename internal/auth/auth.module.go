package auth

import "github.com/gofiber/fiber/v2"

type AuthModule struct {
	authController *AuthController
}

func New(app *fiber.App) *AuthModule {

	authService := createService()

	authController := createController(authService)
	authController.RegisterRoutes(app.Group("/auth"))

	return &AuthModule{
		authController: authController,
	}
}
