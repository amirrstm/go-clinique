package middleware

import (
	"errors"
	"time"

	config "github.com/amirrstm/go-clinique/pkg/config"
	"github.com/gofiber/fiber/v2"
	jwtWare "github.com/gofiber/jwt/v2"
	"github.com/golang-jwt/jwt/v5"
)

// JWTProtected func for specify route group with JWT authentication.
// See: https://github.com/gofiber/jwt
func JWTProtected() func(*fiber.Ctx) error {
	// Create config for JWT authentication middleware.
	jwtWareConfig := jwtWare.Config{
		SigningKey:     []byte(config.AppCfg().JWTSecretKey),
		ContextKey:     "user", // used in private route
		ErrorHandler:   jwtError,
		SuccessHandler: verifyTokenExpiration,
	}

	return jwtWare.New(jwtWareConfig)
}

func verifyTokenExpiration(c *fiber.Ctx) error {
	user := c.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	expires := int64(claims["exp"].(float64))
	if time.Now().Unix() > expires {
		return jwtError(c, errors.New("token expired"))
	}

	return c.Next()
}

func jwtError(c *fiber.Ctx, err error) error {
	// Return status 401 and failed authentication error.
	if err.Error() == "Missing or malformed JWT" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"msg": err.Error(),
		})
	}

	// Return status 401 and failed authentication error.
	return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
		"msg": err.Error(),
	})
}
