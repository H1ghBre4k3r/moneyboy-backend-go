package session

import (
	"errors"
	"fmt"

	"git.pesca.dev/pesca-dev/moneyboy-backend/internal/models"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
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

func (s *SessionService) Middleware() func(*fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		userClaims := c.Locals("user")
		if userClaims == nil {
			return c.Next()
		}
		claims := userClaims.(*jwt.Token).Claims.(jwt.MapClaims)
		sessionId := claims["id"].(string)

		// get the session
		session := s.GetSession(sessionId)
		if session == nil {
			// if there is no session, return error
			return c.SendStatus(fiber.StatusUnauthorized)
		}

		// get user
		user := s.userService.GetUser(session.UserID)
		if user == nil {
			// if there is no user, destroy session and return error
			s.DestroySession(sessionId)
			return c.SendStatus(fiber.StatusUnauthorized)
		}

		// attach session to current request
		session.User = *user
		c.Locals("session", session)
		return c.Next()
	}
}
