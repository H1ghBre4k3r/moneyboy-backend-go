package auth

import (
	"errors"
	"time"

	"git.pesca.dev/pesca-dev/moneyboy-backend/internal/database"
	"git.pesca.dev/pesca-dev/moneyboy-backend/internal/global"
	"git.pesca.dev/pesca-dev/moneyboy-backend/internal/models"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	db  *database.Connection
	jwt JWT
}

func New(db *database.Connection, jwt JWT) *AuthService {
	return &AuthService{
		db,
		jwt,
	}
}

// Login a user
func (s *AuthService) Login(user *global.LoginDTO) (interface{}, error) {

	dbUser := s.db.Users().FindByUsername(user.Username)

	if dbUser == nil ||
		dbUser.Username != user.Username || bcrypt.CompareHashAndPassword([]byte(dbUser.Password), []byte(user.Password)) != nil {
		return nil, errors.New("credentials do not match")
	}

	token, err := s.jwt.Sign(map[string]interface{}{
		"id":  dbUser.Id,
		"exp": time.Now().Add(time.Minute * 15).Unix(),
	})
	if err != nil {
		return nil, err
	}

	return struct {
		Token string `json:"token"`
	}{
		token,
	}, nil
}

// Register a new user
// returns (bool, error), where the bool is a flag for indicating an internal server error
func (s *AuthService) Register(user *global.RegisterDTO) (bool, error) {

	// check for existance of user
	if check := s.db.Users().FindByUsername(user.Username); check != nil {
		return false, errors.New("username already taken")
	}

	newUser, err := createUserFromDTO(user)
	if err != nil {
		return true, err
	}

	if err := s.db.Users().Create(newUser); err != nil {
		return true, err
	}
	return false, nil
}

func createUserFromDTO(user *global.RegisterDTO) (*models.User, error) {

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	return &models.User{
		Id:            uuid.NewString(),
		Username:      user.Username,
		DisplayName:   user.DisplayName,
		Password:      string(hashedPassword),
		Email:         user.Email,
		EmailVerified: false,
	}, nil
}
