package server

import (
	"git.pesca.dev/pesca-dev/moneyboy-backend/internal/auth"
	"git.pesca.dev/pesca-dev/moneyboy-backend/internal/database"
	"git.pesca.dev/pesca-dev/moneyboy-backend/internal/jwt"
	"git.pesca.dev/pesca-dev/moneyboy-backend/internal/router"
	"git.pesca.dev/pesca-dev/moneyboy-backend/internal/session"
	"git.pesca.dev/pesca-dev/moneyboy-backend/internal/user"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	gojwt "github.com/golang-jwt/jwt"
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
	s.app.Use(jwt.Middleware([]string{"/auth/login", "/auth/register"}))
	// TODO lome: move to module and load session
	s.app.Use(func(c *fiber.Ctx) error {
		userClaims := c.Locals("user")
		if userClaims != nil {
			claims := userClaims.(*gojwt.Token).Claims.(gojwt.MapClaims)
			id := claims["id"].(string)
			c.Locals("sessionId", id)
		}
		return c.Next()
	})

	user := user.New(db.Users())
	session := session.New(db.Sessions(), user)
	auth := auth.New(db, jwt, session)
	router := router.New(s.app.Group("/api/v1"), &router.RouterParams{
		UserService:    user,
		SessionService: session,
		AuthService:    auth,
	})
	s.modules = append(s.modules, db, jwt, auth, user, router)
}
