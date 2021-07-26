package jwt

import (
	"github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v2"
	"github.com/golang-jwt/jwt"
)

type JwtParams interface {
	GetSecretKey() string
}

type JWT struct {
	secretKey string
}

func New(secretKey string) *JWT {
	return &JWT{
		secretKey,
	}
}

func (j *JWT) Sign(cls map[string]interface{}) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	for key, value := range cls {
		claims[key] = value
	}

	return token.SignedString([]byte(j.secretKey))
}

func (j *JWT) Middleware() fiber.Handler {
	return jwtware.New(jwtware.Config{
		SigningKey: []byte("mySigningKey"),
		Filter: func(c *fiber.Ctx) bool {
			// TODO lome: add param or sth else for filtered routes
			return string(c.Request().URI().LastPathSegment()) == "login"
		},
	})
}
