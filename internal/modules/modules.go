package modules

import (
	"git.pesca.dev/pesca-dev/moneyboy-backend/internal/auth"
	"github.com/gofiber/fiber/v2"
)

type ModuleManager struct {
	modules []interface{}
}

func New() *ModuleManager {
	return &ModuleManager{}
}

func (m *ModuleManager) Init(app *fiber.App) {
	auth := auth.New(app)

	m.register(auth)
}

func (m *ModuleManager) register(modules ...interface{}) {
	m.modules = append(m.modules, modules)
}
