package auth

import (
	"git.pesca.dev/pesca-dev/moneyboy-backend/internal/validation"
	"github.com/gofiber/fiber/v2"
)

type AuthController struct {
	service *AuthService
}

func createController(authService *AuthService) *AuthController {
	controller := &AuthController{
		authService,
	}
	return controller
}

func (ctrl *AuthController) postLogin(c *fiber.Ctx) error {
	user := new(LoginDTO)
	if validation.New(c).Validate(user) != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}
	return ctrl.service.Login(c, user)
}

func (ctrl *AuthController) postRegister(c *fiber.Ctx) error {
	user := new(RegisterDTO)
	if validation.New(c).Validate(user) != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}
	return ctrl.service.Register(c, user)
}

func (ctrl *AuthController) RegisterRoutes(router fiber.Router) {
	router.Post("/login", ctrl.postLogin)
	router.Post("/register", ctrl.postRegister)
}
