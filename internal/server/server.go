package server

import (
	"git.pesca.dev/pesca-dev/moneyboy-backend/internal/modules"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	jwtware "github.com/gofiber/jwt/v2"
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
	s.app.Use(jwtware.New(jwtware.Config{
		SigningKey: []byte("mySigningKey"),
		Filter: func(c *fiber.Ctx) bool {
			return string(c.Request().URI().LastPathSegment()) == "login"
		},
	}))
}

func (s *Server) init() {
	s.registry.InitV1(s.app.Group("/api/v1"))
}
