package user

import "git.pesca.dev/pesca-dev/moneyboy-backend/internal/models"

type Database interface {
	FindByUsername(string) *models.User
	FindById(string) *models.User
	Create(*models.User) error
	DeleteById(string) error
	DeleteByUsername(string) error
}

type UserService struct {
	db Database
}

func New(db Database) *UserService {
	return &UserService{
		db,
	}
}

func (s *UserService) GetUser(id string) *models.User {
	return s.db.FindById(id)
}
