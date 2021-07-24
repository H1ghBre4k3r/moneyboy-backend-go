package server

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

type server struct {
	app *fiber.App
}

func New() *server {

	app := fiber.New()
	app.Use(logger.New())
	setupRoutes(app)

	return &server{
		app,
	}
}

func (s *server) Start(address string) {
	s.app.Listen(address)
}
