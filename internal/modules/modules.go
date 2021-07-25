package modules

import (
	"git.pesca.dev/pesca-dev/moneyboy-backend/internal/auth"
	"git.pesca.dev/pesca-dev/moneyboy-backend/internal/database"
	"git.pesca.dev/pesca-dev/moneyboy-backend/internal/user"
	"github.com/gofiber/fiber/v2"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type ModuleManager struct {
	modules []interface{}
}

func New() *ModuleManager {
	return &ModuleManager{}
}

func (m *ModuleManager) InitV1(router fiber.Router) {
	dsn := "root:12345678@tcp(127.0.0.1:3306)/moneyboy?charset=utf8mb4&parseTime=True&loc=Local"
	db := database.New(mysql.Open(dsn), &gorm.Config{})
	auth := auth.New(router, db)
	user := user.New(router)
	m.register(db, auth, user)
}

func (m *ModuleManager) register(modules ...interface{}) {
	m.modules = append(m.modules, modules)
}
