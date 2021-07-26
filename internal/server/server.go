package server

import (
	"git.pesca.dev/pesca-dev/moneyboy-backend/internal/auth"
	"git.pesca.dev/pesca-dev/moneyboy-backend/internal/database"
	"git.pesca.dev/pesca-dev/moneyboy-backend/internal/jwt"
	"git.pesca.dev/pesca-dev/moneyboy-backend/internal/user"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Server struct {
	app     *fiber.App
	modules []interface{}
}

func New() *Server {

	app := fiber.New()

	server := &Server{
		app: app,
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
	s.loadModules(s.app.Group("/api/v1"))
}

func (s *Server) loadModules(router fiber.Router) {
	dsn := "root:12345678@tcp(127.0.0.1:3306)/moneyboy?charset=utf8mb4&parseTime=True&loc=Local"
	db := database.New(mysql.Open(dsn), &gorm.Config{})

	jwt := jwt.New("mySigningKey")
	router.Use(jwt.Middleware())

	auth := auth.New(router, db, jwt)
	user := user.New(router, db)
	s.modules = append(s.modules, db, jwt, auth, user)
}
