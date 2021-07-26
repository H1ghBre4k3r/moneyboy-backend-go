package user

import (
	"git.pesca.dev/pesca-dev/moneyboy-backend/internal/database"
)

type UserService struct {
	db *database.Connection
}

func createService(db *database.Connection) *UserService {
	return &UserService{
		db,
	}
}

func (s *UserService) GetProfile(id string) interface{} {
	return s.db.Users().FindById(id)
}
