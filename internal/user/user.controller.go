package user

import (
	"github.com/gofiber/fiber/v2"
)

type UserController struct {
	service *UserService
}

func createController(service *UserService) *UserController {
	controller := &UserController{}
	return controller
}

func (ctrl *UserController) getProfile(c *fiber.Ctx) error {
	return ctrl.service.GetProfile(c)
}

func (ctrl *UserController) RegisterRoutes(router fiber.Router) {
	router.Get("/profile/:id", ctrl.getProfile)
}
