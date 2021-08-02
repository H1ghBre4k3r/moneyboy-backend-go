package router

import (
	"git.pesca.dev/pesca-dev/moneyboy-backend/internal/global"
	"git.pesca.dev/pesca-dev/moneyboy-backend/internal/models"
	"git.pesca.dev/pesca-dev/moneyboy-backend/internal/validation"
	"github.com/gofiber/fiber/v2"
)

type AuthService interface {
	Login(*global.LoginDTO) (string, string, error)
	Register(*global.RegisterDTO) (bool, error)
	RefreshToken(*global.RefreshTokenDTO) (string, error)
	Logout(*models.Session) error
}

type AuthController struct {
	authService AuthService
}

// Return object after a successful login
type LoginReturn struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

// Return object after a successful token refresh
type RefreshReturn struct {
	AccessToken string `json:"access_token"`
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
	accessToken, refreshToken, err := ctrl.authService.Login(user)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).SendString(err.Error())
	}
	return c.JSON(&LoginReturn{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	})
}

// POST /auth/register
func (ctrl *AuthController) postRegister(c *fiber.Ctx) error {
	user := new(global.RegisterDTO)
	// validate payload and on error, return error
	if validation.New(c).Validate(user) != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}
	status := fiber.StatusAccepted

	// try to register a new user
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

// POST /auth/refresh
func (ctrl *AuthController) postRefresh(c *fiber.Ctx) error {
	// validate the payload
	payload := new(global.RefreshTokenDTO)
	if validation.New(c).Validate(payload) != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	// generate a new access token
	returnValue, err := ctrl.authService.RefreshToken(payload)
	if err != nil {
		if returnValue == "" {
			return c.SendStatus(fiber.StatusInternalServerError)
		} else {
			return c.Status(fiber.StatusUnauthorized).SendString(returnValue)
		}
	}

	return c.JSON(&RefreshReturn{
		AccessToken: returnValue,
	})
}

// DELETE /auth/logout
func (ctrl *AuthController) deleteLogout(c *fiber.Ctx) error {
	session := c.Locals("session").(*models.Session)
	if err := ctrl.authService.Logout(session); err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}
	return c.SendStatus(fiber.StatusAccepted)
}

func (ctrl *AuthController) _registerRoutes(router fiber.Router) {
	router.Post("/login", ctrl.postLogin)
	router.Post("/register", ctrl.postRegister)
	router.Post("/refresh", ctrl.postRefresh)
	router.Delete("/logout", ctrl.deleteLogout)
}
