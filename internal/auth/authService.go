package auth

import (
	"errors"
	"fmt"
	"time"

	"git.pesca.dev/pesca-dev/moneyboy-backend/internal/database"
	"git.pesca.dev/pesca-dev/moneyboy-backend/internal/global"
	"git.pesca.dev/pesca-dev/moneyboy-backend/internal/models"
	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

const SESSION_ID_FIELD = "id"

type JWT interface {
	Sign(map[string]interface{}) (string, error)
	Decode(string) (jwt.Claims, error)
}

type SessionService interface {
	GetSession(string) *models.Session
	CreateSession(string) (string, error)
	DestroySession(string) error
}

type AuthService struct {
	db             *database.Connection
	tokenJwt       JWT
	refreshJwt     JWT
	sessionService SessionService
}

func New(db *database.Connection, tokenJwt JWT, refreshJwt JWT, sessionService SessionService) *AuthService {
	return &AuthService{
		db,
		tokenJwt,
		refreshJwt,
		sessionService,
	}
}

// Login a user
func (s *AuthService) Login(user *global.LoginDTO) (string, string, error) {

	dbUser := s.db.Users().FindByUsername(user.Username)

	if dbUser == nil ||
		dbUser.Username != user.Username || bcrypt.CompareHashAndPassword([]byte(dbUser.Password), []byte(user.Password)) != nil {
		return "", "", errors.New("credentials do not match")
	}

	if !dbUser.EmailVerified {
		return "", "", errors.New("email not verified")
	}

	sessionId, err := s.sessionService.CreateSession(dbUser.ID)
	if err != nil {
		return "", "", err
	}

	// optain a normal token
	token, err := s.tokenJwt.Sign(createAccessTokenClaims(sessionId))
	if err != nil {
		return "", "", err
	}

	// optain a refresh token
	refreshToken, err := s.refreshJwt.Sign(map[string]interface{}{
		SESSION_ID_FIELD: sessionId,
	})
	if err != nil {
		return "", "", err
	}

	return token, refreshToken, nil
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

// Refresh a token for a valid session
// If token cant be decoded, an error will be returned, but the string will be empty
// If the token can be decoded, but is malformed or the session is not valid, an error will
// be returned aswell, but the string will contain the reason.
// If everything is successful, the new access token will be returned and the error will be nil
func (s *AuthService) RefreshToken(payload *global.RefreshTokenDTO) (string, error) {
	// try to decode the provided payload
	claims, err := s.refreshJwt.Decode(payload.RefreshToken)
	if err != nil || claims == nil {
		return "", errors.New("malformed token")
	}

	// try to access the session id
	sessionId, ok := claims.(jwt.MapClaims)[SESSION_ID_FIELD].(string)
	if !ok {
		return "invalid token provided", errors.New("invalid token provided")
	}

	// try to get the session
	session := s.sessionService.GetSession(sessionId)
	if session == nil {
		fmt.Println("no session")
		return "session invalid", errors.New("session invalid")
	}

	// if all went fine, create a new access tokne
	return s.tokenJwt.Sign(createAccessTokenClaims(sessionId))
}

// Log a user out (delete the session)
func (s *AuthService) Logout(session *models.Session) error {
	return s.sessionService.DestroySession(session.ID)
}

func createUserFromDTO(user *global.RegisterDTO) (*models.User, error) {

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	return &models.User{
		ID:            uuid.NewString(),
		Username:      user.Username,
		DisplayName:   user.DisplayName,
		Password:      string(hashedPassword),
		Email:         user.Email,
		EmailVerified: false,
	}, nil
}

// create the claims for an access token based on a session id
func createAccessTokenClaims(sessionId string) map[string]interface{} {
	return map[string]interface{}{
		SESSION_ID_FIELD: sessionId,
		"exp":            time.Now().Add(time.Minute * 15).Unix(),
	}
}
