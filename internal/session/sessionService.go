package session

import (
	"errors"
	"fmt"

	"git.pesca.dev/pesca-dev/moneyboy-backend/internal/models"
	"github.com/google/uuid"
)

type Database interface {
	Create(*models.Session) error
	Get(string) *models.Session
	Delete(string) error
}

type UserService interface {
	GetUser(string) *models.User
}

type SessionService struct {
	db          Database
	userService UserService
}

func New(db Database, userService UserService) *SessionService {
	return &SessionService{db, userService}
}

// Create a new session for a user
func (s *SessionService) CreateSession(userId string) (string, error) {
	user := s.userService.GetUser(userId)
	if user == nil {
		return "", errors.New("user not found")
	}
	session := &models.Session{
		ID:   uuid.NewString(),
		User: *user,
	}
	fmt.Println(session.ID)
	err := s.db.Create(session)
	return session.ID, err
}

// Get a session by its id
func (s *SessionService) GetSession(sessionId string) *models.Session {
	return s.db.Get(sessionId)
}

// Destroy a session
func (s *SessionService) DestroySession(sessionId string) error {
	return s.db.Delete(sessionId)
}
