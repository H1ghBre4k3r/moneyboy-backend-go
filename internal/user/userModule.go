package user

import (
	"git.pesca.dev/pesca-dev/moneyboy-backend/internal/database"
	"github.com/gofiber/fiber/v2"
)

type UserModule struct {
	controller *UserController
	service    *UserService
}

func New(router fiber.Router, db *database.Connection) *UserModule {

	userService := createService(db)

	userController := createController(userService)
	userController.RegisterRoutes(router.Group("/user"))

	return &UserModule{
		service:    userService,
		controller: userController,
	}
}
