package user

import "github.com/gofiber/fiber/v2"

type UserService struct {
}

func createService() *UserService {
	return &UserService{}
}

func (s *UserService) GetProfile(c *fiber.Ctx) error {
	return c.SendString("Hello, world!")
}
