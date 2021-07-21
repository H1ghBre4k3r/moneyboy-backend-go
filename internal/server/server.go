package server

import "github.com/gofiber/fiber/v2"

type server struct {
	app *fiber.App
}

func New() *server {

	app := fiber.New()

	setupRoutes(app)

	return &server{
		app,
	}
}

func (s *server) Start(address string) {
	s.app.Listen(address)
}
