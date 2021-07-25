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
	retVal, err := ctrl.service.Login(user)
	if err != nil {
		return c.SendStatus(fiber.StatusUnauthorized)
	}
	return c.JSON(retVal)
}

func (ctrl *AuthController) postRegister(c *fiber.Ctx) error {
	user := new(RegisterDTO)
	if validation.New(c).Validate(user) != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}
	status := fiber.StatusAccepted
	internal, err := ctrl.service.Register(user)
	if err != nil {
		if internal {
			status = fiber.StatusInternalServerError
		} else {
			status = fiber.StatusBadRequest
		}
	}
	return c.SendStatus(status)
}

func (ctrl *AuthController) RegisterRoutes(router fiber.Router) {
	router.Post("/login", ctrl.postLogin)
	router.Post("/register", ctrl.postRegister)
}
