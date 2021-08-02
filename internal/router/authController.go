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
	authController._registerRoutes(router)
	return authController
}

// POST /auth/login
func (ctrl *AuthController) postLogin(c *fiber.Ctx) error {
	user := new(global.LoginDTO)
	// validate payload and on error, return error
	if validation.New(c).Validate(user) != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}
	// process user data
	retVal, err := ctrl.authService.Login(user)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).SendString(err.Error())
	}
	return c.JSON(retVal)
}

// POST /auth/register
func (ctrl *AuthController) postRegister(c *fiber.Ctx) error {
	user := new(global.RegisterDTO)
	// validate payload and on error, return error
	if validation.New(c).Validate(user) != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}
	status := fiber.StatusAccepted
	isInternalError, err := ctrl.authService.Register(user)
	if err != nil {
		if isInternalError {
			status = fiber.StatusInternalServerError
		} else {
			status = fiber.StatusBadRequest
		}
	}
	return c.SendStatus(status)
}

func (ctrl *AuthController) postRefresh(c *fiber.Ctx) error {
	// TODO lome: implement
	return c.SendStatus(fiber.StatusInternalServerError)
}

func (ctrl *AuthController) _registerRoutes(router fiber.Router) {
	router.Post("/login", ctrl.postLogin)
	router.Post("/register", ctrl.postRegister)
	router.Post("/refresh", ctrl.postRefresh)
}
