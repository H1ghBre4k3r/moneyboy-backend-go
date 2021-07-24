package server

import (
	"git.pesca.dev/pesca-dev/moneyboy-backend/internal/modules"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

type Server struct {
	app      *fiber.App
	registry *modules.ModuleManager
}

func New() *Server {

	app := fiber.New()

	server := &Server{
		app:      app,
		registry: modules.New(),
	}

	server.setup()
	server.init()

	return server
}

func (s *Server) Start(address string) {
	s.app.Listen(address)
}

func (s *Server) setup() {
	s.app.Use(logger.New())
}

func (s *Server) init() {
	s.registry.Init(s.app)
}
