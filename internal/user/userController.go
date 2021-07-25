package user

import (
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
)

type UserController struct {
	service *UserService
}

func createController(service *UserService) *UserController {
	controller := &UserController{}
	return controller
}

func (ctrl *UserController) getProfile(c *fiber.Ctx) error {
	user := c.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	id := claims["id"].(string)
	return c.SendString(id)
}

func (ctrl *UserController) RegisterRoutes(router fiber.Router) {
	router.Get("/profile", ctrl.getProfile)
}
