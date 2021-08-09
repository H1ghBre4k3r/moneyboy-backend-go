package server

import (
	"fmt"

	"git.pesca.dev/pesca-dev/moneyboy-backend/internal/auth"
	"git.pesca.dev/pesca-dev/moneyboy-backend/internal/database"
	"git.pesca.dev/pesca-dev/moneyboy-backend/internal/jwt"
	"git.pesca.dev/pesca-dev/moneyboy-backend/internal/router"
	"git.pesca.dev/pesca-dev/moneyboy-backend/internal/session"
	"git.pesca.dev/pesca-dev/moneyboy-backend/internal/user"
	"git.pesca.dev/pesca-dev/moneyboy-backend/internal/variables"
	"github.com/gofiber/fiber/v2"
	fiberlog "github.com/gofiber/fiber/v2/middleware/logger"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	gormlog "gorm.io/gorm/logger"
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
	s.app.Use(fiberlog.New())
}

func (s *Server) init() {
	s.loadModules()
}

func getDbDns() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", variables.DATABASE.Username, variables.DATABASE.Password, variables.DATABASE.Host, variables.DATABASE.Port, variables.DATABASE.Name)
}

func (s *Server) loadModules() {
	dsn := getDbDns()
	db := database.New(mysql.Open(dsn), &gorm.Config{
		Logger: gormlog.Default.LogMode(gormlog.Error),
	})

	tokenJwt := jwt.New(variables.TOKEN.AccessTokenSecret)
	refreshJwt := jwt.New(variables.TOKEN.RefreshTokenSecret)
	s.app.Use(tokenJwt.Middleware([]string{"/auth/login", "/auth/register", "/auth/refresh"}))

	user := user.New(db.Users())

	session := session.New(db.Sessions(), user)
	s.app.Use(session.Middleware())

	auth := auth.New(db, tokenJwt, refreshJwt, session)

	router := router.New(s.app.Group("/api/v1"), &router.RouterParams{
		UserService:    user,
		SessionService: session,
		AuthService:    auth,
	})
	s.modules = append(s.modules, db, tokenJwt, refreshJwt, user, auth, router)
}
