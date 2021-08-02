package router

import (
	"git.pesca.dev/pesca-dev/moneyboy-backend/internal/models"
	"github.com/gofiber/fiber/v2"
)

type UserService interface {
	GetUser(string) *models.User
}

type SessionService interface {
	GetSession(string) *models.Session
}

type UserController struct {
	userService    UserService
	sessionService SessionService
}

func userController(router fiber.Router, userService UserService, sessionService SessionService) *UserController {
	controller := &UserController{
		userService,
		sessionService,
	}
	controller._registerRoutes(router)
	return controller
}

// GET /user/profile
func (ctrl *UserController) getProfile(c *fiber.Ctx) error {
	id := c.Locals("sessionId").(string)
	session := ctrl.sessionService.GetSession(id)
	user := ctrl.userService.GetUser(session.UserID)
	return c.JSON(user)
}

func (ctrl *UserController) _registerRoutes(router fiber.Router) {
	router.Get("/profile", ctrl.getProfile)
}
