package jwt

import (
	"strings"

	"github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v2"
	"github.com/golang-jwt/jwt"
)

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

func (j *JWT) Middleware(filteredRoutes []string) fiber.Handler {
	return jwtware.New(jwtware.Config{
		SigningKey: []byte("mySigningKey"),
		Filter: func(c *fiber.Ctx) bool {
			for _, route := range filteredRoutes {
				if strings.HasSuffix(string(c.Request().URI().Path()), route) {
					return true
				}
			}
			return false
		},
	})
}
