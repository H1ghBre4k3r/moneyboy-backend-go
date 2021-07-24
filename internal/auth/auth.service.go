package auth

import "github.com/gofiber/fiber/v2"

type AuthService struct {
}

func createService() *AuthService {
	return &AuthService{}
}

func (s *AuthService) Login(c *fiber.Ctx, user *LoginDTO) error {
	return c.JSON(user)
}

func (s *AuthService) Register(c *fiber.Ctx, user *RegisterDTO) error {
	return c.JSON(user)
}
