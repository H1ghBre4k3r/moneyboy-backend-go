package modules

import (
	"git.pesca.dev/pesca-dev/moneyboy-backend/internal/auth"
	"git.pesca.dev/pesca-dev/moneyboy-backend/internal/user"
	"github.com/gofiber/fiber/v2"
)

type ModuleManager struct {
	modules []interface{}
}

func New() *ModuleManager {
	return &ModuleManager{}
}

func (m *ModuleManager) InitV1(router fiber.Router) {
	auth := auth.New(router)
	user := user.New(router)
	m.register(auth, user)
}

func (m *ModuleManager) register(modules ...interface{}) {
	m.modules = append(m.modules, modules)
}
