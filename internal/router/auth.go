package router

import (
	"git.pesca.dev/pesca-dev/moneyboy-backend/internal/global"
	"git.pesca.dev/pesca-dev/moneyboy-backend/internal/validation"
	"github.com/gofiber/fiber/v2"
)

type AuthService interface {
	Login(*global.LoginDTO) (interface{}, error)
	Register(*global.RegisterDTO) (bool, error)
}

type AuthController struct {
	authService AuthService
}

func authController(router fiber.Router, authService AuthService) *AuthController {
	authController := &AuthController{
		authService,
	}
	authController.registerRoutes(router)
	return authController
}

func (ctrl *AuthController) postLogin(c *fiber.Ctx) error {
	user := new(global.LoginDTO)
	if validation.New(c).Validate(user) != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}
	retVal, err := ctrl.authService.Login(user)
	if err != nil {
		return c.SendStatus(fiber.StatusUnauthorized)
	}
	return c.JSON(retVal)
}

func (ctrl *AuthController) postRegister(c *fiber.Ctx) error {
	user := new(global.RegisterDTO)
	if validation.New(c).Validate(user) != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}
	status := fiber.StatusAccepted
	internal, err := ctrl.authService.Register(user)
	if err != nil {
		if internal {
			status = fiber.StatusInternalServerError
		} else {
			status = fiber.StatusBadRequest
		}
	}
	return c.SendStatus(status)
}

func (ctrl *AuthController) registerRoutes(router fiber.Router) {
	router.Post("/login", ctrl.postLogin)
	router.Post("/register", ctrl.postRegister)
}
