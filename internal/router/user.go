package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
)

type UserService interface {
	GetProfile(string) interface{}
}

type UserController struct {
	userService UserService
}

func userController(router fiber.Router, userService UserService) *UserController {
	controller := &UserController{
		userService,
	}
	controller.register(router)
	return controller
}

func (ctrl *UserController) getProfile(c *fiber.Ctx) error {
	userClaims := c.Locals("user").(*jwt.Token)
	claims := userClaims.Claims.(jwt.MapClaims)
	id := claims["id"].(string)
	user := ctrl.userService.GetProfile(id)
	return c.JSON(user)
}

func (ctrl *UserController) register(router fiber.Router) {
	router.Get("/profile", ctrl.getProfile)
}
