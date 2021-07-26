package user

import (
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
)

type UserController struct {
	service *UserService
}

func createController(service *UserService) *UserController {
	controller := &UserController{service}
	return controller
}

func (ctrl *UserController) getProfile(c *fiber.Ctx) error {
	userClaims := c.Locals("user").(*jwt.Token)
	claims := userClaims.Claims.(jwt.MapClaims)
	id := claims["id"].(string)
	user := ctrl.service.GetProfile(id)
	return c.JSON(user)
}

func (ctrl *UserController) RegisterRoutes(router fiber.Router) {
	router.Get("/profile", ctrl.getProfile)
}
