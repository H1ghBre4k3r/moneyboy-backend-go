package auth

import (
	"git.pesca.dev/pesca-dev/moneyboy-backend/internal/models"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type AuthService struct {
	db *gorm.DB
}

func createService() *AuthService {

	// TODO lome: move this to own module
	dsn := "root:12345678@tcp(127.0.0.1:3306)/moneyboy?charset=utf8mb4&parseTime=True&loc=Local"

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect to database")
	}

	db.AutoMigrate(&models.User{})
	return &AuthService{
		db,
	}
}

func (s *AuthService) Login(c *fiber.Ctx, user *LoginDTO) error {

	dbUser := &models.User{}
	s.db.Where("username = ?", user.Username).First(dbUser)

	if dbUser.Username != user.Username || bcrypt.CompareHashAndPassword([]byte(dbUser.Password), []byte(user.Password)) != nil {
		return c.SendStatus(fiber.StatusUnauthorized)
	}

	return c.JSON(dbUser)
}

func (s *AuthService) Register(c *fiber.Ctx, user *RegisterDTO) error {

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	newUser := &models.User{
		Id:            uuid.NewString(),
		Username:      user.Username,
		DisplayName:   user.DisplayName,
		Password:      string(hashedPassword),
		Email:         user.Email,
		EmailVerified: false,
	}

	s.db.Create(newUser)
	return c.SendStatus(fiber.StatusAccepted)
}
