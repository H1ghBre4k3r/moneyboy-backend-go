package user

import "github.com/gofiber/fiber/v2"

type UserModule struct {
	controller *UserController
	service    *UserService
}

func New(router fiber.Router) *UserModule {

	userService := createService()

	userController := createController(userService)
	userController.RegisterRoutes(router.Group("/user"))

	return &UserModule{
		service:    userService,
		controller: userController,
	}
}
