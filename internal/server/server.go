package server

import (
	"git.pesca.dev/pesca-dev/moneyboy-backend/internal/auth"
	"git.pesca.dev/pesca-dev/moneyboy-backend/internal/database"
	"git.pesca.dev/pesca-dev/moneyboy-backend/internal/jwt"
	"git.pesca.dev/pesca-dev/moneyboy-backend/internal/router"
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
	s.loadModules()
}

func (s *Server) loadModules() {
	dsn := "root:12345678@tcp(127.0.0.1:3306)/moneyboy?charset=utf8mb4&parseTime=True&loc=Local"
	db := database.New(mysql.Open(dsn), &gorm.Config{})

	jwt := jwt.New("mySigningKey")
	s.app.Use(jwt.Middleware())

	auth := auth.New(db, jwt)
	user := user.New(db)
	router := router.New(s.app.Group("/api/v1"), &router.RouterParams{
		UserService: user,
		AuthService: auth,
	})
	s.modules = append(s.modules, db, jwt, auth, user, router)
}
